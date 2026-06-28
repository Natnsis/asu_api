package models

import "gorm.io/gorm"

type EventCategory struct {
	gorm.Model
	UniversityID string
	Name         string
	ColorHex     string

	University University
}
