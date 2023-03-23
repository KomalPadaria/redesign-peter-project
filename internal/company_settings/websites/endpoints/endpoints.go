package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/service"
)

type Endpoints struct {
	CreateWebsiteEndpoint  endpoint.Endpoint
	GetWebsiteByIdEndpoint endpoint.Endpoint
	GetAllWebsitesEndpoint endpoint.Endpoint
	UpdateWebsiteEndpoint  endpoint.Endpoint
	DeleteWebsiteEndpoint  endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateWebsiteEndpoint:  makeCreateWebsiteEndpoint(svc),
		GetWebsiteByIdEndpoint: makeGetWebsiteByIdEndpoint(svc),
		GetAllWebsitesEndpoint: makeGetAllWebsitesEndpoint(svc),
		UpdateWebsiteEndpoint:  makeUpdateWebsiteEndpoint(svc),
		DeleteWebsiteEndpoint:  makeDeleteWebsiteEndpoint(svc),
	}
}

func makeCreateWebsiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateWebsiteRequest) //nolint:errcheck

		return svc.CreateWebsite(ctx, &req.CompanyUuid, &req.UserUuid, req.CompanyWebsite)
	}
}
func makeGetWebsiteByIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetWebsiteByIdRequest) //nolint:errcheck

		return svc.GetWebsiteById(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanyWebsiteUuid)
	}
}

func makeGetAllWebsitesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllWebsitesRequest) //nolint:errcheck

		return svc.GetAllWebsites(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
func makeUpdateWebsiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateWebsiteRequest) //nolint:errcheck

		return svc.UpdateWebsite(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanyWebsiteUuid, req.CompanyWebsite)
	}
}
func makeDeleteWebsiteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteWebsiteRequest)
		err := svc.DeleteWebsite(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanyWebsiteUuid)
		return "", err
	}
}
