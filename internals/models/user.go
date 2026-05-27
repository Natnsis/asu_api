package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	PasswordHash    string
	RoleIds         []uint
	Status          string
	TotalInvitation int
	UniversityId    string
	Role            Role    `gorm:"foreignKey:RoleID"`
	Profile         Profile `gorm:"constraint:OnDelete:CASCADE"`
}
