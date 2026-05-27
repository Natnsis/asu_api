package models

import (
	"log"

	"UniCore/internals/db"
)

func AutoMigrateModels() {
	err := db.DbConnection.AutoMigrate(
		&User{},
		&Profile{},
		&Role{},
	)
	if err != nil {
		log.Fatal("unable to migrate db", err)
	}
}
