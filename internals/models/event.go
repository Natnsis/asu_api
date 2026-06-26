package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	universityId         string
	categoryId           string
	createdById          string
	title                string
	descritpion          string
	location             string
	startDate            time.Time
	endDate              time.Time
	coverImageUrl        string
	requiresRegistration bool
	maxAttendees         int
	status               string
}
