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
	"os"

	"github.com/uniolo1/monica/internal/abstractions"
)

func aiAsk(args *abstractions.CommandContext) (string, bool) {
	if args.UniqueCallerID == "" {
		return "Unknown caller!", true
	}

	prompt := args.ArgumentGetter("prompt", abstractions.Command_argument_string)
	if prompt == "" {
		return "Missing required " + args.FormatQuote("prompt") + " argument!", true
	}

	key := os.Getenv("AI_KEY")
	model := os.Getenv("AI_MODEL")
	endpoint := os.Getenv("AI_ENDPOINT")

	if key == "" || model == "" || endpoint == "" {
		return "AI not configured, missing one of: " +
				args.FormatQuote("AI_KEY") + ", " +
				args.FormatQuote("AI_MODEL") + ", " +
				args.FormatQuote("AI_ENDPOINT"),
			true
	}

	return "NotImplemented", false
}

func aiReset(args *abstractions.CommandContext) (string, bool) {
	if args.UniqueCallerID == "" {
		return "Unknown caller!", true
	}

	return "NotImplemented", false
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
