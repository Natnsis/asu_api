package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	univerisityId string
	key           string
	value         string
	description   string
	updatedById   string
}
