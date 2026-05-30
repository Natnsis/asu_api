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
	DepartmentID    uint
	RecentActivity  []RecentActivity
	Event           []Event
	Gallery         []Gallery
	Department      Department
	StudentTypeID   uint
	StudentType     StudentType
	RoleId          uint
	Role            Role
}
