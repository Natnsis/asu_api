package models

import "gorm.io/gorm"

type ComplaintCategory struct {
	gorm.Model
	UniveristyID uint
	Name         string

	University University
}
