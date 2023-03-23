package entities

import (
	"time"

	"github.com/google/uuid"
)

type GetPhishingDetailsResponse struct {
	CurrentNextCampaign CurrentNextCampaign `json:"campaign"`
	PhishingStat        PhishingStat        `json:"stat"`
	TopClickers         []Clicker           `json:"top_clickers"`
}

type Clicker struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	ClickedCount int    `json:"clicked_count"`
}

type GetPhishingDetailsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type CurrentNextCampaign struct {
	Current Campaign `json:"current"`
	Next    Campaign `json:"next"`
}

type Campaign struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type PhishingStat struct {
	Passed           int `json:"passed"`
	Failed           int `json:"failed"`
	TotalParticipant int `json:"total_participant"`
}

type GetTrainingDetailsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetTrainingDetailsResponse struct {
	CurrentNextCampaign CurrentNextCampaign `json:"campaign"`
	TrainingStat        TrainingStat        `json:"stat"`
}

type TrainingStat struct {
	NotStarted       int `json:"not_started"`
	InProgress       int `json:"in_progress"`
	Completed        int `json:"completed"`
	Passed           int `json:"passed"`
	PastDue          int `json:"past_due"`
	TotalEnrollments int `json:"total_enrollments"`
}
