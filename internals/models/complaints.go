package models

import (
	"time"

	"gorm.io/gorm"
)

// custom types
type (
	Priorities string
	Statuses   string
)

// custom enums
const (
	LowPriority      Priorities = "low"
	MediumPriority   Priorities = "medium"
	HighPriority     Priorities = "high"
	CriticalPriority Priorities = "critical"
)

const (
	SubmittedStatus   Statuses = "submitted"
	UnderReviewStatus Statuses = "under_review"
	ResolvedStatus    Statuses = "resolved"
	RejectedStatus    Statuses = "rejected"
)

type Complaint struct {
	gorm.Model
	UniversityID        uint
	UserID              uint
	CategoryID          uint
	DepartmentID        uint
	Title               string
	Description         string
	AttachmentUrl       string
	Priority            Priorities `gorm:"default:'low'"`
	Status              Statuses   `gorm:"default:'submitted'"`
	AssignedToId        uint
	ResolvedComplaintAt time.Time

	ComplaintResponse []ComplaintResponse
	ComplaintCategory ComplaintCategory `gorm:"foreignKey:CategoryID"`
}
