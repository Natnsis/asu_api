package models

import "gorm.io/gorm"

type EventCategory struct {
	gorm.Model
	universityId string
	name         string
	colorHex     string
}
