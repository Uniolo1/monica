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
	"os"
	"os/signal"
	"syscall"

	"codeberg.org/jersey/lightning/pkg/lightning"
	"codeberg.org/jersey/lightning/pkg/platforms/irc"

	"github.com/uniolo1/monica/internal/abstractions"
	"github.com/uniolo1/monica/internal/commands"
	"github.com/uniolo1/monica/internal/helpers"
)

func formatFunctionThing(input string) string {
	return "'" + input + "'"
}

func adder(command abstractions.Command) error {
	executor := func(opts *lightning.CommandOptions) {
		// implement ArgumentGetter
		get_args := func(input string, _ abstractions.Command_argument_type) string {
			return opts.Arguments[input]
		}

		// supply context
		context := abstractions.CommandContext{
			Framework:         "lightning",
			FormatCodeOrQuote: formatFunctionThing,
			ArgumentGetter:    get_args,
		}

		reply, _ := command.Handler(&context)

		// NOTE: we always pass 'false' here due to IRC concerns
		opts.Reply(&lightning.Message{Content: reply}, false)
	}

	arguments := make([]lightning.CommandArgument, 0, len(command.Arguments))
	for _, arg := range command.Arguments {
		arguments = append(arguments, lightning.CommandArgument{
			Name:        arg.Name,
			Description: arg.Description,
		})
	}

	bot.AddCommand(lightning.Command{
		Name:        command.Name,
		Description: command.Description,
		Arguments:   arguments,
		Executor:    executor,
	})
	return nil
}

var bot *lightning.Bot = nil

func main() {
	helpers.SetupEnv()
	prefix := "!"

	// change prefix based on environment variable
	prefix_env := os.Getenv("BOT_PREFIX")
	if prefix_env != "" {
		prefix = prefix_env
	}

	bot = lightning.NewBot(prefix)
	commands.AddCommands(adder)

	// data thingy - currently unused
	var data lightning.PluginDataStorage

	platforms := 0

	// add IRC stuff
	if os.Getenv("DO_IRC") != "" && os.Getenv("DO_IRC") != "false" {
		println("Adding IRC bot...")
		platforms++

		bot.AddPluginType("irc", irc.New)
		bot.UsePluginType(
			"irc",
			"irc",
			map[string]string{
				// basic IRC config
				"server": os.Getenv("IRC_SERVER"),
				"nick":   os.Getenv("IRC_NICK"),
				"tls":    os.Getenv("IRC_TLS"),

				// SASL stuff
				"sasl_user": os.Getenv("IRC_SASL_USER"),
				"sasl_pass": os.Getenv("IRC_SASL_PASS"),

				// any connect commands:
				"connect_cmds": os.Getenv("IRC_COMMANDS_ONCONNECT"),
			},
			data,
		)
	}

	// clear some sensitive environment variables
	os.Unsetenv("IRC_SASL_USER")
	os.Unsetenv("IRC_SASL_PASS")
	os.Unsetenv("IRC_COMMANDS_ONCONNECT") // likely contains NickServ commands

	if platforms == 0 {
		println("Bot was not configured or no platform was specified!")
		os.Exit(1)
	}

	println("Started! ^C to stop...")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
	<-quitChannel

	// '\n' is necessary so '^C' appears on it's own line
	println("\nStopped!")
}
