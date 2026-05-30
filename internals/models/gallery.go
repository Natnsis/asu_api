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
	AlbumId             Album
	UserId              uint
	User                User
	AlbumID             uint
	Album               Album
	GalleryCategoriesID uint
	GalleryCategories   GalleryCategories
}
