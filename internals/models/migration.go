package models

import (
	"log"

	"UniCore/internals/db"
)

func MigrateModels() {
	err := db.Db.AutoMigrate(
		User{},
	)
	if err != nil {
		log.Fatal("table unable to migrate with err: ", err)
		return
	}
	log.Println("db migrated")
}
