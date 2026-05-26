package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	PasswordHash    string
	RoleId          []string
	Status          string
	TotalInvitation int
}
