package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID       uint
	UniversityID uint
	TemplateID   uint
	Title        string
	Body         string
	Channel      DeliveryChannels
	IsRead       bool
	ReadAt       time.Time
	EntryType    string
	EntryID      string

	University University
	User       User
}
