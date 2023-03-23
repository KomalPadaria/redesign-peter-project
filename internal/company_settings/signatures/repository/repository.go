package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetSignatures(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetSignaturesResponse, error)
	UpdateStatus(ctx context.Context, userUUID, companySignatureUuid *uuid.UUID, req *entities.UpdateStatusRequestBody) error
	CreateSignatures(ctx context.Context, companySignature *entities.CompanySignatures) (*entities.CompanySignatures, error)
	GetCompanySignature(ctx context.Context, companyUUID, signatureUuid *uuid.UUID) (*entities.CompanySignatures, error)
}

// New repository for websites.
func New(db *sql.DB, gormDB *gorm.DB) Repository {
	repo := &sqlRepository{db, gormDB}

	return repo
}
