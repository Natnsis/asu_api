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
	RoleID             datatypes.JSONSlice[string]
	ApprovedByID       uint
	MaxParticipants    int
	AttendanceRequired bool
	UserID             uint
	EventTypeID        uint
	EventCategoriesID  datatypes.JSONSlice[string]

	ApprovedBy      User `gorm:"foreignKey:ApprovedByID"`
	User            User `gorm:"foreignKey:UserID"`
	EventCategories EventCategories
	EventType       EventType
	Role            Role
}
