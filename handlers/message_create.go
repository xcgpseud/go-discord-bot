package handlers

import (
	"discord-bot/handlers/command_handlers"
	"github.com/bwmarrin/discordgo"
	"os"
)

type Route struct {
	Command string
	Handler func(*discordgo.Session, *discordgo.MessageCreate)
}

var routes []Route = []Route{
	{
		"ping",
		command_handlers.PingHandler,
	}, {
		"pong",
		command_handlers.PongHandler,
	},
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix, command := GetParts(m.Content)

	if prefix == os.Getenv("PREFIX") {
		for i := range routes {
			if routes[i].Command == command {
				routes[i].Handler(s, m)
			}
		}
	}
}
