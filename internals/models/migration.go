package models

import "UniCore/internals/config"

func AutoMigrateModels() {
	config.Db.AutoMigrate(
		&User{},
		&Audit{},
		&Cafeteria{},
		&ComplaintCategory{},
		&ComplaintResponse{},
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
		&LoungeAvailablity{},
		&LoungeBooking{},
		&MealOverride{},
		&MealPlan{},
		&MealSlot{},
		&Notification{},
		&NotificationTemplete{},
		&Profile{},
		&Program{},
		&Records{},
		&Report{},
		&ReportCategory{},
		&Semester{},
		&Setting{},
		&University{},
		&User{},
	)
}
