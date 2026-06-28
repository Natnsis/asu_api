package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	UniversityID uint
	DepartmentID uint
	ProgramID    string
	Name         string
	Code         string
	CreditHours  int
	Descritpion  string
	IsElective   bool

	CourseSchedule []CourseSchedule
	CourseMaterial []CourseMaterial
	Department     Department
}
