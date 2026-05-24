package db

import (
	"fmt"
	"log"

	"unicore/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConnection() {
	config.GetEnvs()
	conn, err := gorm.Open(postgres.Open(config.DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to the db")
	}
	Db = conn

	fmt.Println("connected to Db")
}
