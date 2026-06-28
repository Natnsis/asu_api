package models

import "gorm.io/gorm"

type ReportCategory struct {
	gorm.Model
	UniversityId string
	Name         string
	AccessRoles  []string
}
