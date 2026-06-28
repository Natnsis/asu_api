package models

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	UserID        uint
	UniversityID  uint
	Action        string
	EntityType    string
	EntityId      string
	PreviousValue string
	NewValue      string
	IpAddress     string
	UserAgent     string

	User       User
	University []University
}
