package models

import "gorm.io/gorm"

// custom type
type (
	DaysType  string
	MealTypes string
)

// custom enum
const (
	Monday    DaysType = "monday"
	Tuesday   DaysType = "tuesday"
	Wednesday DaysType = "wednesday"
	Thursday  DaysType = "thursday"
	Friday    DaysType = "friday"
	Saturday  DaysType = "saturday"
	Sunday    DaysType = "sunday"
)

const (
	Breakfast MealType = "breakfast"
	Lunch     MealType = "lunch"
	Dinner    MealType = "dinner"
)

type MealSlot struct {
	gorm.Model
	MealPlanID  uint
	DayOfWeek   DaysType
	MealPeriod  string
	MealName    MealType
	Descritpion string

	MealPlan MealPlan
}
