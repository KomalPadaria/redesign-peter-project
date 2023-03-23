package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetAllTechInfoIpRange(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.TechInfoIpRanges, error)
	CreateTechInfoIpRange(ctx context.Context, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error)
	UpdateTechInfoIpRange(ctx context.Context, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error)
	UpdateTechInfoIpRangePatch(ctx context.Context, userUUID, ipRangeUuid *uuid.UUID, req *entities.UpdateTechInfoIpRangePatchRequestBody) error
	DeleteTechInfoIpRange(ctx context.Context, applicationUuid *uuid.UUID) (*uuid.UUID, error)
	UpdateTechInfoIpRangeStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error
	GetTechInfoIpRangeByUuid(ctx context.Context, ipRangeUuid *uuid.UUID) (*entities.TechInfoIpRanges, error)
}

// New repository for tech_info_applications.
func New(db *gorm.DB) Repository {
	repo := &sqlRepository{db}

	return repo
}
