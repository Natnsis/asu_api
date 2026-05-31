package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name         string
	ParentRoleID *uint
	User         []User
	SubRoles     []Role `gorm:"foreignKey:ParentRoleID"`
}
