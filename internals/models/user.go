package models

import (
	"time"

	"gorm.io/gorm"
)

// define custom role
type UserRole string

// define enum constants
const (
	RoleSuperAdmin     UserRole = "super_admin"
	RoleAdmin          UserRole = "admin"
	RoleDepartmentHead UserRole = "department_head"
	RoleLecturer       UserRole = "lecturer"
	RoleStudent        UserRole = "student"
	RoleClassRep       UserRole = "class_rep"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	passwordHash string
	Phone        string
	AvatarUrl    string
	UniversityID uint
	DepartmentID uint
	Role         UserRole `gorm:"default:'student'"`
	IsActive     bool
	LastLoginAt  *time.Time

	Notification []Notification
	Audit        []Audit
}
