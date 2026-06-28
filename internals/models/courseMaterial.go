package models

import "gorm.io/gorm"

// custom types
type Types string

// custom enum
const (
	TypePdf      Types = "pdf"
	TypeVideo    Types = "video"
	TypeLink     Types = "link"
	TypeSlide    Types = "slide"
	TypeDocument Types = "document"
)

type CourseMaterial struct {
	gorm.Model
	CourseID     uint
	UploadedById uint
	Title        string
	Types        Types `gorm:"default:'pdf'"`
	Url          string
	WeekNumber   int
	SemesterID   uint

	Course Course
}
