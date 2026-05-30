package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type University struct {
	gorm.Model
	Name                    string
	Description             string
	EstablishedYear         string
	AccreditionStatus       string
	LogoUrl                 string
	CoverImageUrl           string
	Motto                   string
	OfficialEmail           string
	SupportEmail            string
	PhoneNumber             int
	WebUrl                  string
	Country                 string
	Region                  string
	City                    string
	PostalCode              string
	Latitude                string
	Longitude               string
	MapUrl                  string
	ApproximateTotalStudent int
	ApproximateTotalStaff   int
	TotalDepartment         int
	TotalPrograms           int
	MainAdminId             string
	MaxAdminAllowed         int
	MaxStudentsAllowed      int
	SocialLinkID            datatypes.JSONSlice[string]
	UniversityTypeId        uint
	College                 []College
	SocialLink              []SocialLink
	UniversityType          UniversityType
}
