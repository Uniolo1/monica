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
	"log/slog"
	"math/rand"
	"os"
	"strings"

	"codeberg.org/jersey/lightning/pkg/lightning"
)

// NOTE: we return a pointer here to fix a nil-pointer bug (go is supposed to be memory safe BTW)
// Interestingly enough, adding a pointer fixed a nil-pointer bug
func load_from_env_or_list(envname string, defaults []string) *[]string {
	env := os.Getenv(envname)
	if env == "" {
		return &defaults
	} else {
		ret := strings.Split(env, ":")
		return &ret
	}
}

func rand_item(list *[]string) string {
	if list == nil || len(*list) == 0 {
		slog.Error("'rand_item' got an empty list as input")
		return "Error: list was empty!"
	}
	return (*list)[rand.Intn(len(*list))]
}

func insult_command(opts *lightning.CommandOptions) {
	insults := load_from_env_or_list("INSULTS", DEFAULT_INSULTS)
	opts.Reply(&lightning.Message{Content: rand_item(insults)}, false)
}

func compliment_command(opts *lightning.CommandOptions) {
	compliments := load_from_env_or_list("COMPLIMENTS", DEFAULT_COMPLIMENTS)
	opts.Reply(&lightning.Message{Content: rand_item(compliments)}, false)
}

func flirt_command(opts *lightning.CommandOptions) {
	// FYI if you need a bot to flirt successfully, you are completely and utterly cooked.

	flirts := load_from_env_or_list("FLIRTS", DEFAULT_FLIRTS)
	opts.Reply(&lightning.Message{Content: rand_item(flirts)}, false)
}

// requested command
func animegirl_command(opts *lightning.CommandOptions) {
	girls := load_from_env_or_list("ANIMEGIRLS", DEFAULT_ANIMEGIRL_GIFS)
	opts.Reply(&lightning.Message{Content: rand_item(girls)}, false)
}

// doesn't fit with the others... putting it here anyways
func ship_command(opts *lightning.CommandOptions) {
	person_a := opts.Arguments["person_a"]
	person_b := opts.Arguments["person_b"]
	if person_a == "" || person_b == "" {
		opts.Reply(&lightning.Message{Content: "Missing required 'person_a' and 'person_b' parameters!"}, false)
	} else {
		output := person_a[:int(len(person_a)/2)] + person_b[int(len(person_b)/2):]
		opts.Reply(&lightning.Message{Content: output}, false)
	}
}

func addConversationMessageGenerationCommands(bot *lightning.Bot) {
	bot.AddCommand(lightning.Command{
		Name:        "insult",
		Description: "Generate an insult",
		Executor:    insult_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "compliment",
		Description: "Generate a compliment",
		Executor:    compliment_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "flirt",
		Description: "Generate a flirtatious compliment",
		Executor:    flirt_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "animegirl",
		Description: "Retrieve a random gif of an anime girl!",
		Executor:    animegirl_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "ship",
		Description: "Ships 2 people!",
		Arguments: []lightning.CommandArgument{
			{
				Name:        "person_a",
				Description: "Person A",
			},
			{
				Name:        "person_b",
				Description: "Person B",
			},
		},
		Executor: ship_command,
	})
}
