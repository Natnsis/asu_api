package models

import "gorm.io/gorm"

type ComplaintCategory struct {
	gorm.Model
	univeristyId string
	name         string
}
