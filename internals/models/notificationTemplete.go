package models

import "gorm.io/gorm"

type NotificationTemplete struct {
	gorm.Model
	universityId string
	key          string
	title        string
	body         string
	channel      string
}
