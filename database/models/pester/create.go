package pester

import (
	"discord-bot/database"
	"discord-bot/database/entities"
	"discord-bot/helpers"
	"errors"
)

func Create(uidFrom, uidTo, msg string) (entities.Pester, error) {
	db := database.GetDb()

	entity, err := GetByUidTo(uidTo)

	if err == nil {
		helpers.LogError(err)
		return entity, errors.New("entity already exists")
	}

	p := entities.Pester{
		UserIdFrom: uidFrom,
		UserIdTo:   uidTo,
		Message:    msg,
	}

	db.Create(&p)

	return p, nil
}
