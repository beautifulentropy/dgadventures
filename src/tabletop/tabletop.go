package tabletop

import (
	"errors"
	"time"

	"github.com/beautifulentropy/dgadventures/src/helpers"
	"github.com/bwmarrin/discordgo"
)

// ActionDefinition ...
type ActionDefinition func(*DiceMat, *discordgo.MessageReaction)

// DiceMat ...
type DiceMat struct {
	Display   *discordgo.MessageEmbed
	Message   *discordgo.Message
	Session   *discordgo.Session
	ChannelID string
	Timeout   time.Duration
	Channel   chan bool
	Actions   map[string]ActionDefinition
	Keys      []string
}

// Start ...
func (d *DiceMat) Start() error {
	if d.Display == nil {
		return errors.New("embed is nil")
	}
	startTime := time.Now()
	dgMessage, err := d.Session.ChannelMessageSendEmbed(d.ChannelID, d.Display)
	if err != nil {
		return err
	}
	d.Message = dgMessage
	for _, v := range d.Keys {
		d.Session.MessageReactionAdd(d.Message.ChannelID, d.Message.ID, v)
	}
	var dgReact *discordgo.MessageReaction
	for {
		if d.Timeout != 0 {
			select {
			case dgReactAdd := <-helpers.AddHandlerForNextReaction(d.Session):
				dgReact = dgReactAdd.MessageReaction
			case <-time.After(startTime.Add(d.Timeout).Sub(time.Now())):
				return nil
			case <-d.Channel:
				return nil
			}
		} else {
			select {
			case dgReactAdd := <-helpers.AddHandlerForNextReaction(d.Session):
				dgReact = dgReactAdd.MessageReaction
			case <-d.Channel:
				return nil
			}
		}
		if (dgReact.MessageID != d.Message.ID) || (d.Session.State.User.ID == dgReact.UserID) {
			continue
		}
		if rollHandler, ok := d.Actions[dgReact.Emoji.Name]; ok {
			rollHandler(d, dgReact)
		}
		time.Sleep(time.Millisecond * 250)
		d.Session.MessageReactionRemove(dgReact.ChannelID, dgReact.MessageID, dgReact.Emoji.Name, dgReact.UserID)
	}
}

// AppendAction ...
func (d *DiceMat) AppendAction(emojiName string, action ActionDefinition) error {
	if _, ok := d.Actions[emojiName]; !ok {
		d.Keys = append(d.Keys, emojiName)
		d.Actions[emojiName] = action
	}
	if d.Message != nil {
		return d.Session.MessageReactionAdd(d.Message.ChannelID, d.Message.ID, emojiName)
	}
	return nil
}

// DisplayUpdate ...
func (d *DiceMat) DisplayUpdate(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	if d.Message == nil {
		return nil, errors.New("message is nil")
	}
	return d.Session.ChannelMessageEditEmbed(d.ChannelID, d.Message.ID, embed)
}
