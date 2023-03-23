package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"gorm.io/gorm"
)

type Repository interface {
	GetFrameworks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.Framework, error)
	GetFrameworkControls(ctx context.Context, companyUuid, userUuid, frameworkUuid *uuid.UUID) ([]*entities.FrameworkControl, error)
	CreateFrameworkControlRemediations(ctx context.Context, companyUUID, userUUID *uuid.UUID) error
	GetFrameworkStats(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetFrameworkStatsResponse, error)
	CreateCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error
	GetCompanyByUUID(ctx context.Context, companyUuid *uuid.UUID) (*companyEntities.Company, error)
	DeleteCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error
	CreateCompanySubscription(ctx context.Context, cs *companyEntities.CompanySubscription) error
}

// New repository for websites.
func New(gormDB *gorm.DB, sqlDB *sql.DB) Repository {
	repo := &sqlRepository{gormDB, sqlDB}

	return repo
}
