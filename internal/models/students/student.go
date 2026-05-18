package students

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentId       string `gorm:"uniqueIndex;not null" json:"student_id"`
	StudentName     string `gorm:"not null" json:"student_name"`
	IsActive        bool   `gorm:"not null; default:" json:"is_active"`
	TotalInvitation int    `gorm:"default:0" json:"total_invitation"`
	SpecialRight    string `gorm:"not null" json:"special_right"`
}
