package models

import "gorm.io/gorm"

type MealSlot struct {
	gorm.Model
	mealPlanId  string
	dayOfWeek   string
	mealPeriod  string
	mealName    string
	descritpion string
}
