package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name          string
	Status        string
	Code          string
	HeadName      string
	OfficeAddress string

	CollegeID uint

	Instructors datatypes.JSONSlice[string]
	Curriculum  []Curriculum
	User        []User
}
