package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/entities"
	"gorm.io/gorm"
)

// Repository for wireless assessment.
type Repository interface {
	CreateTechInfoWireless(ctx context.Context, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error)
	UpdateTechInfoWireless(ctx context.Context, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error)
	GetTechInfoWirelessById(ctx context.Context, wirelessUuid *uuid.UUID) (*entities.TechInfoWireless, error)
	GetAllTechInfoWirelesss(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.TechInfoWireless, error)
	DeleteTechInfoWireless(ctx context.Context, wirelessUuid *uuid.UUID) (*uuid.UUID, error)
	UpdateTechInfoWirelessPatch(ctx context.Context, userUUID, wirelessUuid *uuid.UUID, req *entities.UpdateTechInfoWirelessPatchRequestBody) error
	UpdateTechInfoWirelessStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error
}

// New repository for websites.
func New(db *sql.DB, gormDB *gorm.DB) Repository {
	repo := &sqlRepository{db, gormDB}

	return repo
}
