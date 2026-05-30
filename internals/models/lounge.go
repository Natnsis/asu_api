package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Lounge struct {
	gorm.Model
	Name                string
	Description         string
	ImageUrl            string
	MapUrl              string
	Capacity            int
	HasWifi             bool
	HasPowerOutlet      bool
	HasAirConditioning  bool
	HasProjector        bool
	HasWhiteBoard       bool
	NoiseLevel          int
	Amenities           datatypes.JSONSlice[string]
	OpeningTime         time.Time
	ClosingTime         time.Time
	AvailbleDays        datatypes.JSONSlice[string]
	AllowedDepartmentID datatypes.JSONSlice[string]
	LoungeTypesID       datatypes.JSONSlice[string]

	ManagedBy          string
	Status             string
	Manager            []Role
	AllowedDepartments []Department `gorm:"foreignKey:AllowedDepartmentID"`
	LoungeType         LoungeType
}
