package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserId    uint   `gorm:"unique;not null"`
	User      *User  `gorm:foreignKey:UsreId`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}
