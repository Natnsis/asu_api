package models

import (
	"time"

	"gorm.io/gorm"
)

type MealOverride struct {
	gorm.Model
	CafeteriaID uint
	Date        time.Time
	MealPeriod  MealType
	IsClosed    bool
	MealName    string
	Reason      string
	CreatedById uint

	Cafeteria Cafeteria
}
