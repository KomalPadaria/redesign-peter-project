package company

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/service"
)

type Client interface {
	CreateCompany(ctx context.Context, user *entities.UsersCompany) (uuid.UUID, error)
	FindByExternalId(ctx context.Context, req *entities.FindCompanyByExternalIdRequest) (*entities.FindCompanyByExternalIdResponse, error)
	GetUserCompaniesByUserUuid(ctx context.Context, id *entities.GetCompaniesByUserIdRequest) (*entities.GetCompaniesByUserIdResponse, error)
	FindByUUID(ctx context.Context, companyUuid *entities.GetCompanyByIdRequest) (*entities.Company, error)
	UpdateCompany(ctx context.Context, updateData *entities.UpdateCompanyRequest) error
	GetAllCompanies(ctx context.Context, req *entities.GetAllCompaniesRequest) ([]*entities.Company, error)

	CreateCompanySubscription(ctx context.Context, cs *entities.CompanySubscription) error
	UpdateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string, cs *entities.CompanySubscription) error
	DeleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string) error
	GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
	GetSubscriptionByName(ctx context.Context, companyUuid *uuid.UUID, subscriptionName string) (*entities.CompanySubscription, error)
	FetchCompanySubscriptionFromSF(ctx context.Context, companyUuid *uuid.UUID) error
	GetConsultingHoursSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
	GetServiceReviewSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
}

func NewClient(ep *endpoints.Endpoints, svc service.Service) Client {
	return &localClient{ep, svc}
}

type localClient struct {
	ep  *endpoints.Endpoints
	svc service.Service
}

func (l localClient) GetServiceReviewSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	return l.svc.GetServiceReviewSubscriptions(ctx, companyUuid)
}

func (l localClient) GetSubscriptionByName(ctx context.Context, companyUuid *uuid.UUID, subscriptionName string) (*entities.CompanySubscription, error) {
	return l.svc.GetSubscriptionByName(ctx, companyUuid, subscriptionName)
}

func (l localClient) GetConsultingHoursSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	return l.svc.GetConsultingHoursSubscriptions(ctx, companyUuid)
}

func (l localClient) FetchCompanySubscriptionFromSF(ctx context.Context, companyUuid *uuid.UUID) error {
	return l.svc.FetchCompanySubscriptionFromSF(ctx, companyUuid)
}

func (l localClient) GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	return l.svc.GetAllSubscriptionsByCompany(ctx, companyUuid)
}

func (l localClient) UpdateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string, cs *entities.CompanySubscription) error {
	return l.svc.UpdateCompanySubscription(ctx, companyUuid, subscriptionId, cs)
}

func (l localClient) DeleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string) error {
	return l.svc.DeleteCompanySubscription(ctx, companyUuid, subscriptionId)
}

func (l localClient) CreateCompanySubscription(ctx context.Context, cs *entities.CompanySubscription) error {
	return l.svc.CreateCompanySubscription(ctx, cs)
}

func (l localClient) CreateCompany(ctx context.Context, req *entities.UsersCompany) (uuid.UUID, error) {
	res, err := l.ep.CreateCompanyEndpoint(ctx, req)

	r, _ := res.(uuid.UUID) //nolint:errcheck

	return r, err
}

func (l localClient) FindByExternalId(ctx context.Context, req *entities.FindCompanyByExternalIdRequest) (*entities.FindCompanyByExternalIdResponse, error) {
	res, err := l.ep.FindByExternalIdEndpoint(ctx, req)

	r, _ := res.(*entities.FindCompanyByExternalIdResponse) //nolint:errcheck

	return r, err
}

func (l localClient) GetUserCompaniesByUserUuid(ctx context.Context, req *entities.GetCompaniesByUserIdRequest) (*entities.GetCompaniesByUserIdResponse, error) {
	res, err := l.ep.GetUserCompaniesByUserUuidEndpoint(ctx, req)

	r, _ := res.(*entities.GetCompaniesByUserIdResponse) //nolint:errcheck

	return r, err
}

func (l localClient) FindByUUID(ctx context.Context, companyUuid *entities.GetCompanyByIdRequest) (*entities.Company, error) {
	res, err := l.ep.GetCompanyEndpoint(ctx, companyUuid)

	r, _ := res.(*entities.Company)

	return r, err
}

func (l localClient) UpdateCompany(ctx context.Context, updateData *entities.UpdateCompanyRequest) error {
	_, err := l.ep.UpdateCompanyEndpoint(ctx, updateData)
	return err
}

func (l localClient) GetAllCompanies(ctx context.Context, req *entities.GetAllCompaniesRequest) ([]*entities.Company, error) {
	res, err := l.ep.GetAllCompanies(ctx, req)
	r, _ := res.([]*entities.Company) //nolint:errcheck

	return r, err
}
