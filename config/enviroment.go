package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	log.Println("Environment detected")
}

func buildDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	return fmt.Sprintf(dsn, host, user, password, dbname, port)
}
