package message_handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func PongHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Ping!")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
