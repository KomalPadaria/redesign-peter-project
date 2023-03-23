package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetAllPolicyHistory(ctx context.Context, companyUuid *uuid.UUID, keyword string) ([]*entities.GetAllPoliciesResponse, error)
	CreatePolicy(ctx context.Context, policy *entities.Policy) (*entities.Policy, error)
	GetPolicyDocument(ctx context.Context, companyUuid, policyUuid *uuid.UUID, version int) (*entities.PolicyHistory, error)
	SaveDocument(ctx context.Context, req *entities.SaveDocumentRequest) (*entities.PolicyHistory, error)
	GetDocument(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, version int) (*entities.PolicyHistory, error)
	GetPolicyHistoriesByPolicyUuid(ctx context.Context, policyUuid *uuid.UUID) ([]*entities.PolicyHistory, error)
	DeletePolicy(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID) error
	UpdatePolicyStatus(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, req *entities.UpdatePolicyDocumentStatusPatchRequestBody) error
	GetPoliciesStats(ctx context.Context, companyUuid *uuid.UUID) (*entities.GetPoliciesStatsResponse, error)
	GetTemplates(ctx context.Context, companyTypes []string) ([]*entities.PolicyTemplates, error)
	GetTemplateByUuid(ctx context.Context, policyTemplateUuid *uuid.UUID) (*entities.PolicyTemplates, error)
}

// New repository for tech_info_applications.
func New(db *gorm.DB) Repository {
	repo := &sqlRepository{db}

	return repo
}
