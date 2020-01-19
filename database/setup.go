package database

import (
	"discord-bot/database/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetDb() gorm.DB {
	db, err := gorm.Open("sqlite3", "database/vorkath.db")
	if err != nil {
		panic("Error connecting to sqlite DB.")
	}
	return *db
}

func Migrate() {
	db := GetDb()
	defer db.Close()
	db.AutoMigrate(&entities.Pester{})
}
