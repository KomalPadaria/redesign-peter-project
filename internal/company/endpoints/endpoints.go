package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	CreateCompanyEndpoint               endpoint.Endpoint
	FindByExternalIdEndpoint            endpoint.Endpoint
	GetUserCompaniesByUserUuidEndpoint  endpoint.Endpoint
	GetCompanyEndpoint                  endpoint.Endpoint
	UpdateCompanyEndpoint               endpoint.Endpoint
	GetCompanyInfoEndpoint              endpoint.Endpoint
	GetAllCompanies                     endpoint.Endpoint
	UploadSecurityCampaignUsersEndpoint endpoint.Endpoint
	GetSecurityCampaignUsersEndpoint    endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateCompanyEndpoint:               makeCreateCompanyEndpoint(svc),
		FindByExternalIdEndpoint:            makeFindByExternalIdEndpoint(svc),
		GetUserCompaniesByUserUuidEndpoint:  makeGetCompaniesByUserUuidEndpoint(svc),
		GetCompanyEndpoint:                  makeGetCompanyByUuidEndpoint(svc),
		UpdateCompanyEndpoint:               makeUpdateCompanyEndpoint(svc),
		GetCompanyInfoEndpoint:              makeGetCompanyInfoEndpoint(svc),
		GetAllCompanies:                     makeGetAllCompaniesEndpoint(svc),
		UploadSecurityCampaignUsersEndpoint: makeUploadSecurityCampaignUsersEndpoint(svc),
		GetSecurityCampaignUsersEndpoint:    makeGetSecurityCampaignUsersEndpoint(svc),
	}
}

func makeCreateCompanyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UsersCompany) //nolint:errcheck

		return svc.CreateCompany(ctx, req)
	}
}

func makeFindByExternalIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.FindCompanyByExternalIdRequest) //nolint:errcheck

		return svc.FindByExternalId(ctx, req)
	}
}

func makeGetCompaniesByUserUuidEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompaniesByUserIdRequest) //nolint:errcheck
		resp, err := svc.GetUserCompaniesByUserUuid(ctx, req.UserUuid, req.Keyword)

		return resp, err
	}
}

func makeGetCompanyByUuidEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompanyByIdRequest) //nolint:errcheck
		resp, err := svc.FindByUUID(ctx, &req.CompanyUuid)

		return resp, err
	}
}

func makeGetCompanyInfoEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uuid.UUID) //nolint:errcheck
		resp, err := svc.GetCompanyInfo(ctx, req)

		return resp, err
	}
}

func makeUpdateCompanyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateCompanyRequest) //nolint:errcheck
		err := svc.UpdateCompany(ctx, req)

		return nil, err
	}
}

func makeGetAllCompaniesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllCompaniesRequest) //nolint:errcheck
		resp, err := svc.GetAllCompanies(ctx, req.Keyword)

		return resp, err
	}
}

func makeUploadSecurityCampaignUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UploadSecurityCampaignUsersRequest) //nolint:errcheck
		err := svc.UploadSecurityCampaignUsers(ctx, req)

		return nil, err
	}
}

func makeGetSecurityCampaignUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCampaignUsersRequest) //nolint:errcheck
		return svc.GetSecurityCampaignUsers(ctx, req)
	}
}
