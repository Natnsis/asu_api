package models

import "gorm.io/gorm"

type LoungeType struct {
	gorm.Model
	Name        string
	Description string
	Lounge      []Lounge
}
