package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	jiraEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/entities"
	sfWebhookError "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	SubscriptionStatusCreated = "Created"
	SubscriptionStatusUpdated = "Updated"
	SubscriptionStatusDeleted = "Deleted"
)

type Service interface {
	SFAccountWebhook(ctx context.Context, req *entities.AccountWebhookData) error
	SFAccountSubscriptionsWebhook(ctx context.Context, req entities.AccountSubscriptionWebhookData) error
}

func New(logger *zap.SugaredLogger, companyClient company.Client, jiraClient jira.Client, sfClient salesforce.Client) Service {
	return &service{logger, companyClient, jiraClient, sfClient}
}

type service struct {
	logger        *zap.SugaredLogger
	companyClient company.Client
	jiraClient    jira.Client
	sfClient      salesforce.Client
}

func (s service) SFAccountSubscriptionsWebhook(ctx context.Context, req entities.AccountSubscriptionWebhookData) error {
	s.logger.Info("Salesforce account subscription  webhook data : ", req)

	sfSess, err := s.sfClient.NewSession(ctx)
	if err != nil {
		return err
	}
	for _, data := range req {
		companyRes, err := s.companyClient.FindByExternalId(ctx, &companyEntities.FindCompanyByExternalIdRequest{
			ExternalId: data.AccountId,
		})
		if err != nil {
			return err
		}

		if companyRes.Company != nil {
			company := companyRes.Company

			subscription, err := s.sfClient.GetSubscriptionsByID(ctx, sfSess, data.SubscriptionId)
			if err != nil {
				return err
			}

			startD, endD, err := salesforce.ParseStartDateEndDate(subscription.SBQQStartDateC, subscription.SBQQEndDateC)
			if err != nil {
				s.logger.Error(errors.Wrap(sfWebhookError.ErrSFWebhook, fmt.Sprintf("%v", err)))
				return nil
			}

			switch data.Status {
			case SubscriptionStatusCreated:
				err = s.createCompanySubscription(ctx, company.CompanyUuid, subscription.SBQQProductNameC, subscription.ID, subscription.SBQQProductC, startD, endD)
				if err != nil {
					s.logger.Error(errors.Wrap(sfWebhookError.ErrSFWebhook, fmt.Sprintf("create company subscription %v", err)))
					return nil
				}
			case SubscriptionStatusUpdated:
				err = s.updateCompanySubscription(ctx, &company.CompanyUuid, subscription.SBQQProductNameC, subscription.ID, subscription.SBQQProductC, startD, endD)
				if err != nil {
					s.logger.Error(errors.Wrap(sfWebhookError.ErrSFWebhook, fmt.Sprintf("update company subscription %v", err)))
					return nil
				}
			case SubscriptionStatusDeleted:
				err = s.deleteCompanySubscription(ctx, &company.CompanyUuid, subscription.ID)
				if err != nil {
					s.logger.Error(errors.Wrap(sfWebhookError.ErrSFWebhook, fmt.Sprintf("delete company subscription %v", err)))
					return nil
				}
			}
		}
	}

	return nil
}

func (s service) SFAccountWebhook(ctx context.Context, req *entities.AccountWebhookData) error {
	s.logger.Info("Salesforce account webhook data : ", req)

	res, err := s.companyClient.FindByExternalId(ctx, &companyEntities.FindCompanyByExternalIdRequest{
		ExternalId: req.Id,
	})
	if err != nil {
		s.logger.Error(errors.Wrap(sfWebhookError.ErrSFWebhook, fmt.Sprintf("finding  sf data %v", err)))
	}

	if res.Company == nil {
		// Create EPIC  in Jira
		createEpicRes, err := s.jiraClient.CreateEpic(ctx, &jiraEntities.CreateIssueRequest{
			Name:    req.Name,
			Summary: req.Name,
		})
		if err != nil {
			s.logger.Error(errors.Wrap(sfWebhookError.ErrSFWebhook, fmt.Sprintf("creating jira epic %v", err)))
		}

		jiraEpicId := createEpicRes.Key

		sfSess, err := s.sfClient.NewSession(ctx)
		if err != nil {
			return err
		}

		_, err = s.sfClient.UpdateJiraEpicIdForAccount(ctx, sfSess, req.Id, jiraEpicId)
		if err != nil {
			return err
		}

		return nil
	}

	company := res.Company
	var rapid7Ids []string
	for _, id := range strings.Split(req.Rapid7, ",") {
		rapid7Ids = append(rapid7Ids, strings.TrimSpace(id))
	}

	updateCompanyReq := companyEntities.UpdateCompanyRequest{
		CompanyUuid: company.CompanyUuid,
		Data: map[string]interface{}{
			"knowbe4_token":   strings.TrimSpace(req.KnowBe4),
			"rapid7_site_ids": rapid7Ids,
			"jira_epic_id":    strings.TrimSpace(req.Jira),
			"updated_at":      nullable.NewNullTime(time.Now().UTC()),
		},
	}

	err = s.companyClient.UpdateCompany(ctx, &updateCompanyReq)
	if err != nil {
		return nil
	}

	return nil
}

func (s service) createCompanySubscription(ctx context.Context, companyUuid uuid.UUID, name, sfSubscriptionID, sfProductID string, startDate, endDate time.Time) error {
	return s.companyClient.CreateCompanySubscription(ctx, &companyEntities.CompanySubscription{
		CompanyUuid:      companyUuid,
		Name:             name,
		SfSubscriptionID: sfSubscriptionID,
		SfProductID:      sfProductID,
		StartDate:        startDate,
		EndDate:          endDate,
	})
}

func (s service) updateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, name, sfSubscriptionID, sfProductID string, startDate, endDate time.Time) error {
	return s.companyClient.UpdateCompanySubscription(ctx, companyUuid, sfSubscriptionID, &companyEntities.CompanySubscription{
		Name:        name,
		SfProductID: sfProductID,
		StartDate:   startDate,
		EndDate:     endDate,
	})
}

func (s service) deleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, sfSubscriptionID string) error {
	return s.companyClient.DeleteCompanySubscription(ctx, companyUuid, sfSubscriptionID)
}
