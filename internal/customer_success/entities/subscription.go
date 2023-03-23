package entities

import "github.com/google/uuid"

type GetSubscriptionsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
}

type GetSubscriptionsResponse struct {
	SubscriptionName string    `json:"subscription_name"`
	StartDate        string    `json:"start_date"`
	EndDate          string    `json:"end_date"`
	Services         []Service `json:"services"`
}

type Service struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Type      string `json:"type"`
}
