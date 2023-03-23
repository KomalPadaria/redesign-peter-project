package entities

import "time"

type TrainingEnrollment struct {
	EnrollmentID       int       `json:"enrollment_id"`
	ContentType        string    `json:"content_type"`
	ModuleName         string    `json:"module_name"`
	User               User      `json:"user"`
	CampaignName       string    `json:"campaign_name"`
	EnrollmentDate     time.Time `json:"enrollment_date"`
	StartDate          time.Time `json:"start_date"`
	CompletionDate     time.Time `json:"completion_date"`
	Status             string    `json:"status"`
	TimeSpent          int       `json:"time_spent"`
	PolicyAcknowledged bool      `json:"policy_acknowledged"`
}
