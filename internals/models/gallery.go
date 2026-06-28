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
	universityId string
	createdById  string
	title        string
	description  string
	coverMediaId string
	status       GalleryStatus `gorm:"default:'draft'`
	publishedAt  time.Time
}
