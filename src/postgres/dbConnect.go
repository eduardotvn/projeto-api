package postgres

import (
	"log"
	"os"

	"github.com/eduardotvn/projeto-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	var err error
	dsn := os.Getenv("DB_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})

}
