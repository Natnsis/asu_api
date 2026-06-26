package models

import "gorm.io/gorm"

type ReportCategory struct {
	gorm.Model
	universityId string
	name         string
	accessRoles  string
}
