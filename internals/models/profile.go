package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	studentId      string
	programId      string
	departmentId   string
	enrollmentYear int
	currentYear    int
	status         string
	dateOfBirth    time.Time
	gender         string
}
