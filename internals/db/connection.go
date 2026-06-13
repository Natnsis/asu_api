package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConnection() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load env file")
		return
	}
	log.Println("env file loaded!")

	// get url
	database_url := os.Getenv("LOCAL_DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(database_url), &gorm.Config{})
	if err != nil {
		log.Fatal("connection failed")
	}

	// showing connection
	Db = conn
}
