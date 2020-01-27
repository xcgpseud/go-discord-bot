package message_handlers

import (
	"discord-bot/helpers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/disintegration/imaging"
	"os"
)

func FlipHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	img := helpers.GetImageFromUrl(m.Author.AvatarURL(""))
	imgFlipped := imaging.FlipV(img)

	path := helpers.SaveImage(imgFlipped)
	r := helpers.GetReaderFromImage(path)

	_, err := s.ChannelFileSend(
		m.ChannelID,
		fmt.Sprintf("%s.jpg", m.Author.Username),
		r,
	)
	helpers.LogError(err)

	err = os.Remove(path)
	helpers.LogError(err)
}
