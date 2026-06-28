package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID       uint
	UniveristyID uint
	TempleteID   string
	Title        string
	Body         string
	Channel      DeliveryChannels
	IsRead       bool
	ReadAt       time.Time
	EntryType    string
	EntryID      string

	User User
}
