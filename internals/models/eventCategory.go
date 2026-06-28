package models

import "gorm.io/gorm"

type EventCategory struct {
	gorm.Model
	UniversityID uint
	Name         string
	ColorHex     string

	University University
	Events     []Event
}
