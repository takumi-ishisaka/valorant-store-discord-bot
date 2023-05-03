package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if mc.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasSuffix(mc.Content, "/store") {
		s.ChannelMessageSend(mc.ChannelID, "valorant store show")
        return
	}
}