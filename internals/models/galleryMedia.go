package models

import (
	"gorm.io/gorm"
)

type GalleryMedia struct {
	gorm.Model
	galleryId    string
	uplaodedById string
	types        string
	url          string
	thumbnailUrl string
	caption      string
	sortOrder    int
}
