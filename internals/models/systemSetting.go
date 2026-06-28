package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	UniverisityId string
	Key           string
	Value         string
	Description   string
	UpdatedById   string
}
