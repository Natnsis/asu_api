package models

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	universityId string
	name         string
	startDate    time.Time
	endDate      time.Time
	isCurrent    bool
}
