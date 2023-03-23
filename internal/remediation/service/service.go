package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/eko/gocache/lib/v4/store"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/remediation/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cache"
	"github.com/pkg/errors"
)

type Service interface {
	ListRemediation(ctx context.Context, req *entities.ListTopRemediationRequest) ([]*entities.Remediation, error)
}

type service struct {
	rapid7Client rapid7.Client
	cacheClient  cache.Client
}

// New service for user.
func New(rapid7Client rapid7.Client, cacheClient cache.Client) Service {
	svc := &service{rapid7Client, cacheClient}

	return svc
}

func (s *service) ListRemediation(ctx context.Context, req *entities.ListTopRemediationRequest) ([]*entities.Remediation, error) {
	//siteName := s.rapid7Client.GetDefaultSiteName(ctx)
	//get site_id by site name
	var limit int
	siteIds, err := s.rapid7Client.GetSiteIds(ctx, req.CompanyUuid, "")
	if err != nil {
		return nil, err
	}

	catchPath := fmt.Sprintf("%s-%s", "ListRemediation", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(siteIds)), ","), "[]"))
	log.Printf("ListRemediation Catch path: %+v", catchPath)

	var res []*entities.Remediation

	err = s.getWithCache(ctx, catchPath, &res)
	if err != nil {
		return nil, err
	}

	if res != nil {
		return res, nil
	}

	if req.Top != 0 {
		limit = req.Top
	} else {
		limit = 100
	}

	rs, err := s.rapid7Client.GetRemediations(ctx, siteIds, limit)
	if err != nil {
		return nil, err
	}

	res = make([]*entities.Remediation, 0)

	for _, r := range rs {
		rd := entities.RemediationDetail{}
		rd.RemediationName = r.IssueName.String
		rd.RemediationLink = r.RecommendationUrl.String
		rd.IssueDescription = r.IssueDescription.String
		t := &entities.Remediation{
			Severity:       r.Severity.String,
			Instances:      r.Instances,
			Source:         r.Source.String,
			IssueName:      r.IssueName.String,
			Recommendation: r.Recommendation.String,
		}

		var steps []string
		solSteps, err := s.rapid7Client.GetSolutionSupercedenceSteps(ctx, r.SolutionId)
		if err != nil {
			return nil, err
		}
		steps = append(steps, r.Recommendation.String)
		steps = append(steps, solSteps...)

		rd.Recommendations = steps
		t.RemediationDetail = rd
		res = append(res, t)
	}

	resStr, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = s.cacheClient.Set(ctx, catchPath, string(resStr))
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *service) getWithCache(ctx context.Context, path string, out any) error {
	val, err := s.cacheClient.Get(ctx, path)
	if err != nil {
		if err.Error() != store.NOT_FOUND_ERR {
			return errors.WithMessage(err, "cache error")
		}
	}
	if val != "" {
		return json.Unmarshal([]byte(val), out)
	}
	return nil
}
