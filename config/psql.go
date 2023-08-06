package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	loadEnv()
	dsn := buildDSN()
	// dsn := "host=localhost user=postgres password=1234 dbname=practicedb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	log.Println("Conexión establecida con Postgress")
}
