package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load doenv")
	}

	db_variable := os.Getenv("DATABASE_URL")
	connect, err := gorm.Open(postgres.Open(db_variable), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to db")
	}

	fmt.Println("successfuly connected to db")

	err = connect.AutoMigrate(&User{})
}
