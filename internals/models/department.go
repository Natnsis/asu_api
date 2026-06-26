package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	universityId string
	name         string
	code         string
	headId       string
	isActive     bool
}
