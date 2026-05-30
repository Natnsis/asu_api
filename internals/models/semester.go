package models

import "gorm.io/gorm"

type Semester struct {
	gorm.Model
	Year       int
	Semester   int
	Curriculum []Curriculum
}
