package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	IsActive bool   `gorm:"dafault:true" json:"is_active"`
	RoleId   uint   `json:"role_id"`
	Role     Role   `gorm:"foreignKey: RoleId" json:"role"`
}
