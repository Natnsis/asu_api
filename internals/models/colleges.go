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
	CollegeId       uint
	Department      []Department

	Universities University
}
