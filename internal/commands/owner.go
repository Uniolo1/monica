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
	"fmt"
	"os"

	"github.com/uniolo1/monica/internal/abstractions"
)

func ownerGetenvCommand(args *abstractions.CommandContext) (string, bool) {
	key := args.ArgumentGetter("key", abstractions.Command_argument_string)
	if key == "" {
		return "Missing required " + args.FormatQuote("key") + " parameter!", true
	} else {
		value := os.Getenv(key)
		if value == "" {
			return "Not found!", true
		} else {
			return value, true
		}
	}
}

func ownerUnsetenvCommand(args *abstractions.CommandContext) (string, bool) {
	key := args.ArgumentGetter("key", abstractions.Command_argument_string)
	if key == "" {
		return "Missing required " + args.FormatQuote("key") + " parameter!", true
	} else {
		err := os.Unsetenv(key)
		if err != nil {
			return "Unable to unset " + args.FormatQuote(key) + ", most likely not found!", true
		} else {
			return "Ok", true
		}
	}
}

func ownerSetenvCommand(args *abstractions.CommandContext) (string, bool) {
	key := args.ArgumentGetter("key", abstractions.Command_argument_string)
	value := args.ArgumentGetter("value", abstractions.Command_argument_string)
	if key == "" || value == "" {
		return "Missing required " + args.FormatQuote("key") + " and/or " + args.FormatQuote("value") + " parameter(s)!", true
	} else {
		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Sprintf("Error: %s", err), true
		} else {
			return "Ok", true
		}
	}
}

func ownerClearaiCommand(args *abstractions.CommandContext) (string, bool) {
	aiHistoryMu.Lock()
	defer aiHistoryMu.Unlock()
	clear(aiHistory)

	return "All AI zones cleared!", false
}

func addOwnerCommands(adder abstractions.CommandAdder) error {
	var err error

	err = adder(abstractions.Command{
		Name:        "owner getenv",
		Description: "Owner-only command to get an environment variable",
		Arguments: []abstractions.CommandArguments{{
			Name:        "key",
			Description: "Name of the environment variable",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		OwnerOnly: true,
		Handler:   ownerGetenvCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "owner unsetenv",
		Description: "Owner-only command to unset an environment variable",
		Arguments: []abstractions.CommandArguments{{
			Name:        "key",
			Description: "Name of the environment variable",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		OwnerOnly: true,
		Handler:   ownerUnsetenvCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "owner setenv",
		Description: "Owner-only command to set an environment variable",
		Arguments: []abstractions.CommandArguments{{
			Name:        "key",
			Description: "Name of the environment variable",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "value",
			Description: "Value to set the environment variable to",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		OwnerOnly: true,
		Handler:   ownerSetenvCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "owner clearai",
		Description: "Owner-only command to clear everyone's AI history!",
		Arguments:   []abstractions.CommandArguments{},
		OwnerOnly:   true,
		Handler:     ownerClearaiCommand,
	})
	if err != nil {
		return err
	}

	return nil
}
