package handlers

import (
	pester2 "discord-bot/database/models/pester"
	"discord-bot/helpers"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

type Route struct {
	Command string
	Handler func(*discordgo.Session, *discordgo.MessageCreate)
}

var routes []Route = []Route{
	{"ping", PingHandler},
	{"pong", PongHandler},
	{"dyl", DylHandler},
	{"flip", FlipHandler},
	{"pester", PesterHandler},
	{"unpester", PesterHandler},
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if isCommand(m.Content) {
		command, _ := GetCommandAndContent(m.Content)
		for i := range routes {
			if routes[i].Command == command {
				routes[i].Handler(s, m)
			}
		}
	}

	pester, err := pester2.GetByUidTo(m.Author.ID)
	if err != nil {
		return
	}
	_, err = s.ChannelMessageSend(m.ChannelID, pester.Message)
	helpers.LogError(err)
}

func isCommand(s string) bool {
	prefix := os.Getenv("PREFIX")
	return len(s) > len(prefix) && s[0:len(prefix)] == prefix
}

func GetCommandAndContent(s string) (command string, content string) {
	parts := strings.Split(s, " ")
	command = parts[0][len(os.Getenv("PREFIX")):]
	if len(parts) < 2 {
		return
	}
	content = strings.Join(parts[1:], " ")
	return
}
