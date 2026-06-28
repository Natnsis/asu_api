package models

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	UniversityID uint
	Name         string
	StartDate    time.Time
	EndDate      time.Time
	IsCurrent    bool

	CourseSchedule []CourseSchedule
}
