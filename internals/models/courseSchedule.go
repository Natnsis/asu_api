package models

import "gorm.io/gorm"

type CourseSchedule struct {
	gorm.Model
	courseId   string
	semesterId string
	dayOfWeek  string
	startTime  string
	endTime    string
	venue      string
	types      string
}
