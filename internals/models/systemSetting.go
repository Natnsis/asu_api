package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	UniverisityID uint
	Key           string
	Value         string
	Description   string
	UpdatedById   string
}
