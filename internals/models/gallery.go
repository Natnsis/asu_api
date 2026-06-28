package models

import (
	"time"

	"gorm.io/gorm"
)

// custom type
type GalleryStatus string

// custom enum
const (
	GalleryDraft     GalleryStatus = "draft"
	GalleryPublished GalleryStatus = "published"
	GalleryArchived  GalleryStatus = "archived"
)

type Gallery struct {
	gorm.Model
	UniversityID uint
	CreatedById  uint
	Title        string
	Description  string
	CoverMediaId uint
	Status       GalleryStatus `gorm:"default:'draft'"`
	PublishedAt  time.Time

	GalleryMedia []GalleryMedia
	University   University
}
