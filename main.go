package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mascanio/disc-dice-go/parser"
)

type Messager interface {
	Message() string
}

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) {
		return
	}
	go func() {
		defer func(timeStart time.Time) {
			fmt.Println("Time elapsed: ", time.Since(timeStart))
		}(time.Now())
		messager := parser.InputToMessager(m.Content)
		if messager != nil {
			s.ChannelMessageSendReply(m.ChannelID, messager.Message(), m.Reference())
		}
	}()
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

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

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
