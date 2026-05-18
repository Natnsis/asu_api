package complaints

import "gorm.io/gorm"

type Complaints struct {
	gorm.Model
	ComplaintId   string `gorm:"uniqueIndex;not null" json:"student_id"`
	Title         string `json:"title"`
	Description   string `gorm:"not null" json:"description"`
	AttachmentUrl string `json:"attachment_url"`
}
