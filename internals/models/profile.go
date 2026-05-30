package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	AvatarUrl string
	Address   string
	Major     string
	UserID    uint
	User      User
}
