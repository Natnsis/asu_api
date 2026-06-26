package models

import (
	"time"

	"gorm.io/gorm"
)

type Gallery struct {
	gorm.Model
	universityId string
	createdById  string
	title        string
	description  string
	coverMediaId string
	status       string
	publishedAt  time.Time
}
