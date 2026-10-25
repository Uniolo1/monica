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
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/uniolo1/monica/internal/abstractions"
	"github.com/uniolo1/monica/internal/helpers"
)

func aboutCommand(args *abstractions.CommandContext) (string, bool) {
	// the about text
	about_text := fmt.Sprintf("Monica (%s) - https://github.com/uniolo1/monica", args.Framework)
	return about_text, false
}

func pingCommand(args *abstractions.CommandContext) (string, bool) {
	// debug command - the first added!
	return "pong!", false
}

func defineCommand(args *abstractions.CommandContext) (string, bool) {
	arg := args.ArgumentGetter("word", abstractions.Command_argument_string)
	if arg == "" {
		return "No definition provided!", true
	} else {
		var upper_limit int

		upper_limit_arg := args.ArgumentGetter("limit", abstractions.Command_argument_int)
		if upper_limit_arg == "" {
			upper_limit = 5 // default
		} else {
			n, err := strconv.Atoi(upper_limit_arg)
			if err != nil {
				return "Invalid " + args.FormatQuote("limit") + "!", true
			}

			// set upper_limit to n
			upper_limit = n
		}

		if upper_limit <= 0 {
			return "Upper limit of dictionary items must be no less then 1!", true
		}

		return helpers.FormatTabs(helpers.GetDefinition(arg, upper_limit)), false
	}
}

func evaluateCommand(args *abstractions.CommandContext) (string, bool) {
	if os.Getenv("DISABLE_CALC") != "" && os.Getenv("DISABLE_CALC") != "false" {
		return "The calculator has been disabled!", false
	} else {
		arg := args.ArgumentGetter("expression", abstractions.Command_argument_string)
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

func diceCommand(args *abstractions.CommandContext) (string, bool) {
	arg := args.ArgumentGetter("input", abstractions.Command_argument_string)
	if arg == "" {
		arg = "5d6"
	}

	result, rolls, err := helpers.RollDice(arg)
	if err != nil {
		return fmt.Sprintf("Error: %s", err), true
	}

	// logic to check for args.NoUnicode
	diceIcon := "🎲"
	if args.NoUnicode {
		diceIcon = ""
	}

	var output string = fmt.Sprintf("%v:", result)
	for _, i := range rolls {
		output += fmt.Sprintf(" %s%v", diceIcon, i)
	}

	return output, false
}

func uwuifyCommand(args *abstractions.CommandContext) (string, bool) {
	arg := args.ArgumentGetter("input", abstractions.Command_argument_string)
	if arg == "" {
		return "No text provided!", true
	} else {
		return helpers.Uwuify(arg), false
	}
}

func confessCommand(args *abstractions.CommandContext) (string, bool) {
	arg := args.ArgumentGetter("confession", abstractions.Command_argument_string)
	if arg == "" {
		return "No confession provided!", true
	} else {
		err := helpers.Confess(arg)
		if err != nil {
			return fmt.Sprintf("Error confessing: %s", err), true
		} else {
			return "Confession submitted!", true
		}
	}
}

func shipCommand(args *abstractions.CommandContext) (string, bool) {
	person_a := args.ArgumentGetter("person-a", abstractions.Command_argument_string)
	person_b := args.ArgumentGetter("person-b", abstractions.Command_argument_string)
	if person_a == "" || person_b == "" {
		return "Missing required " + args.FormatQuote("person-a") + " and/or " + args.FormatQuote("person-b") + " parameter(s)!", true
	} else {
		output := person_a[:int(len(person_a)/2)] + person_b[int(len(person_b)/2):]
		return output, false
	}
}

func meterCommand(args *abstractions.CommandContext) (string, bool) {
	thing := strings.ToLower(args.ArgumentGetter("thing", abstractions.Command_argument_string))
	metric := strings.ToLower(args.ArgumentGetter("metric", abstractions.Command_argument_string))
	if thing == "" || metric == "" {
		return "Missing required " + args.FormatQuote("thing") + " and/or " + args.FormatQuote("metric") + " parameter(s)!", true
	} else {
		// logic to get a seed - the seed generated changes every day
		now := time.Now()
		hash := sha256.Sum256([]byte(
			thing + "\x00" +
				metric + "\x00" +
				strconv.Itoa(now.YearDay()) + "\x00" +
				strconv.Itoa(now.Year()),
		))

		seed := int64(binary.BigEndian.Uint64(hash[:8]))

		// calculate the percent using said seed
		percent := rand.New(rand.NewSource(seed)).Intn(101)

		return fmt.Sprintf("%s's %s is at %d%%", args.FormatQuote(thing), args.FormatQuote(metric), percent), false
	}
}

func convertTzCommand(args *abstractions.CommandContext) (string, bool) {
	time_str := args.ArgumentGetter("time", abstractions.Command_argument_string)
	from_tz := args.ArgumentGetter("from_tz", abstractions.Command_argument_string)
	to_tz := args.ArgumentGetter("to_tz", abstractions.Command_argument_string)

	if time_str == "" || from_tz == "" {
		return "Some required parameters are missing, one of: " + args.FormatQuote("time") + ", or " + args.FormatQuote("from_tz") + "!", true
	}

	from_loc, err := time.LoadLocation(from_tz)
	if err != nil {
		return "Invalid " + args.FormatQuote("from_tz") + ": " + err.Error(), true
	}

	to_loc, err := time.LoadLocation(to_tz)
	if err != nil {
		return "Invalid " + args.FormatQuote("to_tz") + ": " + err.Error(), true
	}

	layouts := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"15:04:05",
		"15:04",
	}

	var parsedTime time.Time
	var layout string
	var parseErr error

	for _, candidate := range layouts {
		parsedTime, parseErr = time.ParseInLocation(candidate, time_str, from_loc)
		if parseErr == nil {
			layout = candidate
			break
		}
	}

	if layout == "" {
		return "Time is in an unrecognized format!", true
	}

	result := parsedTime.In(to_loc).Format(layout)

	if to_tz == "" {
		// append " UTC" to result it 'to_tz' is not specified
		result += " UTC"
	}
	converted := result
	return converted, false
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
			Description: "Word to search",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "limit",
			Description: "Maximum number of dictionary entries that are returned",
			Required:    false,
			Arg_type:    abstractions.Command_argument_int,
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
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: evaluateCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "converttz",
		Description: "Convert a the timezone of a time from 'from_tz' to 'to_tz'",
		Arguments: []abstractions.CommandArguments{{
			Name:        "time",
			Description: "The time to convert",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "from_tz",
			Description: "The timezone 'time' is currently in",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "to_tz",
			Description: "The timezone you want 'time' to be in",
			Required:    false,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: convertTzCommand,
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
		Name:        "rtd",
		Description: "Roll some dice!",
		Arguments: []abstractions.CommandArguments{{
			Name:        "input",
			Description: "How many dice and how many sides (e.g. '5d6')",
			Required:    false,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: diceCommand,
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
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: uwuifyCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "confess",
		Description: "Submit an anonymous confession!",
		Arguments: []abstractions.CommandArguments{{
			Name:        "confession",
			Description: "Submitted anonymously!",
			Required:    true,
			Sensitive:   true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: confessCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "ship",
		Description: "Ships 2 people",
		Arguments: []abstractions.CommandArguments{{
			Name:        "person-a",
			Description: "Person A",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "person-b",
			Description: "Person B",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: shipCommand,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "meter",
		Description: "Measure thing's metric",
		Arguments: []abstractions.CommandArguments{{
			Name:        "thing",
			Description: "Thing to measure the thing of",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "metric",
			Description: "Thing to measure",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: meterCommand,
	})
	if err != nil {
		return err
	}

	// run additional setup commands below

	err = addConversationMessageGenerationCommands(adder)
	if err != nil {
		return err
	}

	err = addAiCommands(adder)
	if err != nil {
		return err
	}

	err = addOwnerCommands(adder)
	if err != nil {
		return err
	}

	return nil
}
