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
	Breakfast MealTypes = "breakfast"
	Lunch     MealTypes = "lunch"
	Dinner    MealTypes = "dinner"
)

type MealSlot struct {
	gorm.Model
	MealPlanID  uint
	DayOfWeek   DaysType
	MealPeriod  MealTypes
	MealName    string
	Description string

	MealPlan MealPlan
}
