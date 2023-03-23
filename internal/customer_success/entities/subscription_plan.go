package entities

import (
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type GetSubscriptionPlansRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
}

type GetSubscriptionPlansResponse struct {
	PlanId    uuid.UUID         `json:"subscription_plan_uuid "`
	Status    string            `json:"status"`
	PlanType  string            `json:"plan_type"`
	StartDate nullable.NullTime `json:"start_date"`
	EndDate   nullable.NullTime `json:"end_date"`
}
