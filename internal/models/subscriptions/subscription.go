package subscriptions

import "gorm.io/gorm"

type Subscriptions struct {
	gorm.Model

	UniversityId string `gorm:"uniqueIndex; not null" json:"university_id"`
	Status       string `gorm:"not null" json:"status"`
	PlanType     string `json:"plan_type"`
	BillingCycle string `json:"billing_cycle"`
}
