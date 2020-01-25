package pester

import (
	"discord-bot/database"
	"discord-bot/database/entities"
	"errors"
)

func Create(uidFrom, uidTo, msg string) (entities.Pester, error) {
	db := database.GetDb()

	var existing entities.Pester
	if db.First(&existing, "user_id_to = ?", uidTo).RecordNotFound() {
		pester := entities.Pester{
			UserIdFrom: uidFrom,
			UserIdTo:   uidTo,
			Message:    msg,
		}

		db.Create(&pester)

		return pester, nil
	}

	return entities.Pester{}, errors.New("entity already exists")
}

func Get(uidTo string) (entities.Pester, error) {
	db := database.GetDb()

	var pester entities.Pester
	if db.First(&pester, "user_id_to = ?", uidTo).RecordNotFound() {
		return pester, errors.New("could not find user")
	}

	return pester, nil
}

func Delete(uidTo string) error {
	db := database.GetDb()

	var pester entities.Pester
	if err := db.Delete(&pester, "user_id_to = ?", uidTo).Error; err != nil {
		return err
	}
	return nil
}
