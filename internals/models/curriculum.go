package models

import "gorm.io/gorm"

type Curriculum struct {
	gorm.Model
	universityId string
	departmentId string
	programId    string
	lecturerId   string
	name         string
	code         string
	creditHours  int
	descritpion  string
	isElective   bool
}
