package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	universityId string
	Name         string
	Code         string
	HeadId       string
	IsActive     bool
}
