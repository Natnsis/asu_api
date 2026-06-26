package models

import "gorm.io/gorm"

type CourseMaterial struct {
	gorm.Model
	courseId     string
	uploadedById string
	title        string
	types        string
	url          string
	weekNumber   int
	semesterId   string
}
