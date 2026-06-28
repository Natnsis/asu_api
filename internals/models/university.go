package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	name         string `gorm:"not null"`
	slug         string `gorm:"uniqueIndex"`
	logoUrl      string
	address      string
	contactEmail string `gorm:"not null"`
	contactPhone string
	isActive     bool         `gorm:"not null"`
	Department   []Department // has many
	User         []User
}
