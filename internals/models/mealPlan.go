package models

import (
	"time"

	"gorm.io/gorm"
)

// custom types
type MealType string

// custom enums
const (
	DraftMeal     MealType = "draft"
	PublishedMeal MealType = "published"
	ArchivedMeal  MealType = "archived"
)

type MealPlan struct {
	gorm.Model
	CafeteriaId   string
	WeekStartDate time.Time
	WeekEndDate   time.Time
	CreatedById   string
	Status        MealType `gorm:"default:'draft'"`
}
