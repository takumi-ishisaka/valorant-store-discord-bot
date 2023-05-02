package handlers

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	clientId := os.Getenv("CLIENT_ID")
	u := mc.Author
	if u.ID != clientId {
    	s.ChannelMessageSend(mc.ChannelID, "Hello World!")
	}
}