package models

import "gorm.io/gorm"

type EventType struct {
	gorm.Model
	Name        string
	Description string
}
