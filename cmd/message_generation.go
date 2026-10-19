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

var DEFAULT_INSULTS = []string{
	"You oxygen thief!",

	// everything below this comment is from the internet - https://psycatgames.com/magazine/conversation-starters/creative-insults/
	"There’s some shit on your clothes. Oh, nope. That’s just you.",
	"I would call you a pussy, but you lack warmth and depth.",
	"You are living proof that even ugly people have sex.", // refers to either them, or there parents
	"If you ate trash, it would be cannibalism.",
	"If I throw a stick, will you leave?",
	"If ignorance is bliss, then you must be one happy idiot.",
}

func insult_command(opts *lightning.CommandOptions) {
	insults := load_from_env_or_list("INSULTS", DEFAULT_INSULTS)
	opts.Reply(&lightning.Message{Content: rand_item(insults)}, false)
}

var DEFAULT_COMPLIMENTS = []string{
	"Nice message!",

	// everything below this comment is from wikihow - https://www.wikihow.com/Unique-Compliments
	// (wikihow sucks it had soo little)
	"You could talk to a brick wall and make it feel heard.",
	"Your voice is so calming, you should start an ASMR channel.",
	"Your heart must be 10 times the average size.",
}

func compliment_command(opts *lightning.CommandOptions) {
	compliments := load_from_env_or_list("COMPLIMENTS", DEFAULT_COMPLIMENTS)
	opts.Reply(&lightning.Message{Content: rand_item(compliments)}, false)
}

var DEFAULT_FLIRTS = []string{
	"Heyy",
	"Are you open source? Cuz' I want to fork your left atrium!", // this one is high-risk I would say... meh... they all are

	// An acquaintance of mine (Zachary) recommended I add this... keeping it as a comment
	// "Stawp! I am not a tsundere!",

	// everything below ths comment were flirts I found online - https://bestlifeonline.com/flirty-texts/
	"Hey, stranger. Stop being a stranger.",
	"Wanna trade secrets?",
	"I want to eat everything, including…",
	"My love language is physical touch. Want to see why?",
	"Want to be my little spoon?",
	"Why are you so cute?",
	"Nobody gets me like you do.",
	"I have no idea what I would do without you.",
}

func flirt_command(opts *lightning.CommandOptions) {
	// FYI if you need a bot to flirt successfully, you are completely and utterly cooked.

	flirts := load_from_env_or_list("FLIRTS", DEFAULT_FLIRTS)
	opts.Reply(&lightning.Message{Content: rand_item(flirts)}, false)
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
