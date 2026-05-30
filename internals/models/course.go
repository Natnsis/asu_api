package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string
	CreditHour   float32
	Description  string
	CurriculumID uint
	Curriculum   Curriculum
}
