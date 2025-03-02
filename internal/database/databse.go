package database

import (
	"fmt"
	"log"

	"chat-app-golang/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// Auto-migrate database
	DB.AutoMigrate(&models.Message{}, &models.User{})
	fmt.Println("Database SQLite siap digunakan!")
}
