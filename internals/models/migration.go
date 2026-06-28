package models

import (
	"log"

	"UniCore/internals/config"
)

func AutoMigrateModels() {
	err := config.Db.AutoMigrate(
		&User{},
		&Audit{},
		&Cafeteria{},
		&ComplaintCategory{},
		&ComplaintResponse{},
		&Complaint{},
		&Course{},
		&CourseMaterial{},
		&CourseSchedule{},
		&Department{},
		&Event{},
		&EventCategory{},
		&EventRegistration{},
		&Gallery{},
		&GalleryMedia{},
		&Lounge{},
		&LoungeAvailability{},
		&LoungeBooking{},
		&MealOverride{},
		&MealPlan{},
		&MealSlot{},
		&Notification{},
		&NotificationTemplate{},
		&Profile{},
		&Program{},
		&Records{},
		&Report{},
		&ReportCategory{},
		&Semester{},
		&Setting{},
		&University{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
