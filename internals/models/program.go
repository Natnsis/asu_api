package models

import "gorm.io/gorm"

// define custom level
type DegreeLevel string

// define the enum
const (
	LevelCertificate DegreeLevel = "certificate"
	LevelDeploma     DegreeLevel = "deploma"
	LevelBachlors    DegreeLevel = "bachlors"
	LevelMaster      DegreeLevel = "masters"
)

type Program struct {
	gorm.Model
	UniversityId  string
	DepartmentId  string
	Name          string
	Code          string
	DurationYears int
	DegreeLevel   DegreeLevel `gorm:"default:'bachlors'"`
}
