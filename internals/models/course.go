package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	UniversityID string
	DepartmentID uint
	ProgramId    string
	LecturerId   string
	Name         string
	Code         string
	CreditHours  int
	Descritpion  string
	IsElective   bool

	Department Department
}
