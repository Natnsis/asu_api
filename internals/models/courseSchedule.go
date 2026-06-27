package models

import "gorm.io/gorm"

// custom types
type (
	DayOfWeek string
	Type      string
)

// custom enum
const (
	DayMonday    DayOfWeek = "monday"
	DayTuesday   DayOfWeek = "tuesday"
	DayWednesday DayOfWeek = "wednesday"
	DayThursday  DayOfWeek = "thursday"
	DayFriday    DayOfWeek = "friday"
	DaySaturday  DayOfWeek = "saturday"
	DaySunday    DayOfWeek = "sunday"
)

const (
	CourseLecture  Type = "lecture"
	CourseLab      Type = "lab"
	CourseTutorial Type = "tutorial"
	CourseExam     Type = "exam"
)

type CourseSchedule struct {
	gorm.Model
	CourseId   string
	SemesterId string
	DayOfWeek  DayOfWeek `gorm:"default:'monday'"`
	StartTime  string
	EndTime    string
	Venue      string
	Types      Type `gorm:"default:'lecture'"`
}
