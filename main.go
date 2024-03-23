package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mascanio/disc-dice-go/dice"
	input "github.com/mascanio/disc-dice-go/user-input"
)

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func rollDices(n int, d int) (int, string) {
	sum := 0
	result := "["
	dice := dice.GenericDice{Faces: d}
	for i := 0; i < n; i++ {
		r := dice.Roll()
		sum += r.Sum
		result += r.Result
		if i < n-1 {
			result += ", "
		}
	}
	result += "]"
	return sum, result
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) || !input.IsDiceRoll(m.Content) {
		return
	}

	nDices, diceType, err := input.GetNDiceType(m.Content)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, err.Error(), m.Reference())
		return
	}

	sum, result := rollDices(nDices, diceType)

	response := fmt.Sprintf("Rolling %dd%d: %v\nSum: %d", nDices, diceType, result, sum)
	s.ChannelMessageSendReply(m.ChannelID, response, m.Reference())
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
