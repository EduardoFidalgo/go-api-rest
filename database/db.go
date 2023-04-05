package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSL_MODULE")

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", dbUser, dbName, dbPass, dbHost, dbSSL)
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic(err.Error())
	}
}
