package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	universityID uint
	Name         string
	Code         string
	HeadId       string
	IsActive     bool
	Programs     []string

	Complaint  Complaint
	University University // one
	Course     []Course
	Program    []Program
	User       User
}
