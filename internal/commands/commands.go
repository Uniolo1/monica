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

	"github.com/uniolo1/monica/internal/abstractions"
	"github.com/uniolo1/monica/internal/helpers"
)

func about_command(args *abstractions.Command_context) (string, bool) {
	// the about text
	about_text := fmt.Sprintf("Monica (%s) - https://github.com/uniolo1/monica", args.Framework)
	return about_text, false
}

func ping_command(args *abstractions.Command_context) (string, bool) {
	return "pong!", false
}

func define_command(args *abstractions.Command_context) (string, bool) {
	arg := args.Arguments["word"]
	if arg == "" {
		return "No definition provided!", true
	} else {
		return helpers.FormatTabs(helpers.GetDefinition(arg)), false
	}
}

func evaluate_command(args *abstractions.Command_context) (string, bool) {
	if os.Getenv("DISABLE_CALC") != "" && os.Getenv("DISABLE_CALC") != "false" {
		return "The calculator has been disabled!", false
	} else {
		arg := args.Arguments["expression"]
		if arg == "" {
			return "No expression provided!", true
		} else {
			return helpers.Evaluate_input(arg), false
		}
	}
}

func fortune_command(args *abstractions.Command_context) (string, bool) {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return fmt.Sprintf("Error: %s", err), true
	} else {
		return string(out), false
	}
}

func uwuify_command(args *abstractions.Command_context) (string, bool) {
	arg := args.Arguments["input"]
	if arg == "" {
		return "No text provided!", true
	} else {
		return helpers.Uwuify(arg), false
	}
}

func ship_command(args *abstractions.Command_context) (string, bool) {
	person_a := args.Arguments["person_a"]
	person_b := args.Arguments["person_b"]
	if person_a == "" || person_b == "" {
		return "Missing required 'person_a' and 'person_b' parameters!", true
	} else {
		output := person_a[:int(len(person_a)/2)] + person_b[int(len(person_b)/2):]
		return output, false
	}
}

func AddCommands(adder abstractions.Command_adder) error {
	var err error

	err = adder(abstractions.Command{
		Name:        "about",
		Description: "Display's information about the bot",
		Handler:     about_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "ping",
		Description: "ping-pong",
		Handler:     ping_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "define",
		Description: "find the definition of a word",
		Arguments: []abstractions.Command_arguments{{
			Name:        "word",
			Description: " Word to search",
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: define_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "calc",
		Description: "Calculate an expression",
		Arguments: []abstractions.Command_arguments{{
			Name:        "expression",
			Description: "Expression to evaluate",
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: evaluate_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "fortune",
		Description: "Returns the output of 'fortune'",
		Handler:     fortune_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "uwuify",
		Description: "Uwuifies text",
		Arguments: []abstractions.Command_arguments{{
			Name:        "input",
			Description: "Text to uwuify",
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: uwuify_command,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "ship",
		Description: "Ships 2 people",
		Arguments: []abstractions.Command_arguments{{
			Name:        "person_a",
			Description: "Person A",
			Arg_type:    abstractions.Command_argument_type_string,
		}, {
			Name:        "person_b",
			Description: "Person B",
			Arg_type:    abstractions.Command_argument_type_string,
		}},
		Handler: ship_command,
	})
	if err != nil {
		return err
	}

	return addConversationMessageGenerationCommands(adder)
}
