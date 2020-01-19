package handlers

import (
	pester2 "discord-bot/database/repositories/pester"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func PesterHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd, content := GetCommandAndContent(m.Content)
	if cmd == "pester" {
		mentions := m.Mentions
		if len(mentions) < 1 {
			s.ChannelMessageSend(m.ChannelID, "You must mention somebody to pester...")
		}
		uidTo, _ := strconv.ParseInt(mentions[0].ID, 10, 32)
		uidFrom, _ := strconv.ParseInt(m.Author.ID, 10, 32)
		_, err := pester2.Create(uidFrom, uidTo, content)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "They are already being pestered.")
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Now pestering %s", mentions[0].Username))
	}
}
