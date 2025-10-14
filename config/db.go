package config

import (
	"fmt"
	"log"
	"os"
	"todo-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "db"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "secret"),
		getEnv("DB_NAME", "todo_db"),
		getEnv("DB_PORT", "5432"),
	)

	var db *gorm.DB
	var err error

	if os.Getenv("APP_ENV") == "test" {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

		db.AutoMigrate(&models.Todo{})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	log.Println("Database connection established")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
