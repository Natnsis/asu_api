package models

import "gorm.io/gorm"

type StudentType struct {
	gorm.Model
	Name        string
	Description string
}
