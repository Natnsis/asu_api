package models

import (
	"time"

	"gorm.io/gorm"
)

type Cafeteria struct {
	gorm.Model
	UniveristyId string
	Name         string
	Location     string
	OpenTime     time.Time
	CloseTime    time.Time
	IsActive     bool
}
