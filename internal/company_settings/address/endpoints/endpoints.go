package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/service"
)

type Endpoints struct {
	GetCompanyAddressesEndpoint       endpoint.Endpoint
	CreateCompanyAddressEndpoint      endpoint.Endpoint
	UpdateCompanyAddressEndpoint      endpoint.Endpoint
	DeleteCompanyAddressEndpoint      endpoint.Endpoint
	UpdateCompanyAddressPatchEndpoint endpoint.Endpoint

	GetFacilitiesEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetCompanyAddressesEndpoint:       makeGetAddressesEndpoint(svc),
		CreateCompanyAddressEndpoint:      makeCreateCompanyAddressEndpoint(svc),
		UpdateCompanyAddressEndpoint:      makeUpdateCompanyAddressEndpoint(svc),
		DeleteCompanyAddressEndpoint:      makeDeleteCompanyAddressEndpoint(svc),
		UpdateCompanyAddressPatchEndpoint: makeUpdateCompanyAddressPatchEndpoint(svc),

		GetFacilitiesEndpoint: makeGetFacilitiesEndpoint(svc),
	}
}

func makeGetAddressesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompanyAddressRequest) //nolint:errcheck

		return svc.GetAddresses(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeCreateCompanyAddressEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateCompanyAddressRequest) //nolint:errcheck

		return svc.CreateAddress(ctx, &req.CompanyUuid, &req.UserUuid, req.CompanyAddress)
	}
}

func makeUpdateCompanyAddressEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateCompanyAddressRequest) //nolint:errcheck

		return svc.UpdateAddress(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanyAddressUuid, req.CompanyAddress)
	}
}

func makeDeleteCompanyAddressEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteCompanyAddressRequest)
		err := svc.DeleteAddress(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanyAddressUuid)
		return "", err
	}
}

func makeUpdateCompanyAddressPatchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateCompanyAddressPatchRequest) //nolint:errcheck

		return svc.UpdateCompanyAddressPatch(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanyAddressUuid, req.PatchRequestBody)
	}
}

func makeGetFacilitiesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompanyFacilitiesRequest) //nolint:errcheck

		return svc.GetFacilities(ctx, &req.CompanyUuid, &req.UserUuid, req.Query, req.Status)
	}
}
