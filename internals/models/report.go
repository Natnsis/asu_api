package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	universityId string
	categoryId   string
	createdById  string
	departmentId string
	title        string
	description  string
	fileUrl      string
	fileType     string
	status       string
	publishedAt  time.Time
}
