package models

import "gorm.io/gorm"

// define custom level
type DegreeLevel string

// define the enum
const (
	LevelCertificate DegreeLevel = "certificate"
	LevelDiploma     DegreeLevel = "diploma"
	LevelBachelors   DegreeLevel = "bachelors"
	LevelMaster      DegreeLevel = "masters"
)

type Program struct {
	gorm.Model
	UniversityID  uint
	DepartmentID  uint
	Name          string
	Code          string
	DurationYears int
	DegreeLevel   DegreeLevel `gorm:"default:'bachelors'"`

	Profile    Profile
	Department Department
}
