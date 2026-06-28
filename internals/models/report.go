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
	UniversityID uint
	CategoryID   uint
	UserID       uint
	DepartmentID uint
	Title        string
	Description  string
	FileUrl      string
	FileType     string
	Status       ReportStatus `gorm:"default:'draft'"`
	PublishedAt  time.Time

	University     University
	ReportCategory ReportCategory `gorm:"foreignKey:CategoryID"`
	Department     Department
}
