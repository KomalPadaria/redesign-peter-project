package service

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/repository"
)

type Service interface {
	GetAddresses(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.CompanyAddressW, error)
	CreateAddress(ctx context.Context, companyUuid, userUuid *uuid.UUID, address *entities.CompanyAddress) (*entities.CompanyAddress, error)
	UpdateAddress(ctx context.Context, companyUuid, userUuid, addressUuid *uuid.UUID, address *entities.CompanyAddress) (*entities.CompanyAddress, error)
	DeleteAddress(ctx context.Context, companyUuid, userUuid, addressUuid *uuid.UUID) error
	UpdateCompanyAddressPatch(ctx context.Context, companyUuid, userUUID, addressUuid *uuid.UUID, req *entities.UpdateCompanyAddressPatchRequestBody) (*entities.UpdateCompanyAddressPatchResponse, error)
	GetFacilities(ctx context.Context, companyUUID, userUUID *uuid.UUID, query, status string) ([]*repository.CompanyFacility, error)
}

type service struct {
	repo             repository.Repository
	ipRangesClient   ipranges.Client
	wirelessClient   wireless.Client
	onboardingClient onboarding.Client
}

func (s *service) GetFacilities(ctx context.Context, companyUUID, userUUID *uuid.UUID, query, status string) ([]*repository.CompanyFacility, error) {
	return s.repo.GetFacilities(ctx, companyUUID, query, status)
}

func (s *service) UpdateCompanyAddressPatch(ctx context.Context, companyUuid, userUUID, addressUuid *uuid.UUID, req *entities.UpdateCompanyAddressPatchRequestBody) (*entities.UpdateCompanyAddressPatchResponse, error) {
	err := s.repo.UpdateCompanyAddressPatch(ctx, userUUID, addressUuid, req)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateCompanyFacilityStatusByAddress(ctx, userUUID, addressUuid, req.Status)
	if err != nil {
		return nil, err
	}

	facilities, err := s.repo.GetFacilitiesByAddress(ctx, addressUuid)
	if err != nil {
		return nil, err
	}

	fIds := make([]uuid.UUID, 0)
	for _, f := range facilities {
		fIds = append(fIds, f.CompanyFacilityUuid)
	}

	err = s.ipRangesClient.UpdateTechInfoIpRangeStatusByFacilities(ctx, userUUID, fIds, req.Status)
	if err != nil {
		return nil, err
	}

	err = s.wirelessClient.UpdateTechInfoWirelessStatusByFacilities(ctx, userUUID, fIds, req.Status)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateCompanyAddressPatchResponse{
		CompanyAddressUuid: *addressUuid,
		Status:             req.Status,
	}, nil
}

func (s *service) DeleteAddress(ctx context.Context, companyUuid, userUuid, addressUuid *uuid.UUID) error {
	_, err := s.repo.DeleteAddress(ctx, addressUuid)
	return err
}

func (s *service) UpdateAddress(ctx context.Context, companyUuid, userUuid, addressUuid *uuid.UUID, address *entities.CompanyAddress) (*entities.CompanyAddress, error) {
	address.UpdatedBy = *userUuid
	address.CompanyAddressUuid = *addressUuid

	address, err := s.repo.UpdateAddress(ctx, address)
	if err != nil {
		return nil, err
	}

	address, err = s.repo.GetAddressById(ctx, &address.CompanyAddressUuid)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *service) CreateAddress(ctx context.Context, companyUuid, userUuid *uuid.UUID, address *entities.CompanyAddress) (*entities.CompanyAddress, error) {
	address.CompanyUuid = *companyUuid
	address.CreatedBy = *userUuid

	for _, f := range address.Facilities {
		f.CompanyUuid = *companyUuid
		f.CreatedBy = *userUuid
	}

	address, err := s.repo.CreateAddress(ctx, address)
	if err != nil {
		return nil, err
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.CompanyInfo, *companyUuid, *userUuid)
	return address, nil
}

func (s *service) GetAddresses(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.CompanyAddressW, error) {
	return s.repo.GetAddresses(ctx, companyUUID, userUUID)
}

func New(repo repository.Repository, ipRangesClient ipranges.Client, wirelessClient wireless.Client, onboardingClient onboarding.Client) Service {
	return &service{
		repo:             repo,
		ipRangesClient:   ipRangesClient,
		wirelessClient:   wirelessClient,
		onboardingClient: onboardingClient,
	}
}
