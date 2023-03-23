package service

import (
	"context"
	"strings"

	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"golang.org/x/exp/slices"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"go.uber.org/zap"
)

type Service interface {
	GetFrameworks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.Framework, error)
	GetFrameworkControls(ctx context.Context, companyUuid, userUuid, frameworkUuid *uuid.UUID) ([]*entities.FrameworkControl, error)
	CreateFrameworkControlRemediations(ctx context.Context, companyUUID, userUUID *uuid.UUID) error
	GetFrameworkStats(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetFrameworkStatsResponse, error)
	CreateCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error
	DeleteCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error
}

func New(repo repository.Repository, sfClient salesforce.Client, logger *zap.SugaredLogger) Service {
	return &service{
		repo:     repo,
		sfClient: sfClient,
		logger:   logger,
	}
}

type service struct {
	repo     repository.Repository
	sfClient salesforce.Client
	logger   *zap.SugaredLogger
}

func (s *service) DeleteCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error {
	return s.repo.DeleteCompanyFrameworksLink(ctx, companyUUID, userUUID, frameworkNames)
}

func (s *service) CreateCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error {
	if len(frameworkNames) == 0 {
		return nil
	}
	return s.repo.CreateCompanyFrameworksLink(ctx, companyUUID, userUUID, frameworkNames)
}

func (s *service) GetFrameworks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.Framework, error) {
	return s.repo.GetFrameworks(ctx, companyUuid, userUuid)
}

func (s *service) GetFrameworkControls(ctx context.Context, companyUuid, userUuid, frameworkUuid *uuid.UUID) ([]*entities.FrameworkControl, error) {
	return s.repo.GetFrameworkControls(ctx, companyUuid, userUuid, frameworkUuid)
}

func (s *service) GetFrameworkStats(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetFrameworkStatsResponse, error) {
	//TODO this section need to be reemoved on salesforce webhook is implemented.
	// Because the company framework link will be created via webhook
	frameworks, err := s.GetFrameworks(ctx, companyUuid, userUuid)
	if err != nil {
		s.logger.Warn("error getting  frameworks", err)
	} else {
		if len(frameworks) == 0 {
			err = s.fetchCompanySubscriptionFromSF(ctx, companyUuid)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.repo.GetFrameworkStats(ctx, companyUuid, userUuid)
}

func (s *service) CreateFrameworkControlRemediations(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	return s.repo.CreateFrameworkControlRemediations(ctx, companyUUID, userUUID)
}

func (s *service) fetchCompanySubscriptionFromSF(ctx context.Context, companyUuid *uuid.UUID) error {
	c, err := s.repo.GetCompanyByUUID(ctx, companyUuid)
	if err != nil {
		return err
	}

	sfSess, err := s.sfClient.NewSession(ctx)
	if err != nil {
		return err
	}

	sfSubs, err := s.sfClient.GetSubscriptionsByAccountID(ctx, sfSess, c.ExternalId)
	if err != nil {
		return err
	}

	var productIds []string

	productIds = append(productIds, strings.Split(s.sfClient.GetConfig().MainSubscriptions, ",")...)
	productIds = append(productIds, strings.Split(s.sfClient.GetConfig().ConsultingHoursSubscriptions, ",")...)
	productIds = append(productIds, strings.Split(s.sfClient.GetConfig().ServiceSubscriptions, ",")...)

	valueParts := strings.Split(s.sfClient.GetConfig().FrameworksSubscriptions, ",")
	for _, v := range valueParts {
		parts := strings.Split(v, "-")
		frameworkProductId := parts[0]
		productIds = append(productIds, frameworkProductId)
	}

	for _, sfSub := range sfSubs {
		if slices.Contains(productIds, sfSub.SBQQProductC) {
			startD, endD, err := salesforce.ParseStartDateEndDate(sfSub.SBQQStartDateC, sfSub.SBQQEndDateC)
			if err != nil {
				return err
			}

			s.repo.CreateCompanySubscription(ctx, &companyEntities.CompanySubscription{
				CompanyUuid:      *companyUuid,
				Name:             sfSub.SBQQProductNameC,
				SfSubscriptionID: sfSub.ID,
				SfProductID:      sfSub.SBQQProductC,
				Status:           "active",
				StartDate:        startD,
				EndDate:          endD,
			})
		}
	}

	return nil
}
