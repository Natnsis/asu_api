package models

import (
	"time"

	"gorm.io/gorm"
)

type Complaint struct {
	gorm.Model
	universityId  string
	studentId     string
	categoryId    string
	departmentId  string
	title         string
	descritpion   string
	attachmentUrl string
	priority      string
	status        string
	assignedTold  string
	resolvedAt    time.Time
}
