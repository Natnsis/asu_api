package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_url string

func GetDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load env file")
	}

	DB_url = os.Getenv("DATABASE_URL")
}
