package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	name         string `gorm:"not null"`
	slug         string `gorm:"uniqueIndex"`
	logoUrl      string
	address      string
	contactEmail string `gorm:"not null"`
	contactPhone string
	isActive     bool `gorm:"not null"`

	Records              []Records
	NotificationTemplete []NotificationTemplete
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
