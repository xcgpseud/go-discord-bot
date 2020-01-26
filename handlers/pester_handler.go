package handlers

import (
	"discord-bot/database/models/pester"
	"discord-bot/helpers"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func PesterHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd, content := GetCommandAndContent(m.Content)
	if cmd == "pester" {
		mentions := m.Mentions
		if len(mentions) < 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You must mention somebody to pester...")
		}
		_, err := pester.Create(m.Author.ID, mentions[0].ID, content)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "They are already being pestered.")
			return
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Now pestering %s", mentions[0].Username))
	}

	if cmd == "unpester" {
		mentions := m.Mentions
		if len(mentions) < 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You need to mention somebody to stop pestering.")
		}
		err := pester.DeleteByUidTo(mentions[0].ID)
		helpers.LogError(err)
	}
}
