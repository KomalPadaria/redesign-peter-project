package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/service"
)

type Endpoints struct {
	GetAllTechInfoIpRangeEndpoint      endpoint.Endpoint
	CreateTechInfoIpRangeEndpoint      endpoint.Endpoint
	UpdateTechInfoIpRangeEndpoint      endpoint.Endpoint
	UpdateTechInfoIpRangePatchEndpoint endpoint.Endpoint
	DeleteTechInfoIpRangeEndpoint      endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetAllTechInfoIpRangeEndpoint:      makeGetAllTechInfoIpsEndpoint(svc),
		CreateTechInfoIpRangeEndpoint:      makeCreateTechInfoIpRangeEndpoint(svc),
		UpdateTechInfoIpRangeEndpoint:      makeUpdateTechInfoIpRangeEndpoint(svc),
		UpdateTechInfoIpRangePatchEndpoint: makeUpdateTechInfoIpRangePatchEndpoint(svc),
		DeleteTechInfoIpRangeEndpoint:      makeDeleteTechInfoIpRangeEndpoint(svc),
	}
}

func makeGetAllTechInfoIpsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllTechInfoIpRangeRequest) //nolint:errcheck

		return svc.GetAllTechInfoIpRange(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeCreateTechInfoIpRangeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateTechInfoIpRangeRequest) //nolint:errcheck

		return svc.CreateTechInfoIpRange(ctx, &req.CompanyUuid, &req.UserUuid, req.TechInfoIpRanges)
	}
}

func makeUpdateTechInfoIpRangeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateTechInfoIpRangeRequest) //nolint:errcheck

		return svc.UpdateTechInfoIpRange(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoIpRangeUuid, req.TechInfoIpRanges)
	}
}

func makeUpdateTechInfoIpRangePatchEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateTechInfoIpRangePatchRequest) //nolint:errcheck

		return svc.UpdateTechInfoIpRangePatch(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoIpRangeUuid, req.PatchRequestBody)
	}
}

func makeDeleteTechInfoIpRangeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteTechInfoIpRangeRequest)
		err := svc.DeleteTechInfoIpRange(ctx, &req.CompanyUuid, &req.UserUuid, &req.TechInfoIpRangeUuid)
		return "", err
	}
}
