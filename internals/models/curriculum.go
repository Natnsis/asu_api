package models

import "gorm.io/gorm"

type Curriculum struct {
	gorm.Model
	SemesterId      string
	CourseId        string
	IsMandatory     bool
	MinPassingGrade float32
	CreditHour      int
	description     string
	DepartmentID    uint

	PrerequisiteCourseIds []string
	Department            Department
	StudentTypeIds        []string // regular, summer, half-year, postgrad
}
