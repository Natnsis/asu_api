package schedule

import "gorm.io/gorm"

type (
	DayOfWeek   string
	SessionType string
)

const (
	DayMonday    DayOfWeek = "MONDAY"
	DayTuesday   DayOfWeek = "TUESDAY"
	DayWednesday DayOfWeek = "WEDNESDAY"
	DayThursday  DayOfWeek = "THURSDAY"
	DayFriday    DayOfWeek = "FRIDAY"
	DaySaturday  DayOfWeek = "SATURDAY"
	DaySunday    DayOfWeek = "SUNDAY"

	SessionBreakfast SessionType = "BREAKFAST"
	SessionLunch     SessionType = "LUNCH"
	SessionDinner    SessionType = "DINNER"
)

type Schedule struct {
	gorm.Model

	CafeteriaID uint        `gorm:"uniqueIndex:idx_cafeteria_schedule;not null" json:"cafeteria_id"`
	Day         DayOfWeek   `gorm:"uniqueIndex:idx_cafeteria_schedule;type:varchar(15);not null" json:"day_of_week"`
	Session     SessionType `gorm:"uniqueIndex:idx_cafeteria_schedule;type:varchar(20);not null" json:"session_type"`

	MainDish    string  `gorm:"type:varchar(150);not null" json:"main_dish"`
	SideDish    string  `gorm:"type:varchar(150)" json:"side_dish"`
	Price       float64 `gorm:"type:decimal(10,2);default:0.00" json:"price"`
	DietaryTags string  `gorm:"type:varchar(100)" json:"dietary_tags"`
}
