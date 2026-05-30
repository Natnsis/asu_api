package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	PasswordHash    string
	RoleID          []uint
	Status          string
	TotalInvitation int
	UniversityId    string
	Role            Role
}
