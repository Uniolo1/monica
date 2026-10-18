package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"codeberg.org/jersey/lightning/pkg/lightning"
)

func formatTabs(s string) string {
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		width := 0
		for _, r := range line {
			if r == '\t' {
				width += 8 - (width % 8)
			} else {
				width++
			}
		}

		padding := 8 - (width % 8)
		if padding != 8 {
			line += strings.Repeat(" ", padding)
		}

		lines[i] = strings.ReplaceAll(line, "\t", "        ")
	}

	return strings.Join(lines, "\n")
}

func getDefinition(query string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	apiURL := "https://api.datamuse.com/words?sp=" +
		url.QueryEscape(query) + "&md=d"

	resp, err := client.Get(apiURL)
	if err != nil {
		return "Dictionary is down"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "Dictionary is down"
	}

	var data []struct {
		Defs []string `json:"defs"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "Dictionary is down"
	}

	if len(data) == 0 || len(data[0].Defs) == 0 {
		return "No definition found"
	}

	return strings.Join(data[0].Defs, "\n")
}

func ping_command(opts *lightning.CommandOptions) {
	opts.Reply(&lightning.Message{Content: "pong!"}, false)
}

func define_command(opts *lightning.CommandOptions) {
	opts.Reply(&lightning.Message{Content: formatTabs(getDefinition(opts.Arguments["word"]))}, false)
}

func addCommands(bot *lightning.Bot) {
	bot.AddCommand(lightning.Command{
		Name:        "ping",
		Description: "ping-pong",
		Executor:    ping_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "define",
		Description: "Lookup a word",
		Arguments: []lightning.CommandArgument{{
			Name:        "word",
			Description: "Word to search",
		}},
		Executor: define_command,
	})
}
