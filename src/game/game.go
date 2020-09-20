package game

import (
	"fmt"
	"log"
	"time"

	"github.com/beautifulentropy/dgadventures/src/helpers"
	"github.com/beautifulentropy/dgadventures/src/tabletop"
	"github.com/bwmarrin/discordgo"
)

type turn struct {
	DiceMat *tabletop.DiceMat
	Session *discordgo.Session
}

func (t *turn) start() error {
	defer func() { t.Session.ChannelMessageDelete(t.DiceMat.Message.ChannelID, t.DiceMat.Message.ID) }()
	return t.DiceMat.Start()
}

// LoadCharacter ...
func LoadCharacter(session *discordgo.Session, message *discordgo.Message) {
	turn := &turn{
		Session: session,
		DiceMat: &tabletop.DiceMat{
			ChannelID: message.ChannelID,
			Session:   session,
			Keys:      []string{},
			Actions:   map[string]tabletop.ActionDefinition{},
			Channel:   make(chan bool),
			Display: &discordgo.MessageEmbed{
				Title:       "Welcome, @" + message.Author.String(),
				Description: "Click 📖 below to see your character sheet"},
		},
	}
	turn.DiceMat.AppendAction(characterSheet(message.Author.String()))
	turn.DiceMat.AppendAction(roll("💪", "Strength", "*Result:*", "4d6"))
	turn.DiceMat.AppendAction(roll("🧠", "Wits", "*Result:*", "10d6"))
	turn.DiceMat.AppendAction(endTurn(turn))

	turn.DiceMat.Timeout = time.Minute * 5
	err := turn.start()
	if err != nil {
		log.Fatal(err)
	}
}

func characterSheet(author string) (handler string, actionHandler tabletop.ActionDefinition) {
	return "📖", func(d *tabletop.DiceMat, r *discordgo.MessageReaction) {
		d.DisplayUpdate(&discordgo.MessageEmbed{
			Title:       "Character Sheet",
			Description: fmt.Sprintf("%s, your sheet will be here\n", author)})
	}
}

func roll(emoji, title, description, rollFormula string) (handler string, actionHandler tabletop.ActionDefinition) {
	return emoji, func(d *tabletop.DiceMat, r *discordgo.MessageReaction) {
		d.DisplayUpdate(&discordgo.MessageEmbed{
			Title:       title,
			Description: fmt.Sprintf("%s ```%s```\n", description, helpers.GetResultForRoll(rollFormula))})
	}
}

func endTurn(turn *turn) (handler string, actionHandler tabletop.ActionDefinition) {
	return "🛑", func(d *tabletop.DiceMat, r *discordgo.MessageReaction) {
		close(turn.DiceMat.Channel)
	}
}
