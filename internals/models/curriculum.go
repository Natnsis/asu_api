package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Curriculum struct {
	gorm.Model
	SemesterID      uint
	IsMandatory     bool
	MinPassingGrade float32
	Description     string
	DepartmentID    uint

	PrerequisiteCourseIds datatypes.JSONSlice[string]
	StudentTypeID uint

	Course     []Course
	Semester   Semester
	Department Department
	StudentType StudentType
}
