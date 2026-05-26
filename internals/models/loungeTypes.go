package models

import "gorm.io/gorm"

type LoungeTypes struct {
	gorm.Model
	Name        string
	Description string
}
