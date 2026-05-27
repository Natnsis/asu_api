package models

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	Title               string
	CoverImage          string
	AlbumIds            string
	EventType           string
	CategoryIds         []string
	RelatedEventId      string
	RelatedDepartmentId string
	TargetAudience      string
	AllowLikes          bool
	Likes               int
	AlbumId             Album
}
