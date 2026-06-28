package models

import "gorm.io/gorm"

type ComplaintResponse struct {
	gorm.Model
	ComplaintID   uint
	RespondedById uint
	Message       string
	IsInternal    bool

	Complaint Complaint
}
