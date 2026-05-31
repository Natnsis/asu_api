package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title              string
	Subtitle           string
	Description        string
	ThumbnailUrl       string
	ApprovedByID       uint
	MaxParticipants    int
	AttendanceRequired bool
	UserID             uint
	EventTypeID        uint
	UniversityID       uint

	EventCategoryID uint

	ApprovedBy    User `gorm:"foreignKey:ApprovedByID"`
	University    University
	User          User `gorm:"foreignKey:UserID"`
	EventCategory EventCategories
	EventType     EventType
}
