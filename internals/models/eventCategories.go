package models

import "gorm.io/gorm"

type EventCategories struct {
	gorm.Model
	Name        string
	Description string
	Event       []Event `gorm:"foreignKey:EventCategoryID"`
}
