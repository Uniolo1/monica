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

	"github.com/joho/godotenv"

	"codeberg.org/jersey/lightning/pkg/lightning"
	"codeberg.org/jersey/lightning/pkg/platforms/discord"
)

func main() {
	// load environment from .env
	_ = godotenv.Load()

	bot := lightning.NewBot("!")

	bot.AddCommand(lightning.Command{
		Name:        "ping",
		Description: "ping-pong",
		Executor: func(opts *lightning.CommandOptions) {
			opts.Reply(&lightning.Message{Content: "pong!"}, false)
		},
	})

	// data thingy - currently unused
	var data lightning.PluginDataStorage

	discord_bot_token := os.Getenv("DISCORD_BOT_TOKEN")
	if discord_bot_token != "" {
		println("Adding Discord bot...")
		// add plugin
		bot.AddPluginType("discord", discord.New)
		bot.UsePluginType(
			"discord",
			"discord",
			map[string]string{
				"token": discord_bot_token,
			},
			data,
		)
	}

	println("Started! ^C to stop...")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
	<-quitChannel

	// '\n' is necessary so '^C' appears on it's own line
	println("\nStopped!")
}
