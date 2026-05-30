package models

import "gorm.io/gorm"

type RecentActivity struct {
	gorm.Model
	Username    string
	UserAvatar  string
	UserRole    string
	Action      string
	Title       string
	Description string
	UserID      uint
	User        User
}
