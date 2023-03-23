package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	GetFrameworksEndpoint        endpoint.Endpoint
	GetFrameworkControlsEndpoint endpoint.Endpoint
	GetFrameworkStatsEndpoint    endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetFrameworksEndpoint:        makeGetFrameworksEndpoint(svc),
		GetFrameworkControlsEndpoint: makeGetFrameworkControlsEndpoint(svc),
		GetFrameworkStatsEndpoint:    makeGetFrameworkStatsEndpoint(svc),
	}
}

func makeGetFrameworksEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetFrameworksRequest) //nolint:errcheck

		return svc.GetFrameworks(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeGetFrameworkControlsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetFrameworkControlRequest) //nolint:errcheck

		return svc.GetFrameworkControls(ctx, &req.CompanyUuid, &req.UserUuid, &req.FrameworkUuid)
	}
}

func makeGetFrameworkStatsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetFrameworkStatsRequest) //nolint:errcheck

		return svc.GetFrameworkStats(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
