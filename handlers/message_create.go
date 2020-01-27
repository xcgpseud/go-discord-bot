package handlers

import (
	pester2 "discord-bot/database/models/pester"
	mh "discord-bot/handlers/message_handlers"
	"discord-bot/helpers"
	"github.com/bwmarrin/discordgo"
)

type Route struct {
	Command string
	Handler func(*discordgo.Session, *discordgo.MessageCreate)
}

var routes []Route = []Route{
	{"ping", mh.PingHandler},
	{"pong", mh.PongHandler},
	{"dyl", mh.DylHandler},
	{"flip", mh.FlipHandler},
	{"pester", mh.PesterHandler},
	{"unpester", mh.PesterHandler},
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if helpers.IsCommand(m.Content) {
		command, _ := helpers.GetCommandAndContent(m.Content)
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
