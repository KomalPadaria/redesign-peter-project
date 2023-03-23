package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/entities"
)

// Repository for websites.
type Repository interface {
	CreateTechInfoExternalInfra(ctx context.Context, website *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error)
	UpdateTechInfoExternalInfra(ctx context.Context, website *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error)
	GetTechInfoExternalInfraById(ctx context.Context, websiteUuid *uuid.UUID) (*entities.TechInfoExternalInfra, error)
	GetAllTechInfoExternalInfras(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.TechInfoExternalInfra, error)
	DeleteTechInfoExternalInfra(ctx context.Context, websiteUuid *uuid.UUID) (*uuid.UUID, error)
}

// New repository for websites.
func New(db *sql.DB) Repository {
	repo := &sqlRepository{db}

	return repo
}
