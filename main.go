package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	nre "github.com/mascanio/regexp-named"
)

var RE_DICE = nre.MustCompile(`(?P<n>\d+)?d(?P<d>\d+)`)
var MAX_DICE = 100
var MAX_DICE_TYPE = 100
var PRE_IS_DICE_ROLL_MAX_LEN = 100

func validDiceChar(c rune) bool {
	return c >= '0' && c <= '9' || c == 'd'
}

func isDiceRoll(s string) bool {
	dFound := false
	diceTypeFound := false
	for i, c := range s {
		switch {
		case i > PRE_IS_DICE_ROLL_MAX_LEN:
			return false
		case !validDiceChar(c):
			return false
		case !dFound:
			switch {
			case c == 'd':
				dFound = true
			}
		case dFound:
			if c == 'd' {
				return false
			}
			diceTypeFound = true
		}
	}
	return dFound && diceTypeFound
}

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func rollDice(n int, d int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		// generate random number between 1 and d
		r[i] = rand.Intn(d) + 1
	}
	return r
}

func sumDiceRolls(r []int) int {
	sum := 0
	for _, v := range r {
		sum += v
	}
	return sum
}

func getNDiceType(s string) (int, int, error) {
	_, mm := RE_DICE.FindStringNamed(s)
	if mm == nil {
		return 0, 0, fmt.Errorf("no dice found")
	}
	nDices, err := strconv.Atoi(mm["n"])
	if err != nil {
		nDices = 1
	}
	diceType, err := strconv.Atoi(mm["d"])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting d to int")
	}
	if nDices > MAX_DICE || diceType > MAX_DICE_TYPE {
		return 0, 0, fmt.Errorf("too many dices or incorrect dice type")
	}
	return nDices, diceType, nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) || !isDiceRoll(m.Content) {
		return
	}

	nDices, diceType, err := getNDiceType(m.Content)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, err.Error(), m.Reference())
		return
	}

	roll := rollDice(nDices, diceType)
	sum := sumDiceRolls(roll)

	response := fmt.Sprintf("Rolling %dd%d: %v\nSum: %d", nDices, diceType, roll, sum)
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
