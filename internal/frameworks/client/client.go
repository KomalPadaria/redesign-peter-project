package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/service"
)

type Client interface {
	CreateFrameworkControlRemediations(ctx context.Context, companyUUID, userUUID *uuid.UUID) error
	CreateCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error
	DeleteCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) DeleteCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error {
	return l.svc.DeleteCompanyFrameworksLink(ctx, companyUUID, userUUID, frameworkNames)
}

func (l *localClient) CreateCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error {
	return l.svc.CreateCompanyFrameworksLink(ctx, companyUUID, userUUID, frameworkNames)
}

func (l *localClient) CreateFrameworkControlRemediations(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	return l.svc.CreateFrameworkControlRemediations(ctx, companyUUID, userUUID)
}
