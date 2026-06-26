package models

import (
	"time"

	"gorm.io/gorm"
)

type LoungeBooking struct {
	gorm.Model
	loungeId   string
	bookedById string
	date       time.Time
	startTime  time.Time
	endTime    time.Time
	attendees  int
	purpose    string
	status     string
}
