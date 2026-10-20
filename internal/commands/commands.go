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
	"os/exec"
	"strconv"

	"github.com/uniolo1/monica/internal/abstractions"
	"github.com/uniolo1/monica/internal/helpers"
)

func aboutCommand(args *abstractions.CommandContext) (string, bool) {
	// the about text
	about_text := fmt.Sprintf("Monica (%s) - https://github.com/uniolo1/monica", args.Framework)
	return about_text, false
}

func pingCommand(args *abstractions.CommandContext) (string, bool) {
	return "pong!", false
}

func defineCommand(args *abstractions.CommandContext) (string, bool) {
	arg := args.ArgumentGetter("word", abstractions.Command_argument_type_string)
	if arg == "" {
		return "No definition provided!", true
	} else {
		var upper_limit int
		upper_limit_arg := args.ArgumentGetter("limit", abstractions.Command_argument_type_int)
		if upper_limit_arg == "" {
			upper_limit = 5 // default
		} else {
			n, err := strconv.Atoi(upper_limit_arg)
			if err != nil {
				return "Invalid 'limit'!", true
			}

			// set upper_limit to n
			upper_limit = n
		}

		if upper_limit <= 0 {
			return "Upper limit of dictionary items must be a minimum of 1!", true
		}

		return helpers.FormatTabs(helpers.GetDefinition(arg, upper_limit)), false
	}
}

func evaluateCommand(args *abstractions.CommandContext) (string, bool) {
	if os.Getenv("DISABLE_CALC") != "" && os.Getenv("DISABLE_CALC") != "false" {
		return "The calculator has been disabled!", false
	} else {
		arg := args.ArgumentGetter("expression", abstractions.Command_argument_type_string)
		if arg == "" {
			return "No expression provided!", true
		} else {
			return helpers.EvaluatecCalculatorInput(arg), false
		}
	}
}

func fortuneCommand(args *abstractions.CommandContext) (string, bool) {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return fmt.Sprintf("Error: %s", err), true
	} else {
		return string(out), false
	}
}

func uwuifyCommand(args *abstractions.CommandContext) (string, bool) {
	arg := args.ArgumentGetter("input", abstractions.Command_argument_type_string)
	if arg == "" {
		return "No text provided!", true
	} else {
		return helpers.Uwuify(arg), false
	}
}

func shipCommand(args *abstractions.CommandContext) (string, bool) {
	person_a := args.ArgumentGetter("person_a", abstractions.Command_argument_type_string)
	person_b := args.ArgumentGetter("person_b", abstractions.Command_argument_type_string)
	if person_a == "" || person_b == "" {
		return "Missing required 'person_a' and 'person_b' parameters!", true
	} else {
		output := person_a[:int(len(person_a)/2)] + person_b[int(len(person_b)/2):]
		return output, false
	}
}

func AddCommands(adder abstractions.CommandAdder) error {
	var err error

	err = adder(abstractions.Command{
		Name:        "about",
		Description: "Display's information about the bot",
		Handler:     aboutCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "ping",
		Description: "ping-pong",
		Handler:     pingCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "define",
		Description: "find the definition of a word",
		Arguments: []abstractions.CommandArguments{{
			Name:        "word",
			Description: " Word to search",
			Required:    true,
			Arg_type:    abstractions.Command_argument_type_string,
		}, {
			Name:        "limit",
			Description: "Maximum number of dictionary entries that are returned",
			Required:    false,
			Arg_type:    abstractions.Command_argument_type_int,
		}},
		Handler: defineCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "calc",
		Description: "Calculate an expression",
		Arguments: []abstractions.CommandArguments{{
			Name:        "expression",
			Description: "Expression to evaluate",
			Required:    true,
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: evaluateCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "fortune",
		Description: "Returns the output of 'fortune'",
		Handler:     fortuneCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "uwuify",
		Description: "Uwuifies text",
		Arguments: []abstractions.CommandArguments{{
			Name:        "input",
			Description: "Text to uwuify",
			Required:    true,
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: uwuifyCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "ship",
		Description: "Ships 2 people",
		Arguments: []abstractions.CommandArguments{{
			Name:        "person_a",
			Description: "Person A",
			Required:    true,
			Arg_type:    abstractions.Command_argument_type_string,
		}, {
			Name:        "person_b",
			Description: "Person B",
			Required:    true,
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: shipCommand,
	})
	if err != nil {
		return err
	}

	return addConversationMessageGenerationCommands(adder)
}
