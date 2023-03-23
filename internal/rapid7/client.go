package rapid7

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/service"
)

type Client interface {
	GetLastAndNextVulnerabilityScan(ctx context.Context, siteIds []int) (*entities.LastNexVulnerabilityScan, error)
	GetSiteIds(ctx context.Context, companyUuid uuid.UUID, testType string) ([]int, error)
	GetVulnerabilitySeverityCountsByInterval(ctx context.Context, siteIds []int) ([]*entities.VulnerabilitySeverityCountsByInterval, error)
	GetFactSiteCountsBySiteId(ctx context.Context, siteIds []int) (*entities.FactSiteCounts, error)
	GetTopRemediationTasks(ctx context.Context, siteIds []int) ([]*entities.TopRemediationTask, error)
	GetSolutionSupercedenceSteps(ctx context.Context, solutionId int) ([]string, error)
	GetRemediations(ctx context.Context, siteIds []int, limit int) ([]*entities.Remediation, error)
	FilterSiteIdsByTestType(ctx context.Context, siteIds []int, testType string) ([]int, error)
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) FilterSiteIdsByTestType(ctx context.Context, siteIds []int, testType string) ([]int, error) {
	return l.svc.FilterSiteIdsByTestType(ctx, siteIds, testType)
}

func (l *localClient) GetRemediations(ctx context.Context, siteIds []int, limit int) ([]*entities.Remediation, error) {
	return l.svc.GetRemediations(ctx, siteIds, limit)
}

func (l *localClient) GetSolutionSupercedenceSteps(ctx context.Context, solutionId int) ([]string, error) {
	return l.svc.GetSolutionSupercedenceSteps(ctx, solutionId)
}

func (l *localClient) GetTopRemediationTasks(ctx context.Context, siteIds []int) ([]*entities.TopRemediationTask, error) {
	return l.svc.GetTopRemediationTasks(ctx, siteIds)
}

func (l *localClient) GetFactSiteCountsBySiteId(ctx context.Context, siteIds []int) (*entities.FactSiteCounts, error) {
	return l.svc.GetFactSiteCountsBySiteId(ctx, siteIds)
}

func (l *localClient) GetVulnerabilitySeverityCountsByInterval(ctx context.Context, siteIds []int) ([]*entities.VulnerabilitySeverityCountsByInterval, error) {
	return l.svc.GetVulnerabilitySeverityCountsByInterval(ctx, siteIds)
}

func (l *localClient) GetSiteIds(ctx context.Context, companyUuid uuid.UUID, testType string) ([]int, error) {
	return l.svc.GetSiteIds(ctx, companyUuid, testType)

}

func (l *localClient) GetLastAndNextVulnerabilityScan(ctx context.Context, siteIds []int) (*entities.LastNexVulnerabilityScan, error) {
	return l.svc.GetLastAndNextVulnerabilityScan(ctx, siteIds)
}
