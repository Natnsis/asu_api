package models

import "gorm.io/gorm"

type ComplaintCategory struct {
	gorm.Model
	UniversityID uint
	Name         string

	University University
}
