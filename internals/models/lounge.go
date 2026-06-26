package models

import "gorm.io/gorm"

type Lounges struct {
	gorm.Model
	universityId string
	name         string
	location     string
	capacity     int
	amenities    []string
	imageUrl     string
	isActive     bool
}
