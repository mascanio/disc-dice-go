package messagecontext

import "github.com/bwmarrin/discordgo"

type MessageContext struct {
	Message     string
	Group       string
	SessionName string
	User        string
	Pj          string
}

func New(message discordgo.MessageCreate) MessageContext {
	return MessageContext{
		Message:     message.Content,
		Group:       message.GuildID,
		SessionName: message.ChannelID,
		User:        message.Author.ID,
		Pj:          message.Author.Username,
	}
}
