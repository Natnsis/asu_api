package models

import "gorm.io/gorm"

type GalleryCategories struct {
	gorm.Model
	Name        string
	Description string
	Gallery     []Gallery
}
