package router

import (
	"fmt"

	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/beautifulentropy/dgadventures/src/game"
	"github.com/beautifulentropy/dgadventures/src/helpers"
	"github.com/bwmarrin/discordgo"
)

// PopulateRoutes ...
func PopulateRoutes(session *discordgo.Session, botPrefix *string) {
	router := exrouter.New()
	router.On("turn", func(ctx *exrouter.Context) { game.LoadCharacter(session, ctx.Msg) }).Desc("loads your character sheet and starts your turn")
	router.On("roll", func(ctx *exrouter.Context) {
		ctx.Reply(fmt.Sprintf("```%s```", helpers.GetResultForRoll(ctx.Args[1])))
	}).Desc("given a roll formula (e.g. 5d6) it gives you a result")
	session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		router.FindAndExecute(session, *botPrefix, session.State.User.ID, message.Message)
	})
	router.Default = router.On("help", func(ctx *exrouter.Context) {
		var reply string
		for _, route := range router.Routes {
			reply += fmt.Sprintf("%s: %s\n", route.Name, route.Description)
		}
		ctx.Reply(fmt.Sprintf("```%s```", reply))
	}).Desc("prints this help menu")
}
