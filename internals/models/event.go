package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title              string
	Subtitle           string
	Description        string
	EventTypeId        []string
	CategoryId         []string
	ThumbnailUrl       string
	TargetAudience     []string
	ApprovedBy         []string
	MaxParticipants    int
	AttendanceRequired bool
}
