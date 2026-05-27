package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Name        string
	Description string
	Gallery     []Gallery `gorm:"foreignKey:AlbumID;constraint:OnDelete:CASCADE"`
}
