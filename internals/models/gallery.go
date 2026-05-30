package models

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	Title               string
	CoverImage          string
	RelatedEventId      string
	RelatedDepartmentId string
	TargetAudience      string
	AllowLikes          bool
	Likes               int
	UserID              uint
	AlbumID             uint
	GalleryCategoriesID uint

	Album             Album
	User              User
	GalleryCategories GalleryCategories
}
