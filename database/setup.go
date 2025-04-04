package database

import (
	"chill-pi/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Klarte ikke koble til database:", err)
	}
	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("❌ Migrering feilet:", err)
	}

	DB = db

}
