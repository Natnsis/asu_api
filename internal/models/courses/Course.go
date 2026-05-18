package courses

import "gorm.io/gorm"

type Courses struct {
	gorm.Model

	Code          string `gorm:"not null;uniqueIndex" json:"course_code"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	CreditHours   string `json:"credit_hours"`
	LectureHours  string `json:"lecture_hours"`
	LabHours      int    `json:"lab_hours"`
	Status        string `json:"course_status"`
	CurriculumtId string `json:"curriculum_id"`
}
