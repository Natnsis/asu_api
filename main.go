package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("couldnt connect to the db")
	}
	fmt.Println("Successfuly connected to db via GORM")

	err = db.AutoMigrate(&User{})
}
