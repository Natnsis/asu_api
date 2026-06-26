package models

import "gorm.io/gorm"

type ComplaintResponse struct {
	gorm.Model
	complaintId   string
	respondedById string
	message       string
	isInternal    bool
}
