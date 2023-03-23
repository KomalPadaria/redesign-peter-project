package repository

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	GetLastAndNextVulnerabilityScan(ctx context.Context, siteIds []int) (*entities.LastNexVulnerabilityScan, error)
	GetVulnerabilitySeverityCountsByInterval(ctx context.Context, siteIds []int) ([]*entities.VulnerabilitySeverityCountsByInterval, error)
	GetFactSiteCountsBySiteId(ctx context.Context, siteIds []int) (*entities.FactSiteCounts, error)
	GetTopRemediationTasks(ctx context.Context, siteIds []int) ([]*entities.TopRemediationTask, error)
	GetSolutionSupercedenceSteps(ctx context.Context, solutionId int) ([]string, error)
	GetRemediations(ctx context.Context, siteIds []int, limit int) ([]*entities.Remediation, error)
	FilterSiteIdsByTestType(ctx context.Context, siteIds []int, testType string) ([]int, error)
}

// New repository for websites.
func New(gormDB *gorm.DB) Repository {
	repo := &sqlRepository{gormDB}

	return repo
}
