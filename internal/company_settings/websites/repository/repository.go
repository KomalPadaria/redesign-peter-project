package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/entities"
)

// Repository for websites.
type Repository interface {
	CreateWebsite(ctx context.Context, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error)
	UpdateWebsite(ctx context.Context, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error)
	GetWebsiteById(ctx context.Context, websiteUuid *uuid.UUID) (*entities.CompanyWebsite, error)
	GetAllWebsites(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.CompanyWebsite, error)
	DeleteWebsite(ctx context.Context, websiteUuid *uuid.UUID) (*uuid.UUID, error)
}

// New repository for websites.
func New(db *sql.DB) Repository {
	repo := &sqlRepository{db}

	return repo
}
