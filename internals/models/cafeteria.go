package models

import (
	"time"

	"gorm.io/gorm"
)

type Cafeteria struct {
	gorm.Model
	univeristyId string
	name         string
	location     string
	openTime     time.Time
	closeTime    time.Time
	isActive     bool
}
