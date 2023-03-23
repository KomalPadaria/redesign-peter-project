package repository

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/entities"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
}

func (s *sqlRepository) FilterSiteIdsByTestType(ctx context.Context, siteIds []int, testType string) ([]int, error) {
	testType = testType + "%"

	var ids []int
	err := s.gormDB.WithContext(ctx).Table("dim_site").
		Where("site_id in ?", siteIds).
		Where("name ilike ?", testType).
		Pluck("site_id", &ids).Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *sqlRepository) GetRemediations(ctx context.Context, siteIds []int, limit int) ([]*entities.Remediation, error) {
	rows, err := s.gormDB.WithContext(ctx).Raw(listRemeditation, siteIds, limit).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]*entities.Remediation, 0)

	for rows.Next() {
		res := &entities.Remediation{}
		err = rows.Scan(
			&res.RecommendationUrl,
			&res.IssueDescription,
			&res.Source,
			&res.VulnerabilityId,
			&res.SolutionId,
			&res.RiskScore,
			&res.Severity,
			&res.Instances,
			&res.IssueName,
			&res.Recommendation,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, res)
	}

	return response, nil
}

func (s *sqlRepository) GetSolutionSupercedenceSteps(ctx context.Context, solutionId int) ([]string, error) {
	var steps []string
	err := s.gormDB.WithContext(ctx).Raw(getSolutionSupercedenceStepsSql, solutionId).Scan(&steps).Error
	if err != nil {
		return nil, err
	}
	return steps, nil
}

func (s *sqlRepository) GetTopRemediationTasks(ctx context.Context, siteIds []int) ([]*entities.TopRemediationTask, error) {
	rows, err := s.gormDB.WithContext(ctx).Raw(topRemediationSql, siteIds).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]*entities.TopRemediationTask, 0)

	for rows.Next() {
		res := &entities.TopRemediationTask{}
		err = rows.Scan(
			&res.SolutionId,
			&res.VulnerabilityId,
			&res.RiskScore,
			&res.Title,
			&res.IssueDescription,
			&res.RemediationName,
			&res.RemediationLink,
			&res.SolutionStep,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, res)
	}

	return response, nil
}

func (s *sqlRepository) GetFactSiteCountsBySiteId(ctx context.Context, siteIds []int) (*entities.FactSiteCounts, error) {
	var counts *entities.FactSiteCounts
	err := s.gormDB.WithContext(ctx).Table("fact_site").
		Where("site_id in (?)", siteIds).
		Select("sum(vulnerabilities) as vulnerabilities",
			"sum(critical_vulnerabilities) as critical_vulnerabilities",
			"sum(severe_vulnerabilities) as severe_vulnerabilities",
			"sum(moderate_vulnerabilities) as moderate_vulnerabilities").
		Find(&counts).Error
	if err != nil {
		return nil, err
	}

	return counts, nil
}

func (s *sqlRepository) GetVulnerabilitySeverityCountsByInterval(ctx context.Context, siteIds []int) ([]*entities.VulnerabilitySeverityCountsByInterval, error) {
	rows, err := s.gormDB.WithContext(ctx).Raw(sql, siteIds, siteIds, siteIds, siteIds, siteIds, siteIds, siteIds, siteIds, siteIds, siteIds, siteIds).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]*entities.VulnerabilitySeverityCountsByInterval, 0)

	for rows.Next() {
		res := &entities.VulnerabilitySeverityCountsByInterval{}
		err = rows.Scan(
			&res.Date,
			&res.Severity,
			&res.Count,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, res)
	}

	return response, nil
}

func (s *sqlRepository) GetLastAndNextVulnerabilityScan(ctx context.Context, siteIds []int) (*entities.LastNexVulnerabilityScan, error) {
	var ln *entities.LastNexVulnerabilityScan
	err := s.gormDB.WithContext(ctx).Raw(lastAndNextVulnerabilityScan, siteIds, siteIds).Scan(&ln).Error
	if err != nil {
		return nil, err
	}

	return ln, nil
}
