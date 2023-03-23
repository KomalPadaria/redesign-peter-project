package jira

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/service"
)

type Client interface {
	CreateEpic(ctx context.Context, req *entities.CreateIssueRequest) (*entities.CreateIssueResponse, error)
	AddComment(ctx context.Context, companyUuid *uuid.UUID, comment *entities.Comment) error
	GetIssuesByEpicId(ctx context.Context, epicId string) ([]entities.Issues, error)
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) CreateEpic(ctx context.Context, req *entities.CreateIssueRequest) (*entities.CreateIssueResponse, error) {
	return l.svc.CreateEpic(ctx, req)
}

func (l *localClient) GetIssuesByEpicId(ctx context.Context, epicId string) ([]entities.Issues, error) {
	return l.svc.GetIssuesByEpicId(ctx, epicId)
}

func (l *localClient) AddComment(ctx context.Context, companyUuid *uuid.UUID, comment *entities.Comment) error {
	return l.svc.AddComment(ctx, companyUuid, comment)
}
