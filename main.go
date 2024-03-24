package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mascanio/disc-dice-go/multidice"
	input "github.com/mascanio/disc-dice-go/user-input"
)

type Messager interface {
	Message() string
}

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) || !input.IsDiceRoll(m.Content) {
		return
	}

	nDices, diceType, err := input.GetNDicesAndFaces(m.Content)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, err.Error(), m.Reference())
		return
	}

	multidice := multidice.MultiDice{Faces: diceType, Dices: nDices}
	s.ChannelMessageSendReply(m.ChannelID, multidice.Roll().Message(), m.Reference())
}

func main() {
	token := os.Getenv("TOKEN")

	if token == "" {
		fmt.Println("No token provided. Please set TOKEN environment variable.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents |= discordgo.IntentMessageContent

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}
	defer dg.Close()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
