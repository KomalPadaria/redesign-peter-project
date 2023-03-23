package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/service"
)

type Endpoints struct {
	CreateApplicationEndpoint      endpoint.Endpoint
	GetApplicationByIdEndpoint     endpoint.Endpoint
	GetAllApplicationsEndpoint     endpoint.Endpoint
	UpdateApplicationEndpoint      endpoint.Endpoint
	UpdateApplicationPatchEndpoint endpoint.Endpoint
	DeleteApplicationEndpoint      endpoint.Endpoint

	CreateApplicationEnvEndpoint      endpoint.Endpoint
	UpdateApplicationEnvEndpoint      endpoint.Endpoint
	UpdateApplicationEnvPatchEndpoint endpoint.Endpoint
	DeleteApplicationEnvEndpoint      endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateApplicationEndpoint:      makeCreateApplicationEndpoint(svc),
		GetApplicationByIdEndpoint:     makeGetApplicationByIdEndpoint(svc),
		GetAllApplicationsEndpoint:     makeGetAllApplicationsEndpoint(svc),
		UpdateApplicationEndpoint:      makeUpdateApplicationEndpoint(svc),
		DeleteApplicationEndpoint:      makeDeleteApplicationEndpoint(svc),
		UpdateApplicationPatchEndpoint: makeUpdateApplicationPatchEndpoint(svc),

		CreateApplicationEnvEndpoint:      makeCreateApplicationEnvEndpoint(svc),
		UpdateApplicationEnvEndpoint:      makeUpdateApplicationEnvEndpoint(svc),
		UpdateApplicationEnvPatchEndpoint: makeUpdateApplicationEnvPatchEndpoint(svc),
		DeleteApplicationEnvEndpoint:      makeDeleteApplicationEnvEndpoint(svc),
	}
}

func makeCreateApplicationEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateApplicationRequest) //nolint:errcheck

		return svc.CreateApplication(ctx, &req.CompanyUuid, &req.UserUuid, req.TechInfoApplication)
	}
}

func makeGetApplicationByIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetApplicationByIdRequest) //nolint:errcheck

		return svc.GetApplicationById(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid)
	}
}

func makeGetAllApplicationsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllApplicationsRequest) //nolint:errcheck

		return svc.GetAllApplications(ctx, req)
	}
}

func makeUpdateApplicationEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateApplicationRequest) //nolint:errcheck

		return svc.UpdateApplication(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid, req.TechInfoApplication)
	}
}

func makeDeleteApplicationEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteApplicationRequest)
		err := svc.DeleteApplication(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid)
		return "", err
	}
}

func makeUpdateApplicationPatchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateApplicationPatchRequest) //nolint:errcheck

		return svc.UpdateApplicationPatch(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid, req.PatchRequestBody)
	}
}

func makeCreateApplicationEnvEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateApplicationEnvRequest) //nolint:errcheck

		return svc.CreateApplicationEnv(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid, req.ApplicationEnv)
	}
}

func makeUpdateApplicationEnvEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateApplicationEnvRequest) //nolint:errcheck

		return svc.UpdateApplicationEnv(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid, &req.ApplicationEnvUuid, req.ApplicationEnv)
	}
}

func makeUpdateApplicationEnvPatchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateApplicationEnvPatchRequest) //nolint:errcheck

		return svc.UpdateApplicationEnvPatch(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid, &req.ApplicationEnvUuid, req.PatchRequestBody)
	}
}

func makeDeleteApplicationEnvEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteApplicationEnvRequest)
		err := svc.DeleteApplicationEnv(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoApplicationUuid, &req.ApplicationEnvUuid)
		return "", err
	}
}
