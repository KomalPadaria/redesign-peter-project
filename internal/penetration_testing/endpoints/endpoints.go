package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	GetPenetrationTestsEndpoint     endpoint.Endpoint
	GetPenetrationTestStatsEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetPenetrationTestsEndpoint:     makeGetPenetrationTestsEndpoint(svc),
		GetPenetrationTestStatsEndpoint: makeGetPenetrationTestStatsEndpoint(svc),
	}
}

func makeGetPenetrationTestsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPenetrationTestsRequest) //nolint:errcheck

		return svc.GetPenetrationTopRemediationTasks(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeGetPenetrationTestStatsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPenetrationTestsStatsRequest) //nolint:errcheck

		return svc.GetPenetrationStats(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
