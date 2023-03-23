package service

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/converter"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type Service interface {
	GetAllPolicies(ctx context.Context, req *entities.GetAllPoliciesRequest) ([]*entities.GetAllPoliciesResponse, error)
	CreatePolicy(ctx context.Context, companyUuid, userUuid *uuid.UUID, policy *entities.Policy) (*entities.Policy, error)
	GetPolicyDocument(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, version int) (*entities.GetPolicyDocumentResponse, error)
	GetDocument(ctx context.Context, req *entities.GetPolicyDocumentRequest) (*entities.GetDocumentResponse, error)
	SaveDocument(ctx context.Context, req *entities.SaveDocumentRequest) (*entities.GetPolicyDocumentResponse, error)
	GetPolicyHistoriesByPolicy(ctx context.Context, req *entities.GetPolicyHistoriesByPolicyRequest) ([]*entities.GetPolicyHistoriesByPolicyResponse, error)
	DeletePolicy(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID) error
	UpdatePolicyStatus(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, req *entities.UpdatePolicyDocumentStatusPatchRequestBody) (*entities.UpdatePolicyDocumentResponse, error)
	GetPoliciesStats(ctx context.Context, companyUuid *uuid.UUID) (*entities.GetPoliciesStatsResponse, error)
	GetTemplates(ctx context.Context, req *entities.GetTemplatesRequest) ([]*entities.GetTemplatesResponse, error)
	CreateDocumentFromTemplate(ctx context.Context, req *entities.CreateDocumentFromTemplateRequest) (*entities.GetPolicyDocumentResponse, error)
}

type service struct {
	repo             repository.Repository
	onboardingClient onboarding.Client
}

func (s *service) CreateDocumentFromTemplate(ctx context.Context, req *entities.CreateDocumentFromTemplateRequest) (*entities.GetPolicyDocumentResponse, error) {
	pt, err := s.repo.GetTemplateByUuid(ctx, &req.PolicyTemplateUuid)
	if err != nil {
		return nil, err
	}

	policyTemplateUuid := nullable.NewNullUUID(pt.PolicyTemplateUuid)
	p, err := s.CreatePolicy(ctx, &req.CompanyUuid, &req.UserUuid, &entities.Policy{Name: pt.Name, PolicyTemplateUuid: *policyTemplateUuid})
	if err != nil {
		return nil, err
	}

	response, err := s.SaveDocument(ctx, &entities.SaveDocumentRequest{
		CompanyUuid: req.CompanyUuid,
		UserUuid:    req.UserUuid,
		PolicyUUID:  p.PolicyUuid,
		SaveDocumentRequestBody: &entities.SaveDocumentRequestBody{
			Name:     pt.Name,
			Document: pt.Document,
		},
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *service) GetTemplates(ctx context.Context, req *entities.GetTemplatesRequest) ([]*entities.GetTemplatesResponse, error) {
	templates := make([]*entities.GetTemplatesResponse, 0)
	pts, err := s.repo.GetTemplates(ctx, req.CompanyType)
	if err != nil {
		return nil, err
	}

	for _, t := range pts {
		templates = append(templates, &entities.GetTemplatesResponse{
			PolicyTemplateUuid: t.PolicyTemplateUuid,
			Name:               t.Name,
			Description:        t.Description,
		})
	}

	return templates, nil
}

func (s *service) GetPolicyHistoriesByPolicy(ctx context.Context, req *entities.GetPolicyHistoriesByPolicyRequest) ([]*entities.GetPolicyHistoriesByPolicyResponse, error) {
	phs, err := s.repo.GetPolicyHistoriesByPolicyUuid(ctx, &req.PolicyUUID)
	if err != nil {
		return nil, err
	}

	res := make([]*entities.GetPolicyHistoriesByPolicyResponse, 0)
	for _, ph := range phs {
		r := &entities.GetPolicyHistoriesByPolicyResponse{
			PolicyHistoryUUID: ph.PolicyHistoryUuid,
			Version:           ph.Version,
			CreatedAt:         ph.CreatedAt.Time,
			Owner:             entities.NewUserInfo(ph.Created.UserUuid, ph.Created.FirstName, ph.Created.LastName, "", ph.Created.Email),
		}
		res = append(res, r)
	}

	return res, nil
}

func (s *service) SaveDocument(ctx context.Context, req *entities.SaveDocumentRequest) (*entities.GetPolicyDocumentResponse, error) {
	ph, err := s.repo.SaveDocument(ctx, req)
	if err != nil {
		return nil, err
	}

	return s.GetPolicyDocument(ctx, &req.CompanyUuid, &req.UserUuid, &ph.PolicyUuid, ph.Version)
}

func (s *service) GetDocument(ctx context.Context, req *entities.GetPolicyDocumentRequest) (*entities.GetDocumentResponse, error) {
	pdoc, err := s.repo.GetDocument(ctx, &req.CompanyUuid, &req.UserUuid, &req.PolicyUuid, req.Version)
	if err != nil {
		return nil, err
	}

	if pdoc.Document != "" {
		file := &converter.File{
			From: converter.HTML,
			To:   converter.DOCX,
			Data: io.NopCloser(strings.NewReader(pdoc.Document)),
		}
		docx, err := file.HTMLToDocx()
		if err != nil {
			return nil, err
		}
		fileName := fmt.Sprintf("%s-v%d%s", strings.ReplaceAll(strings.ToLower(pdoc.Policy.Name), " ", "-"), pdoc.Version, converter.DOCX)
		return &entities.GetDocumentResponse{
			Name:     fileName,
			Document: docx,
		}, nil
	}
	return nil, &appError.ErrNotFound{Message: "policy document not found"}
}

func (s *service) GetPolicyDocument(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, version int) (*entities.GetPolicyDocumentResponse, error) {
	p, err := s.repo.GetPolicyDocument(ctx, companyUuid, policyUuid, version)
	if err != nil {
		return nil, err
	}

	res := &entities.GetPolicyDocumentResponse{
		PolicyUUID:      p.Policy.PolicyUuid,
		Name:            p.Policy.Name,
		Status:          p.Policy.Status,
		StatusUpdatedAt: p.Policy.StatusUpdatedAt,
		StatusUpdatedBy: entities.NewUserInfo(p.Policy.StatusUpdated.UserUuid, p.Policy.StatusUpdated.FirstName, p.Policy.StatusUpdated.LastName, "", p.Policy.StatusUpdated.Email),
		Version:         p.Version,
		Document:        p.Document,
		Owner:           entities.NewUserInfo(p.Created.UserUuid, p.Created.FirstName, p.Created.LastName, "", p.Created.Email),
		CreatedAt:       p.CreatedAt.Time,
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.UploadPoliciesAndProcedures, *companyUuid, *userUuid)

	return res, nil
}

func (s *service) CreatePolicy(ctx context.Context, companyUuid, userUuid *uuid.UUID, policy *entities.Policy) (*entities.Policy, error) {
	policy.CompanyUuid = *companyUuid
	policy.CreatedBy = *userUuid
	policy.StatusUpdatedBy = *userUuid

	return s.repo.CreatePolicy(ctx, policy)
}

func (s *service) GetAllPolicies(ctx context.Context, req *entities.GetAllPoliciesRequest) ([]*entities.GetAllPoliciesResponse, error) {
	return s.repo.GetAllPolicyHistory(ctx, &req.CompanyUuid, req.Keyword)
}

func (s *service) DeletePolicy(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID) error {
	return s.repo.DeletePolicy(ctx, companyUuid, userUuid, policyUuid)
}

func (s *service) UpdatePolicyStatus(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, req *entities.UpdatePolicyDocumentStatusPatchRequestBody) (*entities.UpdatePolicyDocumentResponse, error) {
	err := s.repo.UpdatePolicyStatus(ctx, companyUuid, userUuid, policyUuid, req)
	if err != nil {
		return nil, err
	}

	return &entities.UpdatePolicyDocumentResponse{
		PolicyUuid: *policyUuid,
		Status:     req.Status,
		Comment:    req.Comment,
	}, nil
}

func (s *service) GetPoliciesStats(ctx context.Context, companyUuid *uuid.UUID) (*entities.GetPoliciesStatsResponse, error) {
	return s.repo.GetPoliciesStats(ctx, companyUuid)
}

func New(repo repository.Repository, onboardingClient onboarding.Client) Service {
	return &service{
		repo:             repo,
		onboardingClient: onboardingClient,
	}
}
