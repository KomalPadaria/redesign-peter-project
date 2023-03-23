package knowbe4

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4/service"
)

type Client interface {
	GetAllPhishingSecurityTests(ctx context.Context, companyUuid uuid.UUID) ([]entities.SecurityTest, error)
	GetAllRecipientResults(ctx context.Context, companyUuid uuid.UUID, pstId int) ([]entities.RecipientResult, error)
	GetAllTrainingEnrollments(ctx context.Context, companyUuid uuid.UUID) ([]entities.TrainingEnrollment, error)
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) GetAllTrainingEnrollments(ctx context.Context, companyUuid uuid.UUID) ([]entities.TrainingEnrollment, error) {
	return l.svc.GetAllTrainingEnrollments(ctx, companyUuid)

}

func (l *localClient) GetAllRecipientResults(ctx context.Context, companyUuid uuid.UUID, pstId int) ([]entities.RecipientResult, error) {
	return l.svc.GetAllRecipientResults(ctx, companyUuid, pstId)
}

func (l *localClient) GetAllPhishingSecurityTests(ctx context.Context, companyUuid uuid.UUID) ([]entities.SecurityTest, error) {
	return l.svc.GetAllPhishingSecurityTests(ctx, companyUuid)
}
