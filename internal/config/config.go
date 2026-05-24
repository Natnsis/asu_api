package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DbUrl string

func GetEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load the env file")
	}

	DbUrl = os.Getenv("DATABASE_URL")
}
