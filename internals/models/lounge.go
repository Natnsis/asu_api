package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Lounge struct {
	gorm.Model
	UniversityID uint
	Name         string
	Location     string
	Capacity     int
	Amenities    datatypes.JSON
	ImageUrl     string
	IsActive     bool

	University         University
	LoungeAvailability []LoungeAvailability
	LoungeBooking      []LoungeBooking
}
