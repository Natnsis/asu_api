package models

import "gorm.io/gorm"

type ComplaintResponse struct {
	gorm.Model
	ComplaintId   string
	RespondedById string
	Message       string
	IsInternal    bool
}
