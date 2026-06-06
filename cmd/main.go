package main

import (
	"log"
	"net/http"

	"UniCore/internals/db"
	"UniCore/internals/models"
	"UniCore/internals/routes"
)

func main() {
	db.ConnectDb()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	router.Mux

	log.Println("server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
