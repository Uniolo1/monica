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

	"github.com/uniolo1/monica/internal/abstractions"
)

// NOTE: we return a pointer here to fix a nil-pointer bug (go is supposed to be memory safe BTW)
// Interestingly enough, adding a pointer fixed a nil-pointer bug - EDIT: Nope! Turns out lightning
// is (at the time of writing this) very buggy (with IRC)
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

func insult_command(args *abstractions.Command_context) (string, bool) {
	insults := load_from_env_or_list("INSULTS", DEFAULT_INSULTS)
	return rand_item(insults), true
}

func compliment_command(args *abstractions.Command_context) (string, bool) {
	compliments := load_from_env_or_list("COMPLIMENTS", DEFAULT_COMPLIMENTS)
	return rand_item(compliments), true
}

func flirt_command(args *abstractions.Command_context) (string, bool) {
	// FYI if you need a bot to flirt successfully, you are completely and utterly cooked.

	flirts := load_from_env_or_list("FLIRTS", DEFAULT_FLIRTS)
	return rand_item(flirts), true
}

// requested command
func animegirl_command(args *abstractions.Command_context) (string, bool) {
	girls := load_from_env_or_list("ANIMEGIRLS", DEFAULT_ANIMEGIRL_GIFS)
	return rand_item(girls), false
}

func addConversationMessageGenerationCommands(adder abstractions.Command_adder) error {
	var err error

	err = adder(abstractions.Command{
		Name:        "insult",
		Description: "Generate an insult",
		Handler:     insult_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "compliment",
		Description: "Generate a compliment",
		Handler:     compliment_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "flirt",
		Description: "Generate a flirtatious compliment",
		Handler:     flirt_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "animegirl",
		Description: "Retrieve a random gif of an anime girl!",
		Handler:     animegirl_command,
	})
	if err != nil {
		return err
	}

	return nil
}
