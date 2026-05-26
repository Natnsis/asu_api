package models

import "gorm.io/gorm"

type Curriculum struct {
	gorm.Model
	SemesterId            string
	CourseId              string
	IsMandatory           bool
	MinPassingGrade       float32
	CreditHour            int
	description           string
	PrerequisiteCourseIds []string
	StudentTypeIds        []string // regular, summer, half-year, postgrad
}
