package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/remediation/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/remediation/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	ListTopRemediationEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		ListTopRemediationEndpoint: makeListRemmediationEndpointEndpoint(svc),
	}
}

func makeListRemmediationEndpointEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.ListTopRemediationRequest) //nolint:errcheck

		return svc.ListRemediation(ctx, req)
	}
}
