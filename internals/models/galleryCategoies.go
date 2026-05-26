package models

import "gorm.io/gorm"

type GalleryCategories struct {
	gorm.Model
	Name         string
	UniversityId string
	Description  string
}
