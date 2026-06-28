package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ReportCategory struct {
	gorm.Model
	UniversityId uint
	Name         string
	AccessRoles  datatypes.JSON
}
