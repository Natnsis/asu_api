package main

import (
	"fmt"
	"log"
	"net/http"

	"UniCore/internals/db"
	"UniCore/internals/models"
	"UniCore/internals/routes"
)

func main() {
	db.DbConnection()
	models.MigrateModels()
	router := routes.AuthRoutes()
	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
