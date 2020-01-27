package message_handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func PingHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
