package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/service"
)

type Endpoints struct {
	CreateTechInfoWirelessEndpoint      endpoint.Endpoint
	GetTechInfoWirelessByIdEndpoint     endpoint.Endpoint
	GetAllTechInfoWirelesssEndpoint     endpoint.Endpoint
	UpdateTechInfoWirelessEndpoint      endpoint.Endpoint
	DeleteTechInfoWirelessEndpoint      endpoint.Endpoint
	UpdateTechInfoWirelessPatchEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateTechInfoWirelessEndpoint:      makeCreateTechInfoWirelessEndpoint(svc),
		GetTechInfoWirelessByIdEndpoint:     makeGetTechInfoWirelessByIdEndpoint(svc),
		GetAllTechInfoWirelesssEndpoint:     makeGetAllTechInfoWirelesssEndpoint(svc),
		UpdateTechInfoWirelessEndpoint:      makeUpdateTechInfoWirelessEndpoint(svc),
		DeleteTechInfoWirelessEndpoint:      makeDeleteTechInfoWirelessEndpoint(svc),
		UpdateTechInfoWirelessPatchEndpoint: makeUpdateTechInfoWirelessPatchEndpoint(svc),
	}
}

func makeCreateTechInfoWirelessEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateTechInfoWirelessRequest) //nolint:errcheck

		return svc.CreateTechInfoWireless(ctx, &req.CompanyUuid, &req.UserUuid, req.TechInfoWireless)
	}
}
func makeGetTechInfoWirelessByIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetTechInfoWirelessByIdRequest) //nolint:errcheck

		return svc.GetTechInfoWirelessById(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoWirelessUuid)
	}
}

func makeGetAllTechInfoWirelesssEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllTechInfoWirelesssRequest) //nolint:errcheck

		return svc.GetAllTechInfoWirelesss(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
func makeUpdateTechInfoWirelessEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateTechInfoWirelessRequest) //nolint:errcheck

		return svc.UpdateTechInfoWireless(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoWirelessUuid, req.TechInfoWireless)
	}
}
func makeDeleteTechInfoWirelessEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteTechInfoWirelessRequest)
		err := svc.DeleteTechInfoWireless(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoWirelessUuid)
		return "", err
	}
}

func makeUpdateTechInfoWirelessPatchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateTechInfoWirelessPatchRequest) //nolint:errcheck

		return svc.UpdateTechInfoWirelessPatch(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoWirelessUuid, req.PatchRequestBody)
	}
}
