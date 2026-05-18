package recent

import "gorm.io/gorm"

type RecentActivities struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	UserType    string `json:"user_type"`
	Description string `json:"recent_description"`
	Action      string `json:"operation_action"`
}
