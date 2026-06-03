package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string
	PasswordHash    string
	RoleID          *uint
	Status          string
	TotalInvitation int
	UniversityID    *uint
	DepartmentID    *uint
	StudentTypeID   *uint

	Profile        Profile          `gorm:"foreignKey:UserID"`
	RecentActivity []RecentActivity `gorm:"foreignKey:UserID"`
	Event          []Event          `gorm:"foreignKey:UserID"`
	Gallery        []Gallery        `gorm:"foreignKey:UserID"`
	Department     Department       `gorm:"foreignKey:DepartmentID"`
	StudentType    StudentType      `gorm:"foreignKey:StudentTypeID"`
	University     University       `gorm:"foreignKey:UniversityID"`
	Role           Role             `gorm:"foreignKey:RoleID"`
}
