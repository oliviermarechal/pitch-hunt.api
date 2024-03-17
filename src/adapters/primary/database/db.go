package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			panic("DATABASE_URL environment variable not set")
		}

		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
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

		log.Println("Connected to database")
	})

	return db
}
