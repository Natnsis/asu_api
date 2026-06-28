package models

import (
	"time"

	"gorm.io/gorm"
)

// custom type
type EventStatusType string

// enum custom
const (
	EventConfirmedType  EventStatusType = "confirmed"
	EventWaitlistedType EventStatusType = "waitlisted"
	EventCancelledType  EventStatusType = "cancelled"
)

type EventRegistration struct {
	gorm.Model
	EventID      uint
	UserID       uint
	RegisteredAt time.Time
	Status       EventStatusType `gorm:"default:'confirmed'"`

	User  User
	Event Event
}
