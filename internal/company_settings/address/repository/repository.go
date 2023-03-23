package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetAddressById(ctx context.Context, addressUUID *uuid.UUID) (*entities.CompanyAddress, error)
	GetAddresses(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.CompanyAddressW, error)
	CreateAddress(ctx context.Context, address *entities.CompanyAddress) (*entities.CompanyAddress, error)
	UpdateAddress(ctx context.Context, address *entities.CompanyAddress) (*entities.CompanyAddress, error)
	DeleteAddress(ctx context.Context, addressUuid *uuid.UUID) (*uuid.UUID, error)
	UpdateCompanyAddressPatch(ctx context.Context, userUUID, addressUuid *uuid.UUID, req *entities.UpdateCompanyAddressPatchRequestBody) error
	UpdateCompanyFacilityStatusByAddress(ctx context.Context, userUUID, addressUuid *uuid.UUID, status string) error

	GetFacilities(ctx context.Context, companyUUID *uuid.UUID, query, status string) ([]*CompanyFacility, error)
	GetFacilitiesByAddress(ctx context.Context, addressUuid *uuid.UUID) ([]*CompanyFacility, error)
}

// New repository for websites.
func New(db *sql.DB, gormDb *gorm.DB) Repository {
	repo := &sqlRepository{db, gormDb}

	return repo
}
