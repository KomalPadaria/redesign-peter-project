package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/service"
)

type Client interface {
	CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	return l.svc.CreateCompanyQuestionnaires(ctx, companyUUID, userUUID)
}
