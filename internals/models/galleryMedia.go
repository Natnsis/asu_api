package models

import (
	"gorm.io/gorm"
)

// custom type
type FileType string

// custom enum
const (
	FileVideo FileType = "video"
	FilePhoto FileType = "photo"
)

type GalleryMedia struct {
	gorm.Model
	galleryId    string
	uplaodedById string
	types        FileType `gorm:"default:'photo'"`
	url          string
	thumbnailUrl string
	caption      string
	sortOrder    int
}
