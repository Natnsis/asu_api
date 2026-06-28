package models

import (
	"time"

	"gorm.io/gorm"
)

type Cafeteria struct {
	gorm.Model
	UniversityID uint
	Name         string
	Location     string
	OpenTime     time.Time
	CloseTime    time.Time
	IsActive     bool

	MealPlan     []MealPlan
	MealOverride []MealOverride
}
