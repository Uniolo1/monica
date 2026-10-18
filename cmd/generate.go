package main

import (
	"math/rand"
	"os"
	"strings"

	"codeberg.org/jersey/lightning/pkg/lightning"
)

func load_from_env_or_list(envname string, defaults []string) []string {
	env := os.Getenv(envname)
	if env == "" {
		return defaults
	} else {
		return strings.Split(env, ":")
	}
}

var DEFAULT_INSULTS = []string{
	"You oxygen thief!",
}

func insult_command(opts *lightning.CommandOptions) {
	insults := load_from_env_or_list("INSULTS", DEFAULT_INSULTS)
	insult := insults[rand.Intn(len(insults))]

	opts.Reply(&lightning.Message{Content: insult}, false)
}

var DEFAULT_COMPLIMENTS = []string{
	"Nice message!",
}

func compliment_command(opts *lightning.CommandOptions) {
	compliments := load_from_env_or_list("COMPLIMENTS", DEFAULT_COMPLIMENTS)
	compliment := compliments[rand.Intn(len(compliments))]

	opts.Reply(&lightning.Message{Content: compliment}, false)
}

var DEFAULT_FLIRTS = []string{
	"Heyy", // temporary: do not know what to put yet
}

func flirt_command(opts *lightning.CommandOptions) {
	flirts := load_from_env_or_list("FLIRTS", DEFAULT_FLIRTS)
	flirt := flirts[rand.Intn(len(flirts))]

	opts.Reply(&lightning.Message{Content: flirt}, false)
}

func addGenCommands(bot *lightning.Bot) {
	bot.AddCommand(lightning.Command{
		Name:        "insult",
		Description: "Generate an insult",
		Executor:    insult_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "compliment",
		Description: "Generate a compliment",
		Executor:    compliment_command,
	})

	bot.AddCommand(lightning.Command{
		Name:        "flirt",
		Description: "Generate a flirtatious compliment",
		Executor:    flirt_command,
	})
}
