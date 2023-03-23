package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/repository"
)

type Service interface {
	CreateTechInfoExternalInfra(ctx context.Context, companyUuid, userUuid *uuid.UUID, techInfoExternalInfra *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error)
	UpdateTechInfoExternalInfra(ctx context.Context, companyUuid, userUuid, techInfoExternalInfraUuid *uuid.UUID, techInfoExternalInfra *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error)
	GetTechInfoExternalInfraById(ctx context.Context, companyUuid, userUuid, techInfoExternalInfraUuid *uuid.UUID) (*entities.TechInfoExternalInfra, error)
	GetAllTechInfoExternalInfras(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.TechInfoExternalInfra, error)
	DeleteTechInfoExternalInfra(ctx context.Context, companyUuid, userUuid, techInfoExternalInfraUuid *uuid.UUID) error
}

type service struct {
	repo repository.Repository
}

func (s *service) CreateTechInfoExternalInfra(ctx context.Context, companyUuid, userUuid *uuid.UUID, techInfoExternalInfra *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error) {
	techInfoExternalInfra.CompanyUuid = *companyUuid
	techInfoExternalInfra.CreatedBy = *userUuid

	return s.repo.CreateTechInfoExternalInfra(ctx, techInfoExternalInfra)
}

func (s *service) UpdateTechInfoExternalInfra(ctx context.Context, companyUuid, userUuid, techInfoExternalInfraUuid *uuid.UUID, techInfoExternalInfra *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error) {
	techInfoExternalInfra.CompanyUuid = *companyUuid
	techInfoExternalInfra.UpdatedBy = *userUuid
	techInfoExternalInfra.TechInfoExternalInfraUuid = *techInfoExternalInfraUuid

	return s.repo.UpdateTechInfoExternalInfra(ctx, techInfoExternalInfra)
}

func (s *service) GetTechInfoExternalInfraById(ctx context.Context, companyUuid, userUuid, techInfoExternalInfraUuid *uuid.UUID) (*entities.TechInfoExternalInfra, error) {
	return s.repo.GetTechInfoExternalInfraById(ctx, techInfoExternalInfraUuid)
}

func (s *service) GetAllTechInfoExternalInfras(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.TechInfoExternalInfra, error) {
	return s.repo.GetAllTechInfoExternalInfras(ctx, companyUuid)
}

func (s *service) DeleteTechInfoExternalInfra(ctx context.Context, companyUuid, userUuid, techInfoExternalInfraUuid *uuid.UUID) error {
	_, err := s.repo.DeleteTechInfoExternalInfra(ctx, techInfoExternalInfraUuid)
	return err
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
