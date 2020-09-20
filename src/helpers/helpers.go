package helpers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/justinian/dice"
)

// AddHandlerForNextReaction ...
func AddHandlerForNextReaction(session *discordgo.Session) chan *discordgo.MessageReactionAdd {
	out := make(chan *discordgo.MessageReactionAdd)
	session.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageReactionAdd) {
		out <- e
	})
	return out
}

// GetResultForRoll ...
func GetResultForRoll(rollFormula string) string {
	if rollFormula == "" {
		return ""
	} else {
		result, _, _ := dice.Roll(rollFormula)
		res := result.String()
		beg := strings.Index(res, "[")
		end := strings.Index(res, "]")
		return res[beg+1 : end]
	}
}
