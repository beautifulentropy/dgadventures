package main

import (
	"flag"
	"log"

	"github.com/beautifulentropy/dgadventures/src/router"

	"github.com/bwmarrin/discordgo"
)

// botToken  --t flag of your discord bot token
// botPrefix --p prefix your bot will listen for (e.g. "!" would be the prefix for !roll)
var (
	botToken  = flag.String("t", "", "bot token")
	botPrefix = flag.String("p", "!", "bot prefix")
)

func main() {
	flag.Parse()
	discordSession, err := discordgo.New("Bot " + *botToken)
	if err != nil {
		log.Fatal(err)
	}
	router.PopulateRoutes(discordSession, botPrefix)
	err = discordSession.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bot is running, enjoy your game!")
	select {}
}
