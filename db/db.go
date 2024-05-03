package db

import (
	"github.com/vishalpatidar99/Book-Management/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Book{})

	DB = database
}
