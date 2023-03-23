package repository

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/entities"
)

// Repository for user.
type Repository interface {
	CreateApplication(ctx context.Context, application *entities.TechInfoApplication) (*entities.TechInfoApplication, error)
	UpdateApplication(ctx context.Context, application *entities.TechInfoApplication) (*entities.TechInfoApplication, error)
	UpdateApplicationPatch(ctx context.Context, userUUID, applicationUuid *uuid.UUID, req *entities.UpdateApplicationPatchRequestBody) error
	GetApplicationById(ctx context.Context, applicationUuid *uuid.UUID) (*entities.TechInfoApplication, error)
	GetAllApplications(ctx context.Context, companyUuid *uuid.UUID, keyword string) ([]*entities.TechInfoApplication, error)
	DeleteApplication(ctx context.Context, applicationUuid *uuid.UUID) (*uuid.UUID, error)

	CreateApplicationEnv(ctx context.Context, applicationEnv *entities.ApplicationEnv) (*entities.ApplicationEnv, error)
	UpdateApplicationEnv(ctx context.Context, applicationEnv *entities.ApplicationEnv) (*entities.ApplicationEnv, error)
	GetApplicationEnvById(ctx context.Context, applicationEnvUuid *uuid.UUID) (*entities.ApplicationEnv, error)
	UpdateApplicationEnvPatch(ctx context.Context, userUUID, applicationEnvUuid *uuid.UUID, req *entities.UpdateApplicationEnvPatchRequestBody) error
	DeleteApplicationEnv(ctx context.Context, applicationEnvUuid *uuid.UUID) (*uuid.UUID, error)
}

// New repository for tech_info_applications.
func New(db *sql.DB, gormDB *gorm.DB) Repository {
	repo := &sqlRepository{db, gormDB}

	return repo
}
