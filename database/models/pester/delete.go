package pester

import (
	"discord-bot/database"
	"discord-bot/database/entities"
)

func DeleteByUidTo(uidTo string) error {
	db := database.GetDb()

	var pester entities.Pester
	if err := db.Delete(&pester, "user_id_to = ?", uidTo).Error; err != nil {
		return err
	}
	return nil
}
