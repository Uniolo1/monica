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

package main

import (
	"codeberg.org/jersey/lightning/pkg/lightning"
)

func about_command(opts *lightning.CommandOptions) {
	// the about text
	about_text := "Monica - https://github.com/uniolo1/monica"

	opts.Reply(&lightning.Message{Content: about_text}, false)
}

func ping_command(opts *lightning.CommandOptions) {
	opts.Reply(&lightning.Message{Content: "pong!"}, false)
}

func define_command(opts *lightning.CommandOptions) {
	opts.Reply(&lightning.Message{Content: formatTabs(getDefinition(opts.Arguments["word"]))}, false)
}

func addCommands(bot *lightning.Bot) {
	bot.AddCommand(lightning.Command{
		Name:        "about",
		Description: "Display's information about the bot",
		Executor:    about_command,
	})

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
