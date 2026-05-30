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

	CollegeId string

	Instructors datatypes.JSONSlice[string]
	Department  []Curriculum
	User        []User
}
