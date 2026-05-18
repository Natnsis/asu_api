package curriculums

import "gorm.io/gorm"

type Curriculum struct {
	gorm.Model
	ProgramName              string  `gorm:"not null" json:"program_name"`
	ProgramCode              string  `gorm:"not null" json:"program_code"`
	DegreeAwarded            int     `gorm:"not null" json:"semester"`
	DurationInSemesters      int     `gorm:"not null" json:"year"`
	MinimumCreditsToGraduate float32 `gorm:"not null" json:"minimum_credits_to_graduate"`
	PassingGpa               float32 `gorm:"not null" json:"passing_gpa"`
}
