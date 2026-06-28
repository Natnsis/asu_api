package models

import (
	"time"

	"gorm.io/gorm"
)

type Records struct {
	gorm.Model
	EventID      uint
	StudentId    string
	RegisteredAt time.Time
	Status       string
}
