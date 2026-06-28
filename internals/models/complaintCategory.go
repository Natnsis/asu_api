package models

import "gorm.io/gorm"

type ComplaintCategory struct {
	gorm.Model
	UniveristyId string
	Name         string
}
