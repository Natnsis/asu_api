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
	CafeteriaID   uint
	WeekStartDate time.Time
	WeekEndDate   time.Time
	UserID        uint
	Status        MealType `gorm:"default:'draft'"`

	Cafeteria Cafeteria
	MealSlot  []MealSlot
}
