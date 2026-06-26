package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	firstName    string
	lastName     string
	email        string
	passwordHash string
	phone        string
	avatarUrl    string
	universityId string
	departmentId string
	isActive     bool
	lastLoginAt  *time.Time
}
