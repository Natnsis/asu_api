package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error, cant get dontenv file")
	}

	db_url := os.Getenv("DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("couldnt connect to db")
	}

	fmt.Println("connected to db")

	DB = conn
}
