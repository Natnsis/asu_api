package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

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

	EventCategoriesID datatypes.JSONSlice[string]
	RoleID            datatypes.JSONSlice[string]

	ApprovedBy      User `gorm:"foreignKey:ApprovedByID"`
	University      University
	User            User `gorm:"foreignKey:UserID"`
	EventCategories EventCategories
	EventType       EventType
	Role            Role
}
