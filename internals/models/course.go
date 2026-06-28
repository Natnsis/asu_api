package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	UniversityID uint
	DepartmentID uint
	ProgramID    uint
	Name         string
	Code         string
	CreditHours  int
	Description  string
	IsElective   bool

	CourseSchedule []CourseSchedule
	CourseMaterial []CourseMaterial
	Department     Department
}
