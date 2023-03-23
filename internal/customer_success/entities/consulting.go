package entities

import (
	"github.com/google/uuid"
)

type GetConsultingHoursRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetConsultingHoursResponse struct {
	SubscriptionID string  `json:"subscription_id"`
	Name           string  `json:"name"`
	HoursConsumed  float64 `json:"hours_consumed"`
	HoursTotal     float64 `json:"hours_total"`
	From           string  `json:"from"`
	To             string  `json:"to"`
}
