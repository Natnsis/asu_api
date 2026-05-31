package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	PasswordHash    string
	RoleID          uint
	Status          string
	TotalInvitation int
	UniversityID    uint
	DepartmentID    uint
	StudentTypeID   uint

	RecentActivity []RecentActivity
	Event          []Event
	Gallery        []Gallery
	Department     Department
	StudentType    StudentType
	University     University
	Role           Role
}
