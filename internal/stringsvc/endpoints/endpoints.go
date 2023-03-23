// Package endpoints contains app endpoints
package endpoints

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/entities"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/service"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct
type Endpoints struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
}

// New endpoint constructor
func New(svc service.Service) *Endpoints {
	endpoints := &Endpoints{
		UppercaseEndpoint: makeUppercase(svc),
		CountEndpoint:     makeCount(svc),
	}

	return endpoints
}

func makeUppercase(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UppercaseRequest)
		resp := &entities.UppercaseResponse{}

		svcRes := svc.Uppercase(req.Input)

		resp.Output = svcRes

		return resp, nil
	}
}

func makeCount(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CountRequest)
		resp := &entities.CountResponse{}

		svcRes := svc.Count(req.Input)

		resp.Count = svcRes

		return resp, nil
	}
}
