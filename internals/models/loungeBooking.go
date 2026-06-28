package models

import (
	"time"

	"gorm.io/gorm"
)

// custom lounge status types
type LoungeStatus string

// custom enum
const (
	PendingLounge   LoungeStatus = "pending"
	ConfirmedLounge LoungeStatus = "confirmed"
	CancelledLounge LoungeStatus = "cancelled"
	NoShowLounge    LoungeStatus = "no_show"
)

type LoungeBooking struct {
	gorm.Model
	LoungeID   uint
	BookedByID uint
	Date       time.Time
	StartTime  time.Time
	EndTime    time.Time
	Attendees  int
	Purpose    string
	Status     LoungeStatus `gorm:"default:'pending'"`

	Lounge Lounge
}
