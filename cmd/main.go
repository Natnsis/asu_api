package main

import (
	"fmt"

	"UniCore/internals/config"
	"UniCore/internals/models"
)

func main() {
	config.DbConnection()
	models.AutoMigrateModels()
	fmt.Println("server is running on port 8080")
}
