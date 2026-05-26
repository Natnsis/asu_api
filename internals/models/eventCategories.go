package models

import "gorm.io/gorm"

type EventCategories struct {
	gorm.Model
	Name         string
	UniversityId string
	Description  string
}
