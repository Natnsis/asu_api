package db

import (
	"log"

	"UniCore/internals/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func ConnectDb() {
	config.GetDotEnv()
	conn, err := gorm.Open(postgres.Open(config.DB_url), &gorm.Config{})
	if err != nil {
		log.Fatal("databse is not connected")
	}

	DbConnection = conn
}
