package models

import (
	"time"

	"gorm.io/gorm"
)

type Lounge struct {
	gorm.Model
	Name                 string
	Description          string
	LoungeTypeIds        []string
	ImageUrl             string
	MapUrl               string
	Capacity             int
	HasWifi              bool
	HasPowerOutlet       bool
	HasAirConditioning   bool
	HasProjector         bool
	HasWhiteBoard        bool
	NoiseLevel           int
	Amenities            []string
	OpeningTime          time.Time
	ClosingTime          time.Time
	AvailbleDays         []string
	AllowedDepartmentIds []string
	ManagedBy            string
	Status               string
	LoungeTypesID        uint
	LoungeTypes          LoungeTypes
}
