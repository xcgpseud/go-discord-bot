package message_handlers

import (
	"discord-bot/database/models/pester"
	"discord-bot/helpers"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func PesterHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if !helpers.IsAdmin(m.Author) {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		helpers.LogError(err)
		return
	}

	cmd, content := helpers.GetCommandAndContent(m.Content)
	if cmd == "pester" {
		mentions := m.Mentions
		if len(mentions) < 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, "You must mention somebody to pester...")
			helpers.LogError(err)
		}
		_, err := pester.Create(m.Author.ID, mentions[0].ID, content)
		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, "They are already being pestered.")
			helpers.LogError(err)
			return
		}
		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Now pestering %s", mentions[0].Username))
		helpers.LogError(err)
	}

	if cmd == "unpester" {
		mentions := m.Mentions
		if len(mentions) < 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, "You need to mention somebody to stop pestering.")
			helpers.LogError(err)
		}
		err := pester.DeleteByUidTo(mentions[0].ID)
		helpers.LogError(err)
	}
}
