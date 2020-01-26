package pester

import (
	"discord-bot/database"
	"discord-bot/database/entities"
	"errors"
)

func GetByUidTo(uidTo string) (entities.Pester, error) {
	db := database.GetDb()

	var pester entities.Pester
	if db.First(&pester, "user_id_to = ?", uidTo).RecordNotFound() {
		return pester, errors.New("could not find user")
	}

	return pester, nil
}
