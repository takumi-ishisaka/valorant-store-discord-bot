package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/takumi-ishisaka/valorant-store-discord-bot/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// // Discord client new
	// discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	// if err != nil {
	// 	fmt.Println("error creating Discord session,", err)
	// 	return
	// }

	// dg.AddHandler(handlers.MessageCreate)

	// err = dg.Open()
	// if err != nil {
	// 	fmt.Println("error opening connection,", err)
	// 	return
	// }

	// fmt.Println("Bot is now running. Press CTRL-C to exit.")
	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	// dg.Close()

	godotenv.Load(".env")

	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("ログインに失敗しました")
		fmt.Println(err)
	}

	discord.AddHandler(handlers.OnMessageCreate)
	discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	defer discord.Close()

	fmt.Println("bot is running")

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot
}