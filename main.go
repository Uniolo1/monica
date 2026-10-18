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
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"codeberg.org/jersey/lightning/pkg/lightning"
	"codeberg.org/jersey/lightning/pkg/platforms/discord"
	"codeberg.org/jersey/lightning/pkg/platforms/irc"
)

func formatTabs(s string) string {
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		width := 0
		for _, r := range line {
			if r == '\t' {
				width += 8 - (width % 8)
			} else {
				width++
			}
		}

		padding := 8 - (width % 8)
		if padding != 8 {
			line += strings.Repeat(" ", padding)
		}

		lines[i] = strings.ReplaceAll(line, "\t", "        ")
	}

	return strings.Join(lines, "\n")
}

func getDefinition(query string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	apiURL := "https://api.datamuse.com/words?sp=" +
		url.QueryEscape(query) + "&md=d"

	resp, err := client.Get(apiURL)
	if err != nil {
		return "Dictionary is down"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "Dictionary is down"
	}

	var data []struct {
		Defs []string `json:"defs"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "Dictionary is down"
	}

	if len(data) == 0 || len(data[0].Defs) == 0 {
		return "No definition found"
	}

	return strings.Join(data[0].Defs, "\n")
}

func ping_command(opts *lightning.CommandOptions) {
	opts.Reply(&lightning.Message{Content: "pong!"}, false)
}

func define_command(opts *lightning.CommandOptions) {
	opts.Reply(&lightning.Message{Content: formatTabs(getDefinition(opts.Arguments["word"]))}, false)
}

func addCommands(bot *lightning.Bot) {
	bot.AddCommand(lightning.Command{
		Name:        "ping",
		Description: "ping-pong",
		Executor:    ping_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "define",
		Description: "Lookup a word",
		Arguments: []lightning.CommandArgument{{
			Name:        "word",
			Description: "Word to search",
		}},
		Executor: define_command,
	})
}

func main() {
	// load configuration (environment)
	_ = godotenv.Load("/etc/monica.env")
	_ = godotenv.Load("~/.config/monica/monica.env")
	_ = godotenv.Load(".env")

	prefix := "!"

	// change prefix based on environment variable
	prefix_env := os.Getenv("BOT_PREFIX")
	if prefix_env != "" {
		prefix = prefix_env
	}

	bot := lightning.NewBot(prefix)
	addCommands(bot)

	// data thingy - currently unused
	var data lightning.PluginDataStorage

	// add discord stuff
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

	// add IRC stuff
	if os.Getenv("DO_IRC") != "" && os.Getenv("DO_IRC") != "false" {
		println("Adding IRC bot...")
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

	println("Started! ^C to stop...")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
	<-quitChannel

	// '\n' is necessary so '^C' appears on it's own line
	println("\nStopped!")
}
