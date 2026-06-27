package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	StudentId      string
	ProgramId      string
	DepartmentId   string
	EnrollmentYear int
	CurrentYear    int
	Status         string
	DateOfBirth    time.Time
	Gender         string
}
