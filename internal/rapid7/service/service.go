package service

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/repository"
)

type Service interface {
	GetLastAndNextVulnerabilityScan(ctx context.Context, siteIds []int) (*entities.LastNexVulnerabilityScan, error)
	GetSiteIds(ctx context.Context, companyUuid uuid.UUID, testType string) ([]int, error)
	GetVulnerabilitySeverityCountsByInterval(ctx context.Context, siteIds []int) ([]*entities.VulnerabilitySeverityCountsByInterval, error)
	GetFactSiteCountsBySiteId(ctx context.Context, siteIds []int) (*entities.FactSiteCounts, error)
	GetTopRemediationTasks(ctx context.Context, siteIds []int) ([]*entities.TopRemediationTask, error)
	GetSolutionSupercedenceSteps(ctx context.Context, solutionId int) ([]string, error)
	GetRemediations(ctx context.Context, siteIds []int, limit int) ([]*entities.Remediation, error)
	FilterSiteIdsByTestType(ctx context.Context, siteIds []int, testType string) ([]int, error)
}

func New(config config.Config, repo repository.Repository, companyClient company.Client) Service {
	return &service{config, repo, companyClient}
}

type service struct {
	config        config.Config
	repo          repository.Repository
	companyClient company.Client
}

func (s *service) FilterSiteIdsByTestType(ctx context.Context, siteIds []int, testType string) ([]int, error) {
	return s.repo.FilterSiteIdsByTestType(ctx, siteIds, testType)
}

func (s *service) GetRemediations(ctx context.Context, siteIds []int, limit int) ([]*entities.Remediation, error) {
	return s.repo.GetRemediations(ctx, siteIds, limit)
}

func (s *service) GetSolutionSupercedenceSteps(ctx context.Context, solutionId int) ([]string, error) {
	return s.repo.GetSolutionSupercedenceSteps(ctx, solutionId)
}

func (s *service) GetTopRemediationTasks(ctx context.Context, siteIds []int) ([]*entities.TopRemediationTask, error) {
	return s.repo.GetTopRemediationTasks(ctx, siteIds)
}

func (s *service) GetFactSiteCountsBySiteId(ctx context.Context, siteIds []int) (*entities.FactSiteCounts, error) {
	return s.repo.GetFactSiteCountsBySiteId(ctx, siteIds)
}

func (s *service) GetVulnerabilitySeverityCountsByInterval(ctx context.Context, siteIds []int) ([]*entities.VulnerabilitySeverityCountsByInterval, error) {
	return s.repo.GetVulnerabilitySeverityCountsByInterval(ctx, siteIds)

}

func (s *service) GetSiteIds(ctx context.Context, companyUuid uuid.UUID, testType string) ([]int, error) {
	company, err := s.companyClient.FindByUUID(ctx, &companyEntities.GetCompanyByIdRequest{CompanyUuid: companyUuid})
	if err != nil {
		return nil, err
	}

	var ids []int
	for _, id := range company.Rapid7SiteIds {
		i, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}

	siteIds, err := s.repo.FilterSiteIdsByTestType(ctx, ids, testType)
	if err != nil {
		return nil, err
	}

	return siteIds, nil
}

func (s *service) GetLastAndNextVulnerabilityScan(ctx context.Context, siteIds []int) (*entities.LastNexVulnerabilityScan, error) {
	return s.repo.GetLastAndNextVulnerabilityScan(ctx, siteIds)
}
