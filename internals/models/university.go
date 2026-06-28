package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Slug         string `gorm:"uniqueIndex"`
	LogoUrl      string
	Address      string
	ContactEmail string `gorm:"not null"`
	ContactPhone string
	IsActive     bool `gorm:"not null"`

	Records              []Records
	NotificationTemplate []NotificationTemplate
	Notification         []Notification
	Lounge               []Lounge
	EventCategory        []EventCategory
	Gallery              []Gallery
	ComplaintCategory    []ComplaintCategory
	Department           []Department
	User                 []User
	Audit                []Audit
	Setting              Setting
}
