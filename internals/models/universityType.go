package models

import "gorm.io/gorm"

type UniversityType struct {
	gorm.Model
	Name         string
	Universities []University
}
