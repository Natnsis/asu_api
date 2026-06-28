package models

import (
	"time"

	"gorm.io/gorm"
)

type LoungeAvailablity struct {
	gorm.Model
	LoungeId  string
	DayOfWeek DaysType
	OpenTime  time.Time
	CloseTime time.Time
}
