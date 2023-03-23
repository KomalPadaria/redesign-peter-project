package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/service"
)

type Endpoints struct {
	CreateTechInfoExternalInfraEndpoint  endpoint.Endpoint
	GetTechInfoExternalInfraByIdEndpoint endpoint.Endpoint
	GetAllTechInfoExternalInfrasEndpoint endpoint.Endpoint
	UpdateTechInfoExternalInfraEndpoint  endpoint.Endpoint
	DeleteTechInfoExternalInfraEndpoint  endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateTechInfoExternalInfraEndpoint:  makeCreateTechInfoExternalInfraEndpoint(svc),
		GetTechInfoExternalInfraByIdEndpoint: makeGetTechInfoExternalInfraByIdEndpoint(svc),
		GetAllTechInfoExternalInfrasEndpoint: makeGetAllTechInfoExternalInfrasEndpoint(svc),
		UpdateTechInfoExternalInfraEndpoint:  makeUpdateTechInfoExternalInfraEndpoint(svc),
		DeleteTechInfoExternalInfraEndpoint:  makeDeleteTechInfoExternalInfraEndpoint(svc),
	}
}

func makeCreateTechInfoExternalInfraEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateTechInfoExternalInfraRequest) //nolint:errcheck

		return svc.CreateTechInfoExternalInfra(ctx, &req.CompanyUuid, &req.UserUuid, req.TechInfoExternalInfra)
	}
}
func makeGetTechInfoExternalInfraByIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetTechInfoExternalInfraByIdRequest) //nolint:errcheck

		return svc.GetTechInfoExternalInfraById(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoExternalInfraUuid)
	}
}

func makeGetAllTechInfoExternalInfrasEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllTechInfoExternalInfrasRequest) //nolint:errcheck

		return svc.GetAllTechInfoExternalInfras(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
func makeUpdateTechInfoExternalInfraEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateTechInfoExternalInfraRequest) //nolint:errcheck

		return svc.UpdateTechInfoExternalInfra(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoExternalInfraUuid, req.TechInfoExternalInfra)
	}
}
func makeDeleteTechInfoExternalInfraEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteTechInfoExternalInfraRequest)
		err := svc.DeleteTechInfoExternalInfra(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoExternalInfraUuid)
		return "", err
	}
}
