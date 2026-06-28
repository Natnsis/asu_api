package models

import "gorm.io/gorm"

type DeliveryChannels string

const (
	PushChannel  DeliveryChannels = "push"
	EmailChannel DeliveryChannels = "email"
	InAppChannel DeliveryChannels = "in_app"
)

type NotificationTemplate struct {
	gorm.Model
	UniversityID uint
	Key          string
	Title        string
	Body         string
	Channel      DeliveryChannels `gorm:"default:'push'"`

	University University
}
