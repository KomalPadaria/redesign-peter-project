package repository

import (
	"context"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetOnboardingStatus(ctx context.Context, companyUuid uuid.UUID) (*companyEntities.OnboardingGroup, error)
	UpdateOnboardingStatus(ctx context.Context, onboardingGroup *companyEntities.OnboardingGroup, companyUuid, updated_by uuid.UUID) error
}

// New repository for tech_info_applications.
func New(db *gorm.DB) Repository {
	repo := &sqlRepository{db}

	return repo
}
