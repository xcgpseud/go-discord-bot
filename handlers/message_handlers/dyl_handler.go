package message_handlers

import (
	"discord-bot/helpers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func DylHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, subject := helpers.GetCommandAndContent(m.Content)
	mention := m.Author.Mention()
	_, err := s.ChannelMessageSend(
		m.ChannelID,
		fmt.Sprintf("Hey, %s ... DO YOU LIKE %s?", mention, strings.ToUpper(subject)),
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
