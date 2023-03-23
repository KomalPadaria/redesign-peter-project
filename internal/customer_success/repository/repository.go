package repository

import (
	"context"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"gorm.io/gorm"
)

// Repository for user.
type Repository interface {
	CreateEvidences(ctx context.Context, serviceReports companyEntities.ServiceEvidence) error
	DeleteEvidenceFile(ctx context.Context, companySubscriptionsUuid *uuid.UUID, reportName string) error
	GetEvidenceByEvidenceId(ctx context.Context, evidenceUuid *uuid.UUID) (*companyEntities.ServiceEvidence, error)
	DeleteEvidence(ctx context.Context, evidenceUuid *uuid.UUID) error
	UpdateEvidence(ctx context.Context, evidenceUuid *uuid.UUID, updateData map[string]interface{}) error
}

// New repository for company.
func New(gormdb *gorm.DB) Repository {
	repo := &sqlRepository{gormdb}

	return repo
}
