package service

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/repository"

	"github.com/google/uuid"
)

type Service interface {
	CreateWebsite(ctx context.Context, companyUuid, userUuid *uuid.UUID, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error)
	UpdateWebsite(ctx context.Context, companyUuid, userUuid, websiteUuid *uuid.UUID, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error)
	GetWebsiteById(ctx context.Context, companyUuid, userUuid, websiteUuid *uuid.UUID) (*entities.CompanyWebsite, error)
	GetAllWebsites(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.CompanyWebsite, error)
	DeleteWebsite(ctx context.Context, companyUuid, userUuid, websiteUuid *uuid.UUID) error
}

type service struct {
	repo repository.Repository
}

func (s *service) CreateWebsite(ctx context.Context, companyUuid, userUuid *uuid.UUID, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error) {
	website.CompanyUuid = *companyUuid
	website.CreatedBy = *userUuid

	return s.repo.CreateWebsite(ctx, website)
}

func (s *service) UpdateWebsite(ctx context.Context, companyUuid, userUuid, websiteUuid *uuid.UUID, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error) {
	website.CompanyUuid = *companyUuid
	website.UpdatedBy = *userUuid
	website.CompanyWebsiteUuid = *websiteUuid

	return s.repo.UpdateWebsite(ctx, website)
}

func (s *service) GetWebsiteById(ctx context.Context, companyUuid, userUuid, websiteUuid *uuid.UUID) (*entities.CompanyWebsite, error) {
	return s.repo.GetWebsiteById(ctx, websiteUuid)
}

func (s *service) GetAllWebsites(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.CompanyWebsite, error) {
	return s.repo.GetAllWebsites(ctx, companyUuid)
}

func (s *service) DeleteWebsite(ctx context.Context, companyUuid, userUuid, websiteUuid *uuid.UUID) error {
	_, err := s.repo.DeleteWebsite(ctx, websiteUuid)
	return err
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
