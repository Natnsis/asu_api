package lounges

import "gorm.io/gorm"

type Lounge struct {
	gorm.Model

	Name         string `json:"official_name"`
	BuildingName string `json:"building_name"`
	MaxCapacity  int    `json:"max_capacity"`
	IsQuiet      bool   `json:"is_queit"`
}
