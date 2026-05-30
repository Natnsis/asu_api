package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Curriculum struct {
	gorm.Model
	IsMandatory     bool
	MinPassingGrade float32
	description     string

	PrerequisiteCourseIds datatypes.JSONSlice[string]
	StudentTypeIds        datatypes.JSONSlice[string]
	CourseID              uint
	SemesterId            uint
	CourseId              uint
	DepartmentID          uint

	Course      []Course
	Semester    Semester
	Department  Department
	StudentType StudentType
}
