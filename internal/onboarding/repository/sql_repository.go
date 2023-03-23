package repository

import (
	"context"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
}

func (s *sqlRepository) UpdateOnboardingStatus(ctx context.Context, onboardingGroup *companyEntities.OnboardingGroup, companyUuid, updated_by uuid.UUID) error {
	err := s.gormDB.WithContext(ctx).Table("public.companies").Where("company_uuid = ?", companyUuid).Update("onboarding", onboardingGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlRepository) GetOnboardingStatus(ctx context.Context, companyUuid uuid.UUID) (*companyEntities.OnboardingGroup, error) {
	type onboarding struct {
		Onboarding companyEntities.OnboardingGroup `json:"onboarding,omitempty"`
	}

	result := onboarding{}
	err := s.gormDB.WithContext(ctx).Table("public.companies").
		Select("onboarding").Where("company_uuid = ?", companyUuid).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result.Onboarding, nil
}
