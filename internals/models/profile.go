package models

import (
	"time"

	"gorm.io/gorm"
)

// custom status type and gender type
type (
	StatusTypes string
	GenderTypes string
)

// custom enums
const (
	StatusActive    StatusTypes = "active"
	StatusSuspended StatusTypes = "suspended"
	StatusGraduated StatusTypes = "graduated"
	StatusWithdrawn StatusTypes = "withdrawn"
)

const (
	GenderMale   GenderTypes = "male"
	GenderFemale GenderTypes = "female"
)

type Profile struct {
	gorm.Model
	UserID         uint
	StudentId      string
	ProgramID      uint
	DepartmentId   string
	EnrollmentYear int
	CurrentYear    int
	Status         StatusTypes `gorm:"default:'active'"`
	DateOfBirth    time.Time
	Gender         GenderTypes `gorm:"default:'male'"`

	User User
}
