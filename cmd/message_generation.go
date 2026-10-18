package main

import (
	"log/slog"
	"math/rand"
	"os"
	"strings"

	"codeberg.org/jersey/lightning/pkg/lightning"
)

// NOTE: we return a pointer here inorder to fix a NULL-pointer bug (go is supposed to be memory safe BTW)
// Interestingly enough, switching to using a pointer fixed a NULL-pointer bug
func load_from_env_or_list(envname string, defaults []string) *[]string {
	env := os.Getenv(envname)
	if env == "" {
		return &defaults
	} else {
		ret := strings.Split(env, ":")
		return &ret
	}
}

func rand_item(list *[]string) string {
	if list == nil || len(*list) == 0 {
		slog.Error("'rand_item' got an empty list as input")
		return "Error: list was empty!"
	}
	return (*list)[rand.Intn(len(*list))]
}

var DEFAULT_INSULTS = []string{
	"You oxygen thief!",

	// everything below this comment is from the internet - https://psycatgames.com/magazine/conversation-starters/creative-insults/
	"There’s some shit on your clothes. Oh, nope. That’s just you.",
	"I would call you a pussy, but you lack warmth and depth.",
	"You are living proof that even ugly people have sex.", // refers to either them, or there parents
	"If you ate trash, it would be cannibalism.",
	"If I throw a stick, will you leave?",
	"If ignorance is bliss, then you must be one happy idiot.",
}

func insult_command(opts *lightning.CommandOptions) {
	insults := load_from_env_or_list("INSULTS", DEFAULT_INSULTS)
	opts.Reply(&lightning.Message{Content: rand_item(insults)}, false)
}

var DEFAULT_COMPLIMENTS = []string{
	"Nice message!",

	// everything below this comment is from wikihow - https://www.wikihow.com/Unique-Compliments
	// (wikihow sucks it had soo little)
	"You could talk to a brick wall and make it feel heard.",
	"Your voice is so calming, you should start an ASMR channel.",
	"Your heart must be 10 times the average size.",
}

func compliment_command(opts *lightning.CommandOptions) {
	compliments := load_from_env_or_list("COMPLIMENTS", DEFAULT_COMPLIMENTS)
	opts.Reply(&lightning.Message{Content: rand_item(compliments)}, false)
}

var DEFAULT_FLIRTS = []string{
	"Heyy",
	"Are you open source? Cuz' I want to fork your heart!", // this one is high-risk I would say... meh... they all are

	// everything below ths comment were flirts I found online - https://bestlifeonline.com/flirty-texts/
	"Hey, stranger. Stop being a stranger.",
	"Wanna trade secrets?",
	"I want to eat everything, including…",
	"My love language is physical touch. Want to see why?",
	"Want to be my little spoon?",
	"Why are you so cute?",
	"Nobody gets me like you do.",
	"I have no idea what I would do without you.",
}

func flirt_command(opts *lightning.CommandOptions) {
	// FYI if you need a bot to flirt successfully, you are completely and utterly cooked.

	flirts := load_from_env_or_list("FLIRTS", DEFAULT_FLIRTS)
	opts.Reply(&lightning.Message{Content: rand_item(flirts)}, false)
}

func addConversationMessageGenerationCommands(bot *lightning.Bot) {
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
