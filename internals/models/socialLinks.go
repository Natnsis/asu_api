package models

import "gorm.io/gorm"

type SocialLink struct {
	gorm.Model
	Name string
	Url  string
}
