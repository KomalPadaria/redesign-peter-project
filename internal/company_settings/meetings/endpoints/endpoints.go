package endpoints

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetMeetingsEndpoint               endpoint.Endpoint
	GetCompanyMeetingsEndpoint        endpoint.Endpoint
	CalendlyWebhookEndpoint           endpoint.Endpoint
	CreateMeetingFromCalendlyEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetMeetingsEndpoint:               makeGetMeetingsEndpoint(svc),
		GetCompanyMeetingsEndpoint:        makeGetCompanyMeetingsEndpoint(svc),
		CalendlyWebhookEndpoint:           makeCalendlyWebhookEndpoint(svc),
		CreateMeetingFromCalendlyEndpoint: makeCreateMeetingFromCalendlyEndpoint(svc),
	}
}

func makeGetMeetingsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetMeetingsRequest) //nolint:errcheck

		return svc.GetMeetings(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeGetCompanyMeetingsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompanyMeetingsRequest) //nolint:errcheck

		return svc.GetCompanyMeetings(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeCalendlyWebhookEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.WebhookData) //nolint:errcheck
		err := svc.CalendlyWebhook(ctx, req)
		return nil, err
	}
}

func makeCreateMeetingFromCalendlyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		err := svc.CreateMeetingFromCalendly(ctx)
		return nil, err
	}
}
