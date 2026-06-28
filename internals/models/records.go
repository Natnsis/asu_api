package models

import (
	"time"

	"gorm.io/gorm"
)

type Records struct {
	gorm.Model
	UniversityID uint
	EventID      uint
	StudentId    string
	RegisteredAt time.Time
	Status       string
}
