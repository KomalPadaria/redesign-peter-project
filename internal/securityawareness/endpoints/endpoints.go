package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness/service"
)

type Endpoints struct {
	GetPhishingDetailsEndpoint endpoint.Endpoint
	GetTrainingDetailsEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetPhishingDetailsEndpoint: makeGetPhishingDetailsEndpoint(svc),
		GetTrainingDetailsEndpoint: makeGetTrainingDetailsEndpoint(svc),
	}
}

func makeGetPhishingDetailsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPhishingDetailsRequest) //nolint:errcheck

		return svc.GetPhishingDetails(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeGetTrainingDetailsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetTrainingDetailsRequest) //nolint:errcheck

		return svc.GetTrainingDetails(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
