package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name          string
	Code          string
	CollegeId     string
	HeadName      string
	OfficeAddress string
	Instructors   []string
	Status        string
}
