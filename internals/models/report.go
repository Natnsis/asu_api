package models

import (
	"time"

	"gorm.io/gorm"
)

// cusom types
type ReportStatus string

// custom enums
const (
	ReportDraft     ReportStatus = "draft"
	ReportPublished ReportStatus = "published"
	ReportArchived  ReportStatus = "archived"
)

type Report struct {
	gorm.Model
	UniversityId string
	CategoryId   string
	CreatedById  string
	DepartmentId string
	Title        string
	Description  string
	FileUrl      string
	FileType     string
	Status       ReportStatus `gorm:"ReportStatus='draft'"`
	PublishedAt  time.Time
}
