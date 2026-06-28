package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	UniversityID uint
	Key           string
	Value         string
	Description   string
	UpdatedById   string
}
