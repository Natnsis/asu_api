package models

import (
	"time"

	"gorm.io/gorm"
)

// custom type
type EventStatus string

// custom enum
const (
	EventDraft     EventStatus = "draft"
	EventPublished EventStatus = "published"
	EventCancelled EventStatus = "cancelled"
	EventCompleted EventStatus = "completed"
)

type Event struct {
	gorm.Model
	UniversityId         string
	CategoryID           string
	CreatedByID          string
	Title                string
	Descritpion          string
	Location             string
	StartDate            time.Time
	EndDate              time.Time
	CoverImageUrl        string
	RequiresRegistration bool
	MaxAttendees         int
	Status               EventStatus `gorm:"default:'draft'"`

	Records           Records
	EventRegistration []EventRegistration
	EventCategory     []EventCategory
}
