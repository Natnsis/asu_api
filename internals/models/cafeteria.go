package models

import (
	"gorm.io/gorm"
)

type Cafeteria struct {
	gorm.Model
	DayOfWeek  string
	MealType   string
	FoodName   string
	FoodImgUrl string
	IsSpecial  bool
	notes      string
}
