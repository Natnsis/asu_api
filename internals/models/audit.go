package models

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	actorId       string
	universityId  string
	action        string
	entityType    string
	entityId      string
	previousValue string
	newValue      string
	ipAddress     string
	userAgent     string
}
