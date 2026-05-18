package events

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title          string    `json:"event_title"`
	Description    string    `json:"event_description"`
	Type           string    `json:"event_type"`
	StartDateTime  time.Time `json:"event_start_date"`
	EndDateTime    time.Time `json:"event_end_date"`
	Location       string    `json:"event_location"`
	TargetAudience string    `json:"event_audience"`
	Organizer      string    `json:"event_organizer"`
}
