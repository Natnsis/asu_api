package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Name        string
	Description string
	Gallery     []Gallery
}
