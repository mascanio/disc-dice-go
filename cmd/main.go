package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mascanio/disc-dice-go/internal/parser"
)

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) {
		return
	}
	go func() {
		defer func(timeStart time.Time) {
			slog.Debug("messageCreate", "time_taken_us", time.Since(timeStart).Microseconds())
		}(time.Now())
		messager := parser.InputToMessager(m.Content)
		if messager != nil {
			_, err := s.ChannelMessageSendReply(m.ChannelID, messager.Message(), m.Reference())
			if err != nil {
				slog.Error("error sending reply", slog.Any("err", err))
			}
		}
	}()
}

func main() {
	setupLogger()

	token := os.Getenv("TOKEN")
	if token == "" {
		slog.Error("No token provided. Please set TOKEN environment variable.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		slog.Error("Error creating Discord session.", slog.Any("err", err))
		return
	}
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		slog.Error("Error opening connection", slog.Any("err", err))
		return
	}
	defer dg.Close()

	slog.Info("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	defer close(sc)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func setupLogger() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "INFO"
	}

	var logLevel slog.Level

	switch logLevelStr {
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "INFO":
		logLevel = slog.LevelInfo
	case "WARN":
		logLevel = slog.LevelWarn
	case "ERROR":
		logLevel = slog.LevelError
	default:
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}))
		logger.Error("Unrecognized LOG_LEVEL", "inputLogLevel", logLevelStr)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)
}
