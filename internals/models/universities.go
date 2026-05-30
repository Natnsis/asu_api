package models

import "gorm.io/gorm"

type Universities struct {
	gorm.Model
	Name                    string
	Description             string
	EstablishedYear         string
	UniversityType          string
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
	SocialLinkID            string
	UniversityID            uint
	College                 []College
	SocialLink              []SocialLink
}
