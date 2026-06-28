package models

import "gorm.io/gorm"

type Lounge struct {
	gorm.Model
	UniversityID uint
	Name         string
	Location     string
	Capacity     int
	Amenities    []string
	ImageUrl     string
	IsActive     bool

	LoungeAvailablity []LoungeAvailablity
	LoungeBooking     []LoungeBooking
}
