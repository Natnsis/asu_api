package models

import "gorm.io/gorm"

type Curriculum struct {
	gorm.Model
	UniversityId string
	DepartmentId string
	ProgramId    string
	LecturerId   string
	Name         string
	Code         string
	CreditHours  int
	Descritpion  string
	IsElective   bool
}
