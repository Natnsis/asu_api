package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string  `json:"email" gorm:"unique;not null"`
	Username    string  `json:"username" gorm:"unique;not null"`
	Password    string  `json:"-"`
	PhoneNumber string  `json:"phone_number"`
	Profile     Profile `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
}
