package wireless

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/service"
)

type Client interface {
	UpdateTechInfoWirelessStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) UpdateTechInfoWirelessStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error {
	return l.svc.UpdateTechInfoWirelessStatusByFacilities(ctx, userUUID, facilityUuids, status)
}
