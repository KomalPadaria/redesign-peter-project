package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/service"
)

type Endpoints struct {
	GetAllPoliciesEndpoint             endpoint.Endpoint
	CreatePolicyEndpoint               endpoint.Endpoint
	GetPolicyDocumentEndpoint          endpoint.Endpoint
	SaveDocumentEndpoint               endpoint.Endpoint
	GetDocumentEndpoint                endpoint.Endpoint
	GetPolicyHistoriesByPolicyEndpoint endpoint.Endpoint
	DeletePolicyEndpoint               endpoint.Endpoint
	UpdatePolicyDocumentStatusEndpoint endpoint.Endpoint
	GetPoliciesStatsEndpoint           endpoint.Endpoint
	GetTemplatesEndpoint               endpoint.Endpoint
	CreateDocumentFromTemplateEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetAllPoliciesEndpoint:             makeGetAllPoliciesEndpoint(svc),
		CreatePolicyEndpoint:               makeCreatePolicyEndpoint(svc),
		GetPolicyDocumentEndpoint:          makeGetPolicyDocumentEndpoint(svc),
		SaveDocumentEndpoint:               makeSaveDocumentEndpoint(svc),
		GetDocumentEndpoint:                makeGetDocumentEndpoint(svc),
		GetPolicyHistoriesByPolicyEndpoint: makeGetPolicyHistoriesByPolicyEndpoint(svc),
		UpdatePolicyDocumentStatusEndpoint: makeUpdatePolicyDocumentStatusEndpoint(svc),
		DeletePolicyEndpoint:               makeDeletePolicyEndpoint(svc),
		GetPoliciesStatsEndpoint:           makeGetPoliciesStatsEndpoint(svc),
		GetTemplatesEndpoint:               makeGetTemplatesEndpoint(svc),
		CreateDocumentFromTemplateEndpoint: makeCreateDocumentFromTemplateEndpoint(svc),
	}
}

func makeGetAllPoliciesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetAllPoliciesRequest) //nolint:errcheck

		return svc.GetAllPolicies(ctx, req)
	}
}

func makeCreatePolicyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreatePolicyRequest) //nolint:errcheck

		return svc.CreatePolicy(ctx, &req.CompanyUuid, &req.UserUuid, req.Policy)
	}
}

func makeGetPolicyDocumentEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPolicyDocumentRequest) //nolint:errcheck

		return svc.GetPolicyDocument(ctx, &req.CompanyUuid, &req.UserUuid, &req.PolicyUuid, req.Version)
	}
}

func makeSaveDocumentEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.SaveDocumentRequest) //nolint:errcheck

		return svc.SaveDocument(ctx, req)
	}
}

func makeGetDocumentEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPolicyDocumentRequest) //nolint:errcheck

		return svc.GetDocument(ctx, req)
	}
}

func makeGetPolicyHistoriesByPolicyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPolicyHistoriesByPolicyRequest) //nolint:errcheck

		return svc.GetPolicyHistoriesByPolicy(ctx, req)
	}
}

func makeUpdatePolicyDocumentStatusEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdatePolicyDocumentStatus) //nolint:errcheck

		return svc.UpdatePolicyStatus(ctx, &req.CompanyUuid, &req.UserUuid, &req.PolicyUuid, req.PatchRequestBody)
	}
}

func makeDeletePolicyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeletePolicyRequest) //nolint:errcheck
		err := svc.DeletePolicy(ctx, &req.CompanyUuid, &req.UserUuid, &req.PolicyUuid)
		return "", err
	}
}

func makeGetPoliciesStatsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetPoliciesStatsRequest) //nolint:errcheck

		return svc.GetPoliciesStats(ctx, &req.CompanyUuid)
	}
}

func makeGetTemplatesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetTemplatesRequest) //nolint:errcheck

		return svc.GetTemplates(ctx, req)
	}
}

func makeCreateDocumentFromTemplateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateDocumentFromTemplateRequest) //nolint:errcheck

		return svc.CreateDocumentFromTemplate(ctx, req)
	}
}
