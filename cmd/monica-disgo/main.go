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
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"

	"github.com/uniolo1/monica/internal/abstractions"
	"github.com/uniolo1/monica/internal/commands"
	"github.com/uniolo1/monica/internal/helpers"
)

var h *handler.Mux = nil
var slash_commands []discord.ApplicationCommandCreate

func adder(command abstractions.Command) error {
	executor := func(data discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
		// implement ArgumentGetter
		get_args := func(input string) string {
			ret, is_found := data.OptString(input)
			if is_found {
				return ret
			} else {
				return ""
			}
		}

		// supply context
		context := abstractions.CommandContext{
			Framework:      "disgo",
			ArgumentGetter: get_args,
		}

		reply, ephemeral := command.Handler(&context)

		return e.CreateMessage(discord.MessageCreate{
			Content: reply,
		}.WithEphemeral(ephemeral))
	}

	options := []discord.ApplicationCommandOption{}
	for _, arg := range command.Arguments {
		// TODO: Add int's when we add a command that uses them
		options = append(options, discord.ApplicationCommandOptionString{
			Name:        arg.Name,
			Description: arg.Description,
			Required:    arg.Required,
		})
	}

	newCommand := discord.SlashCommandCreate{
		Name:        command.Name,
		Description: command.Description,
		Options:     options,
		NSFW:        &command.NSFW,
	}

	h.SlashCommand("/"+command.Name, executor)

	slash_commands = append(slash_commands, newCommand)
	return nil
}

func main() {
	helpers.SetupEnv()
	if os.Getenv("DISCORD_BOT_TOKEN") == "" {
		fmt.Println("Missing 'DISCORD_BOT_TOKEN'")
		os.Exit(1)
	}

	h = handler.New()
	client, err := disgo.New(
		os.Getenv("DISCORD_BOT_TOKEN"),
		bot.WithDefaultGateway(),
		bot.WithEventListeners(h),
	)
	if err != nil {
		panic(err)
	}

	// add the commands
	commands.AddCommands(adder)
	if err := handler.SyncCommands(client, slash_commands, []snowflake.ID{}); err != nil {
		panic(err)
	}

	// connect to gateway
	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s

	os.Unsetenv("DISCORD_BOT_TOKEN")
}
