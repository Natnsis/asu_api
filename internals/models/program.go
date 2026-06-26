package models

import "gorm.io/gorm"

type Program struct {
	gorm.Model
	universityId  string
	departmentId  string
	name          string
	code          string
	durationYears int
	degreeLevel   string
}
