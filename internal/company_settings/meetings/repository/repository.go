package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.Meetings, error)
	GetCompanyMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetCompanyMeetingsResponse, error)
	CreateOrUpdateCompanyMeeting(ctx context.Context, cMeeting *entities.CompanyMeetings) error
	GetMeetingByUuid(ctx context.Context, meetingUUID *uuid.UUID) (*entities.Meetings, error)
	CreateMeeting(ctx context.Context, meeting *entities.Meetings) error
}

// New repository for websites.
func New(db *sql.DB, gormDB *gorm.DB) Repository {
	repo := &sqlRepository{db, gormDB}

	return repo
}
