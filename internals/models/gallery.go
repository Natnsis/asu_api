package models

import (
	"time"

	"gorm.io/gorm"
)

// custom type
type GalleryStatus string

// custom enum
const (
	GelleryDraft     GalleryStatus = "draft"
	GelleryPublished GalleryStatus = "published"
	GelleryArchived  GalleryStatus = "archived"
)

type Gallery struct {
	gorm.Model
	UniversityID uint
	CreatedById  string
	Title        string
	Description  string
	CoverMediaId string
	Status       GalleryStatus `gorm:"default:'draft'"`
	PublishedAt  time.Time

	GalleryMedia []GalleryMedia
	University   University
}
