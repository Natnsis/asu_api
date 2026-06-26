package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	userId       string
	univeristyId string
	templeteId   string
	title        string
	body         string
	channel      string
	isRead       bool
	readAt       time.Time
	entryType    string
	entryId      string
}
