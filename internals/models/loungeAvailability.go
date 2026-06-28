package models

import (
	"time"

	"gorm.io/gorm"
)

type LoungeAvailability struct {
	gorm.Model
	LoungeID  uint
	DayOfWeek DaysType
	OpenTime  time.Time
	CloseTime time.Time

	Lounge Lounge
}
