package entities

import (
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type GetConsumedHoursRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetConsumedHoursResponse struct {
	TicketNumber string            `json:"ticket_number"`
	Type         string            `json:"type"`
	AssignedTo   string            `json:"assigned_to"`
	CloseDate    nullable.NullTime `json:"close_date"`
	TimeConsumed string            `json:"time_consumed"`
}
