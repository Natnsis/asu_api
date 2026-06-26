package models

import (
	"time"

	"gorm.io/gorm"
)

type Records struct {
	gorm.Model
	eventId      string
	studentId    string
	registeredAt time.Time
	status       string
}
