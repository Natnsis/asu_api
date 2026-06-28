package models

import (
	"time"

	"gorm.io/gorm"
)

type MealOverride struct {
	gorm.Model
	CafeteriaId string
	Date        time.Time
	MealPeriod  MealType
	IsClosed    bool
	MealName    string
	Reason      string
	CreatedById time.Time
}
