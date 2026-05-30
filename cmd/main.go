package main

import (
	"fmt"

	"UniCore/internals/db"
	"UniCore/internals/models"
)

func main() {
	db.ConnectDb()
	models.AutoMigrateModels()
	fmt.Println("server is running...")
}
