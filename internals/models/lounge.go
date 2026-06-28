package models

import "gorm.io/gorm"

type Lounges struct {
	gorm.Model
	UniversityId string
	Name         string
	Location     string
	Capacity     int
	Amenities    []string
	ImageUrl     string
	IsActive     bool
}
