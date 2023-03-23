package onboarding

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/service"
)

type Client interface {
	// Update the onboarding status from "DRAFT" to "COMPLETE" for a given step
	UpdateOnboardingStatus(ctx context.Context, onboardingStep int, companyUuid, updated_by uuid.UUID)
}

func newClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) UpdateOnboardingStatus(ctx context.Context, onboardingStep int, companyUuid, updated_by uuid.UUID) {
	l.svc.UpdateOnboardingStatus(ctx, onboardingStep, companyUuid, updated_by)
}
