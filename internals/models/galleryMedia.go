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
	GalleryID    uint
	UserID       uint
	Types        FileType `gorm:"default:'photo'"`
	Url          string
	ThumbnailUrl string
	Caption      string
	SortOrder    int

	Gallery Gallery
	User    User
}
