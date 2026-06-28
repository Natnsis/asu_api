package models

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	ActorId       string
	UniversityId  string
	Action        string
	EntityType    string
	EntityId      string
	PreviousValue string
	NewValue      string
	IpAddress     string
	UserAgent     string
}
