package main

import (
	"unicore/config"
)

func main() {
	config.ConnectDb()

	db := config.DB
	db.AutoMigrate()
}
