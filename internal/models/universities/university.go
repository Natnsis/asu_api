package universities

import "gorm.io/gorm"

type Universities struct {
	gorm.Model

	OfficialName      string `json:"official_name"`
	Code              string `json:"university_code"`
	EstablishmentYear string `json:"establishment_year"`
	InstitutionType   string `json:"institution_type"`
	OfficialSite      string `json:"official_site"`
}
