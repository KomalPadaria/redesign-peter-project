package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/vulnerability"
	vulnerabilityEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/vulnerability/entities"
)

type Service interface {
	GetPenetrationTopRemediationTasks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*vulnerabilityEntities.TopRemediation, error)
	GetPenetrationStats(ctx context.Context, companyUuid, userUuid *uuid.UUID) (*vulnerabilityEntities.VulnerabilityStats, error)
}

// New service for user.
func New(vulnerabilityClient vulnerability.Client) Service {
	svc := &service{vulnerabilityClient}

	return svc
}

type service struct {
	vulnerabilityClient vulnerability.Client
}

func (s *service) GetPenetrationTopRemediationTasks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*vulnerabilityEntities.TopRemediation, error) {
	return s.vulnerabilityClient.GetPenetrationTopRemediationTasks(ctx, companyUuid, userUuid)
}

func (s *service) GetPenetrationStats(ctx context.Context, companyUuid, userUuid *uuid.UUID) (*vulnerabilityEntities.VulnerabilityStats, error) {
	return s.vulnerabilityClient.GetPenetrationStats(ctx, companyUuid, userUuid)
}
