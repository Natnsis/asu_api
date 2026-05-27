package models

import (
	"log"

	"UniCore/internals/db"
)

func AutoMigrateModels() {
	err := db.DbConnection.AutoMigrate(
		&Album{},
		&Cafeteria{},
		&College{},
		&Curriculum{},
		&Department{},
		&Event{},
		&EventCategories{},
		&EventType{},
		&Gallery{},
		&GalleryCategories{},
		&Lounge{},
		&LoungeTypes{},
		&Profile{},
		&RecentActivity{},
		&StudentType{},
		&Universities{},
		&User{},
	)
	if err != nil {
		log.Fatal("unable to migrate db", err)
	}
}
