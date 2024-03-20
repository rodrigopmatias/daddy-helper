package main

import (
	"log"

	"github.com/rodrigopmatias/daddy-helper/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(
		&models.Message{},
	)
}
