package models

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	UniversityId string
	Name         string
	StartDate    time.Time
	EndDate      time.Time
	IsCurrent    bool
}
