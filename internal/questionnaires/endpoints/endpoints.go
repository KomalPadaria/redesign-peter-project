package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	GetQuestionnairesEndpoint           endpoint.Endpoint
	GetQuestionnairesByCategoryEndpoint endpoint.Endpoint
	PostAnswerEndpoint                  endpoint.Endpoint
	EngineerFeedbackEndpoint            endpoint.Endpoint
	SubmitQuestionnaires                endpoint.Endpoint
	AddAnswerWithEvidenceEndpoint       endpoint.Endpoint
	UpdateAnswerWithEvidenceEndpoint    endpoint.Endpoint
	DownloadEvidenceEndpoint            endpoint.Endpoint
	DeleteEvidenceEndpoint              endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetQuestionnairesEndpoint:           makeGetQuestionnairesEndpoint(svc),
		GetQuestionnairesByCategoryEndpoint: makeGetQuestionnairesByCategoryEndpoint(svc),
		PostAnswerEndpoint:                  makeAddAnswerEndpoint(svc),
		EngineerFeedbackEndpoint:            makeAddEngineerFeedbackEndpoint(svc),
		SubmitQuestionnaires:                makeSubmitQuestionnairesEndpoint(svc),
		AddAnswerWithEvidenceEndpoint:       makeAddAnswerWithEvidenceEndpoint(svc),
		UpdateAnswerWithEvidenceEndpoint:    makeUpdateAnswerWithEvidenceEndpoint(svc),
		DownloadEvidenceEndpoint:            makeDownloadEvidenceEndpoint(svc),
		DeleteEvidenceEndpoint:              makeDeleteEvidenceEndpoint(svc),
	}
}

func makeGetQuestionnairesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetQuestionnairesRequest) //nolint:errcheck

		return svc.GetCategories(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeGetQuestionnairesByCategoryEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetQuestionnairesByCategoryRequest) //nolint:errcheck

		return svc.GetQuestionnairesByCategory(ctx, &req.CompanyUuid, &req.UserUuid, req.Category)
	}
}

func makeAddAnswerEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.PostAnswerRequest) //nolint:errcheck
		err := svc.PostAnswerEndpoint(ctx, &req.CompanyUuid, &req.UserUuid, req.Answers)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func makeAddEngineerFeedbackEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.EngineerFeedbackRequest) //nolint:errcheck
		err := svc.AddEngineerFeedback(ctx, req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
func makeSubmitQuestionnairesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.SubmitQuestionnairesRequest) //nolint:errcheck
		err := svc.SubmitQuestionnaires(ctx, req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
func makeAddAnswerWithEvidenceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.AnswerWithEvidenceRequest) //nolint:errcheck
		err := svc.AddAnswerWithEvidence(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func makeUpdateAnswerWithEvidenceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateAnswerWithEvidenceRequest) //nolint:errcheck
		err := svc.UpdateAnswerWithEvidence(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func makeDownloadEvidenceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DownloadEvidenceRequest) //nolint:errcheck
		return svc.DownloadQuestionnairesEvidence(ctx, req)
	}
}

func makeDeleteEvidenceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteEvidenceRequest) //nolint:errcheck
		err := svc.DeleteQuestionnairesEvidence(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
