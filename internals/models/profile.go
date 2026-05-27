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

	// belongs to (optional) one to one
	User *User `gorm:"foreignKey:UserID"`
}
