package models

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	UserID        string
	UniversityID  string
	Action        string
	EntityType    string
	EntityId      string
	PreviousValue string
	NewValue      string
	IpAddress     string
	UserAgent     string

	User User
}
