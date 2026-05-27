package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load env file")
	}

	DB_url := os.Getenv("LOCAL_DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(DB_url), &gorm.Config{})
	if err != nil {
		log.Fatal("databse is not connected")
	}

	DbConnection = conn
}
