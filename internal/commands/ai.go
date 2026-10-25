// SPDX-FileCopyrightText: NONE
// SPDX-License-Identifier: Unlicense

/*
 * This is free and unencumbered software released into the public domain.
 *
 * Anyone is free to copy, modify, publish, use, compile, sell, or distribute
 * this software, either in source code form or as a compiled binary, for any
 * purpose, commercial or non-commercial, and by any means.
 *
 * In jurisdictions that recognize copyright laws, the author or authors of this
 * software dedicate any and all copyright interest in the software to the
 * public domain. We make this dedication for the benefit of the public at large
 * and to the detriment of our heirs and successors. We intend this dedication
 * to be an overt act of relinquishment in perpetuity of all present and future
 * rights to this software under copyright law.
 *
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/uniolo1/monica/internal/abstractions"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

const maxHistoryMessages = 20

var aiHistory = make(map[string][]Message)
var aiHistoryMu sync.Mutex

// making helpers so I can add thread-safety stuff later
func getAiHistory(zone string) []Message {
	aiHistoryMu.Lock()
	defer aiHistoryMu.Unlock()

	return aiHistory[zone]
}

func appendAiHistory(zone string, messages ...Message) {
	aiHistoryMu.Lock()
	defer aiHistoryMu.Unlock()

	aiHistory[zone] = append(aiHistory[zone], messages...)

	// logic so it does not get too big
	if len(aiHistory[zone]) > maxHistoryMessages {
		aiHistory[zone] = append(
			// empty array:
			[]Message(nil),

			// keep the most recent messages
			aiHistory[zone][len(aiHistory[zone])-maxHistoryMessages:]...,
		)
	}
}

// again, making helpers so I can add thread-safety stuff later
func clearAiHistory(zone string) {
	aiHistoryMu.Lock()
	defer aiHistoryMu.Unlock()

	delete(aiHistory, zone)
}

func sendAiAMessage(message string, endpoint string, key string, model string, historyZone string) (string, error) {
	if !strings.HasSuffix(endpoint, "/chat/completions") {
		endpoint += "/chat/completions"
	}

	var systemPrompt string = os.Getenv("AI_PROMPT")
	if systemPrompt == "" {
		systemPrompt = DEFAULT_SYSTEM_PROMPT
	}

	// add first message with systemPrompt (TODO: make customizable by user)
	messages := []Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
	}

	messages = append(messages, getAiHistory(historyZone)...)
	messages = append(messages, Message{
		Role:    "user",
		Content: message,
	})

	payload := ChatRequest{
		Model:     model,
		MaxTokens: 350,
		Messages:  messages,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal chat request: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		bytes.NewReader(body),
	)
	if err != nil {
		return "", fmt.Errorf("create chat request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if key != "" {
		req.Header.Set("Authorization", "Bearer "+key)
	}

	client := &http.Client{
		Timeout: 65 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send chat request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return "", fmt.Errorf("read chat response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf(
			"API returned %s: %s",
			resp.Status,
			strings.TrimSpace(string(respBody)),
		)
	}

	var result ChatResponse
	err = json.Unmarshal(respBody, &result)

	if err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	// handle result where API returned no choices
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("API returned no choices")
	}

	reply := result.Choices[0].Message.Content
	appendAiHistory(
		historyZone,
		Message{Role: "user", Content: message},
		Message{Role: "assistant", Content: reply},
	)
	return reply, nil
}

func aiAsk(args *abstractions.CommandContext) (string, bool) {
	var historyZone string
	if args.UniqueChannelID != "" {
		historyZone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		historyZone = args.UniqueCallerID
	} else {
		return "Unknown caller!", true
	}
	_ = historyZone

	// only check here, not much harm in clearing AI
	if os.Getenv("DISABLE_AI") != "" && os.Getenv("DISABLE_AI") != "false" {
		return "AI features have been disabled!", false
	}

	prompt := args.ArgumentGetter("prompt", abstractions.Command_argument_string)
	if prompt == "" {
		return "Missing required " + args.FormatQuote("prompt") + " argument!", true
	}

	key := os.Getenv("AI_KEY")
	model := os.Getenv("AI_MODEL")
	endpoint := os.Getenv("AI_BASEURL")

	if key == "" || model == "" || endpoint == "" {
		return "AI not configured, missing either " +
				args.FormatQuote("AI_MODEL") + ", or " +
				args.FormatQuote("AI_BASEURL"),
			true
	}

	endpoint = strings.TrimRight(endpoint, "/")

	// defer
	if args.Defer != nil {
		err := args.Defer(true)
		if err != nil {
			return fmt.Sprintf("Error: %s", err), true
		}
	}

	msg, err := sendAiAMessage(prompt, endpoint, key, model, historyZone)

	if err != nil {
		return err.Error(), true
	}

	return msg, false
}

// TODO: Add history
func aiReset(args *abstractions.CommandContext) (string, bool) {
	var historyZone string
	if args.UniqueChannelID != "" {
		historyZone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		historyZone = args.UniqueCallerID
	} else {
		return "Unknown caller!", true
	}

	clearAiHistory(historyZone)
	return "History cleared for the " + args.FormatQuote(historyZone) + " zone!", false
}

func addAiCommands(adder abstractions.CommandAdder) error {
	var err error

	err = adder(abstractions.Command{
		Name:        "ai chat",
		Description: "Ask a question to an AI",
		Arguments: []abstractions.CommandArguments{{
			Name:        "prompt",
			Description: "What to say to the AI",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: aiAsk,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "ai reset",
		Description: "Reset the AIs history",
		Arguments:   []abstractions.CommandArguments{},
		Handler:     aiReset,
	})
	if err != nil {
		return err
	}

	return nil
}
