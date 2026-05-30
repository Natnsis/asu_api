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
	description     string
	DepartmentID    uint

	PrerequisiteCourseIds datatypes.JSONSlice[string]
	StudentTypeIds        datatypes.JSONSlice[string]
	CourseID              uint

	Course      []Course
	Semester    Semester
	Department  Department
	StudentType StudentType
}
