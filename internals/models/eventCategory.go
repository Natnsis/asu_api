package models

import "gorm.io/gorm"

type EventCategory struct {
	gorm.Model
	UniversityId string
	Name         string
	ColorHex     string
}
