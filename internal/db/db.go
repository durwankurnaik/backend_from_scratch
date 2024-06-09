package db

import (
	"log"

	"github.com/durwankurnaik/glovs/internal/config"
	"github.com/durwankurnaik/glovs/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dbFile := config.Config("DB_FILE")

	if dbFile == "" {
		dbFile = "glovs.db" // default value if not set
	}

	DB, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatically migrate schema  to keep it up to date
	DB.AutoMigrate(&models.Article{})
}