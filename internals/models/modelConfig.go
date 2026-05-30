package models

import (
	"log"

	"UniCore/internals/db"
)

func AutoMigrateModels() {
	err := db.DbConnection.AutoMigrate(
		Album{},
		Cafeteria{},
		College{},
		Course{},
		Curriculum{},
		Department{},
		Event{},
		EventCategories{},
		EventType{},
		Gallery{},
		GalleryCategories{},
		Lounge{},
		LoungeType{},
		RecentActivity{},
		Role{},
		Semester{},
		SocialLink{},
		StudentType{},
		University{},
		UniversityType{},
		&User{},
		&Profile{},
		&Role{},
	)
	if err != nil {
		log.Fatal("unable to migrate db", err)
	}
}
