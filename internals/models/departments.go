package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name          string
	Status        string
	Code          string
	HeadName      string
	OfficeAddress string

	CollegeId string

	Instructors []string
	Department  []Curriculum
	User        []User
}
