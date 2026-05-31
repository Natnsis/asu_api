package models

import (
	"time"

	"gorm.io/gorm"
)

type College struct {
	gorm.Model
	Name            string
	Code            string
	Description     string
	EstablishedYear time.Time
	DeanId          string
	ViceDeanId      string
	Status          string
	UniversityID    uint

	Department []Department
	University University
}
