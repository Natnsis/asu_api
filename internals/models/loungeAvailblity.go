package models

import (
	"time"

	"gorm.io/gorm"
)

type LoungeAvailablity struct {
	gorm.Model
	loungeId  string
	dayOfWeek string
	openTime  time.Time
	closeTime time.Time
}
