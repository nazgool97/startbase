package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nazgool97/startbase/internal/models"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN is not set in environment")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = database
	fmt.Println("✅ Connected to PostgreSQL")

	DB.AutoMigrate(&models.User{})
	fmt.Println("📦 User table migrated")
}