package models

import (
	"log"

	"UniCore/internals/config"
)

func AutoMigrateModels() {
	err := config.Db.AutoMigrate(
		&University{},
		&Setting{},
		&NotificationTemplate{},
		&Department{},
		&EventCategory{},
		&ComplaintCategory{},
		&ReportCategory{},
		&Lounge{},
		&Gallery{},
		&Cafeteria{},
		&Semester{},
		&User{},
		&Program{},
		&Course{},
		&Event{},
		&Complaint{},
		&Notification{},
		&Audit{},
		&Profile{},
		&Records{},
		&EventRegistration{},
		&GalleryMedia{},
		&LoungeAvailability{},
		&LoungeBooking{},
		&MealPlan{},
		&MealSlot{},
		&MealOverride{},
		&CourseSchedule{},
		&CourseMaterial{},
		&ComplaintResponse{},
		&Report{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
