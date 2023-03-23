package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"gorm.io/gorm"
)

// Repository for user.
type Repository interface {
	CreateCompany(ctx context.Context, company *entities.UsersCompany) (uuid.UUID, error)
	FindByExternalId(ctx context.Context, externalId string) (*entities.UsersCompany, error)
	GetUserCompaniesByUserUuid(ctx context.Context, userUUID uuid.UUID, keyword string) ([]*entities.UsersCompany, error)
	FindByUUID(ctx context.Context, companyUuid *uuid.UUID) (*entities.Company, error)
	UpdateCompany(ctx context.Context, updateData *entities.UpdateCompanyRequest) error
	GetAllCompanies(ctx context.Context, keyword string) ([]*entities.Company, error)
	CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error

	CreateCompanySubscription(ctx context.Context, cs *entities.CompanySubscription) error
	UpdateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string, cs *entities.CompanySubscription) error
	DeleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string) (*entities.CompanySubscription, error)
	GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
	GetSubscriptionByName(ctx context.Context, companyUuid *uuid.UUID, subscriptionName string) (*entities.CompanySubscription, error)
}

// New repository for company.
func New(db *sql.DB, gormdb *gorm.DB) Repository {
	repo := &sqlRepository{db, gormdb}

	return repo
}
