package models

import (
	"time"

	"gorm.io/gorm"
)

type MealOverride struct {
	gorm.Model
	cafeteriaId string
	date        time.Time
	mealPeriod  string
	isClosed    bool
	mealName    string
	reason      string
	createdById time.Time
}
