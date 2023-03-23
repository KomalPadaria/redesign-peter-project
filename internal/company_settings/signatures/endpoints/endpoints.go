package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/service"
)

type Endpoints struct {
	GetSignaturesEndpoint          endpoint.Endpoint
	UpdateStatusEndpoint           endpoint.Endpoint
	ViewSignaturesDocumentEndpoint endpoint.Endpoint
	WebhookEndpoint                endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetSignaturesEndpoint:          makeGetSignaturesEndpoint(svc),
		UpdateStatusEndpoint:           makeUpdateStatusEndpoint(svc),
		ViewSignaturesDocumentEndpoint: makeViewSignaturesDocumentEndpointv(svc),
		WebhookEndpoint:                makeWebhookEndpoint(svc),
	}
}

func makeGetSignaturesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetSignaturesRequest) //nolint:errcheck

		return svc.GetSignatures(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeUpdateStatusEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateStatusRequest) //nolint:errcheck

		return svc.UpdateStatus(ctx, &req.CompanyUuid, &req.UserUuid, &req.CompanySignatureUuid, req.PatchRequestBody)
	}
}

func makeViewSignaturesDocumentEndpointv(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.ViewDocumentRequest) //nolint:errcheck

		return svc.ViewDocument(ctx, &req.CompanyUuid, &req.SignatureUuid)
	}
}

func makeWebhookEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DocusignWebhookData) //nolint:errcheck
		err := svc.Webhook(ctx, req)
		return nil, err
	}
}
