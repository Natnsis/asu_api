package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	UniversityID uint
	Name         string
	Code         string
	HeadId       string
	IsActive     bool
	Programs     datatypes.JSON

	Complaint  []Complaint
	University University
	Course     []Course
	Program    []Program
	User       []User
}
