package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserId       string
	UniveristyId string
	TempleteId   string
	Title        string
	Body         string
	Channel      DeliveryChannels
	IsRead       bool
	ReadAt       time.Time
	EntryType    string
	EntryId      string
}
