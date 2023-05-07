package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/takumi-ishisaka/valorant-store-discord-bot/operators"
)

func OnMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if mc.Author.ID == s.State.User.ID {
		return
	}

	// mc.Content is "@clientId message"
	if strings.HasSuffix(mc.Content, "/store") {
		s.ChannelMessageSend(mc.ChannelID, "valorant store show")
        return
	} else if strings.HasSuffix(mc.Content, "/agent") {
		s.ChannelMessageSend(mc.ChannelID, "valorant agent show")
        agents := operators.GetAgents()
		s.ChannelMessageSend(mc.ChannelID, string(agents.Status))
		var count int = 0
		for _, agent := range agents.AgentData {
			s.ChannelMessageSend(mc.ChannelID, string(agent.DisplayName))
            s.ChannelMessageSend(mc.ChannelID, string(agent.DisplayIcon))
			count += 1
			if count == 2 {
				break
			}
        }
		return
	} else {
		s.ChannelMessageSend(mc.ChannelID, "command not found")
		return
	}
}