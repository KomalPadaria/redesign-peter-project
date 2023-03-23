package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
)

// Repository for user.
type Repository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	UpdateCompanyUserLink(ctx context.Context, user *entities.CompanyUser) (*entities.CompanyUser, error)
	DeleteCompanyUserLink(ctx context.Context, companyUUID, reqUserUUID, userUUID *uuid.UUID) error
	FindByUUID(ctx context.Context, userUUID uuid.UUID) (*entities.User, error)
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	FindByEmail(ctx context.Context, username string) (*entities.User, error)
	FindByExternalId(ctx context.Context, externalId string) (*entities.User, error)
	LinkCompanyAndUser(ctx context.Context, companyUser *entities.CompanyUser) (*entities.CompanyUser, error)
	ActivateOrDeactivateByEmail(ctx context.Context, externalId string, isActive bool) (*entities.User, error)
	UpdateIsFirstLoginByUserUuid(ctx context.Context, userUUID uuid.UUID, isFirstLogin bool) (*entities.User, error)
	UpdateCurrentCompany(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) error
	GetCompanyUsers(ctx context.Context, companyUUID, userUUID uuid.UUID) ([]*entities.User, error)
	UpdateUserCompanyStatus(ctx context.Context, userUUID uuid.UUID, status string) error
	UpdateUserDetails(ctx context.Context, userUUID uuid.UUID, details map[string]interface{}) error
	GetCompanyUser(ctx context.Context, userUUID, comapnyUUID uuid.UUID) (*entities.CompanyUser, error)
	GetCompanyUserByUserId(ctx context.Context, userUUID uuid.UUID) ([]*entities.CompanyUser, error)
}

// New repository for user.
func New(db *sql.DB) Repository {
	repo := &sqlRepository{db}

	return repo
}
