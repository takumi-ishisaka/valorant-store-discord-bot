package main

import (
	"fmt"
	"os"
	"os/signal"
	"golang.org/x/exp/slog"
	"syscall"
	"github.com/takumi-ishisaka/valorant-store-discord-bot/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env")
	logger := slog.New(slog.HandlerOptions{AddSource: true,}.NewJSONHandler(os.Stdout))

	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("ログインに失敗しました")
		logger.Error("", err)
	}

	dg.AddHandler(handlers.OnMessageCreate)

	dg.Open()
	if err != nil {
		fmt.Println("Discordのオープンに失敗しました")
		logger.Error("", err)
	}
	defer dg.Close()

	fmt.Println("Bot is running. Press CTRL-C to exit.")

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot
}