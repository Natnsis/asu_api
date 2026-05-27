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
	UserId    uint `gorm:"unique;not null"`
}
