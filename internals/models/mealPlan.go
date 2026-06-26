package models

import (
	"time"

	"gorm.io/gorm"
)

type MealPlan struct {
	gorm.Model
	cafeteriaId   string
	weekStartDate time.Time
	weekEndDate   time.Time
	createdById   string
	status        string
}
