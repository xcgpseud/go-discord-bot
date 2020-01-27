package helpers

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

type Admin struct {
	ID string `json:"id"`
}

type Admins struct {
	Admins []Admin `json:"admins"`
}

func GetAdmins() (out Admins) {
	f, err := os.Open("config/admins.json")
	LogError(err)

	bVal, err := ioutil.ReadAll(f)
	LogError(err)
	
	err = json.Unmarshal(bVal, &out)
	LogError(err)

	return
}

func IsAdmin(user *discordgo.User) bool {
	admins := GetAdmins().Admins

	for _, v := range admins {
		if v.ID == user.ID {
			return true
		}
	}

	return false
}
