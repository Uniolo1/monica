package main

import (
	"codeberg.org/jersey/lightning/pkg/lightning"
)

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
