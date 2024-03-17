package main

import (
	"fmt"
	"log"
	"os"

	models "pitch-hunt/src/hexagon/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get *sql.DB: %v", err))
	}

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("failed to ping database: %v", err))
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Pitch{})
}
