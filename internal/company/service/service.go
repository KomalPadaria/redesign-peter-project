package service

import (
	"context"
	"io"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/repository"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	appErrors "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	httpErr "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

// Service stores business logic
type Service interface {
	CreateCompany(ctx context.Context, company *entities.UsersCompany) (uuid.UUID, error)
	FindByExternalId(ctx context.Context, req *entities.FindCompanyByExternalIdRequest) (*entities.FindCompanyByExternalIdResponse, error)
	GetUserCompaniesByUserUuid(ctx context.Context, userUuid uuid.UUID, keyword string) (*entities.GetCompaniesByUserIdResponse, error)
	FindByUUID(ctx context.Context, companyUuid *uuid.UUID) (*entities.Company, error)
	UpdateCompany(ctx context.Context, updateData *entities.UpdateCompanyRequest) error
	GetCompanyInfo(ctx context.Context, companyUuid uuid.UUID) (*entities.Company, error)
	GetAllCompanies(ctx context.Context, keyword string) ([]*entities.Company, error)
	UploadSecurityCampaignUsers(ctx context.Context, req *entities.UploadSecurityCampaignUsersRequest) error
	GetSecurityCampaignUsers(ctx context.Context, req *entities.GetCampaignUsersRequest) (interface{}, error)

	CreateCompanySubscription(ctx context.Context, cs *entities.CompanySubscription) error
	UpdateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string, cs *entities.CompanySubscription) error
	DeleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string) error
	GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
	GetConsultingHoursSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
	FetchCompanySubscriptionFromSF(ctx context.Context, companyUuid *uuid.UUID) error
	GetSubscriptionByName(ctx context.Context, companyUuid *uuid.UUID, subscriptionName string) (*entities.CompanySubscription, error)
	GetServiceReviewSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error)
}

// New service for company.
func New(
	repo repository.Repository,
	frameworksClient frameworksClient.Client,
	onboardingClient onboarding.Client,
	sfClient salesforce.Client,
) Service {
	svc := &service{
		repo:             repo,
		frameworksClient: frameworksClient,
		onboardingClient: onboardingClient,
		sfClient:         sfClient,
	}

	return svc
}

type service struct {
	repo             repository.Repository
	frameworksClient frameworksClient.Client
	onboardingClient onboarding.Client
	sfClient         salesforce.Client
}

func (s *service) GetServiceReviewSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	allSubs, err := s.GetAllSubscriptionsByCompany(ctx, companyUuid)
	if err != nil {
		return nil, err
	}

	var res []entities.CompanySubscription
	var productIds []string
	productIds = append(productIds, strings.Split(s.sfClient.GetConfig().ServiceSubscriptions, ",")...)
	for _, sfSub := range allSubs {
		if slices.Contains(productIds, sfSub.SfProductID) {
			res = append(res, sfSub)
		}
	}

	return res, nil
}

func (s *service) GetSubscriptionByName(ctx context.Context, companyUuid *uuid.UUID, subscriptionName string) (*entities.CompanySubscription, error) {
	return s.repo.GetSubscriptionByName(ctx, companyUuid, subscriptionName)
}

func (s *service) GetConsultingHoursSubscriptions(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	allSubs, err := s.GetAllSubscriptionsByCompany(ctx, companyUuid)
	if err != nil {
		return nil, err
	}

	var res []entities.CompanySubscription
	var productIds []string
	productIds = append(productIds, strings.Split(s.sfClient.GetConfig().ConsultingHoursSubscriptions, ",")...)
	for _, sfSub := range allSubs {
		if slices.Contains(productIds, sfSub.SfProductID) {
			res = append(res, sfSub)
		}
	}

	return res, nil
}

func (s *service) FetchCompanySubscriptionFromSF(ctx context.Context, companyUuid *uuid.UUID) error {
	c, err := s.FindByUUID(ctx, companyUuid)
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

			s.CreateCompanySubscription(ctx, &entities.CompanySubscription{
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

func (s *service) GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	return s.repo.GetAllSubscriptionsByCompany(ctx, companyUuid)
}

func (s *service) UpdateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string, cs *entities.CompanySubscription) error {
	return s.repo.UpdateCompanySubscription(ctx, companyUuid, subscriptionId, cs)
}

func (s *service) DeleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string) error {
	cs, err := s.repo.DeleteCompanySubscription(ctx, companyUuid, subscriptionId)
	if err != nil {
		return err
	}
	// create company frameworks link based on subscription
	var frameworkNames []string
	valueParts := strings.Split(s.sfClient.GetConfig().FrameworksSubscriptions, ",")
	for _, v := range valueParts {
		parts := strings.Split(v, "-")
		frameworkProductId, frameworkName := parts[0], parts[1]
		if cs.SfProductID == frameworkProductId {
			frameworkNames = append(frameworkNames, frameworkName)
		}
	}

	err = s.frameworksClient.DeleteCompanyFrameworksLink(ctx, &cs.CompanyUuid, nil, frameworkNames)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateCompanySubscription(ctx context.Context, cs *entities.CompanySubscription) error {
	err := s.repo.CreateCompanySubscription(ctx, cs)
	if err != nil {
		return err
	}

	// create company frameworks link based on subscription
	var frameworkNames []string
	valueParts := strings.Split(s.sfClient.GetConfig().FrameworksSubscriptions, ",")
	for _, v := range valueParts {
		parts := strings.Split(v, "-")
		frameworkProductId, frameworkName := parts[0], parts[1]
		if cs.SfProductID == frameworkProductId {
			frameworkNames = append(frameworkNames, frameworkName)
		}
	}

	err = s.frameworksClient.CreateCompanyFrameworksLink(ctx, &cs.CompanyUuid, nil, frameworkNames)
	if err != nil {
		errCause := errors.Cause(err)
		if ec, ok := errCause.(*pgconn.PgError); ok && ec.Code != pgerrcode.UniqueViolation {
			return err
		}
	}

	return nil
}

func (s *service) CreateCompany(ctx context.Context, company *entities.UsersCompany) (uuid.UUID, error) {
	companyUUID, err := s.repo.CreateCompany(ctx, company)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.FetchCompanySubscriptionFromSF(ctx, &companyUUID)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.repo.CreateCompanyQuestionnaires(ctx, &companyUUID, &company.CreatedBy)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.frameworksClient.CreateFrameworkControlRemediations(ctx, &companyUUID, &company.CreatedBy)
	if err != nil {
		return uuid.Nil, err
	}

	return companyUUID, nil
}

func (s *service) FindByExternalId(ctx context.Context, req *entities.FindCompanyByExternalIdRequest) (*entities.FindCompanyByExternalIdResponse, error) {
	company, err := s.repo.FindByExternalId(ctx, req.ExternalId)
	if err != nil {
		return nil, err
	}
	res := &entities.FindCompanyByExternalIdResponse{Company: company}
	return res, nil
}

func (s *service) GetUserCompaniesByUserUuid(ctx context.Context, userUuid uuid.UUID, keyword string) (*entities.GetCompaniesByUserIdResponse, error) {
	companies, err := s.repo.GetUserCompaniesByUserUuid(ctx, userUuid, keyword)
	if err != nil {
		return nil, errors.Wrap(err, "finding user")
	}

	if companies == nil {
		return nil, &appErrors.ErrNotFound{Message: "Companies for the user not found"}
	}

	return s.getCompaniesByUserIdResponse(companies), nil
}

func (s *service) getCompaniesByUserIdResponse(res []*entities.UsersCompany) *entities.GetCompaniesByUserIdResponse {
	companies := make([]*entities.Company, 0)

	for _, c := range res {
		company := &entities.Company{}
		company.CompanyUuid = c.CompanyUuid
		company.Name = c.Name
		company.Type = c.Type
		company.IndustryType = &c.IndustryType

		onboardingGroup := c.Onboarding
		sort.Slice(onboardingGroup, func(i, j int) bool {
			return onboardingGroup[i].Position < onboardingGroup[j].Position
		})

		company.Onboarding = onboardingGroup
		company.Address = &c.Address
		company.UserRole = c.Role

		companies = append(companies, company)
	}

	return &entities.GetCompaniesByUserIdResponse{Companies: companies}
}

func (s *service) FindByUUID(ctx context.Context, companyUuid *uuid.UUID) (*entities.Company, error) {
	return s.repo.FindByUUID(ctx, companyUuid)
}

func (s *service) UpdateCompany(ctx context.Context, updateData *entities.UpdateCompanyRequest) error {
	return s.repo.UpdateCompany(ctx, updateData)
}

func (s *service) GetCompanyInfo(ctx context.Context, companyUuid uuid.UUID) (*entities.Company, error) {
	return s.repo.FindByUUID(ctx, &companyUuid)
}

func (s *service) GetAllCompanies(ctx context.Context, keyword string) ([]*entities.Company, error) {
	return s.repo.GetAllCompanies(ctx, keyword)
}

func (s *service) UploadSecurityCampaignUsers(ctx context.Context, req *entities.UploadSecurityCampaignUsersRequest) error {
	fileBytes, err := io.ReadAll(req.File)
	defer req.File.Close()
	if err != nil {
		return err
	}

	var su entities.SecurityCampaignUsers
	err = gocsv.UnmarshalBytes(fileBytes, &su)
	if err != nil {
		return httpErr.ErrFileNotSupported
	}

	err = s.repo.UpdateCompany(ctx, &entities.UpdateCompanyRequest{
		CompanyUuid: req.CompanyUuid,
		Data: map[string]interface{}{
			"campaign_users": su,
		},
	})
	if err != nil {
		return err
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.SecurityAwarenessTraining, req.CompanyUuid, req.UserUuid)

	return nil
}

func (s *service) GetSecurityCampaignUsers(ctx context.Context, req *entities.GetCampaignUsersRequest) (interface{}, error) {
	data, err := s.repo.FindByUUID(ctx, &req.CompanyUuid)
	if err != nil {
		return nil, err
	}

	if data.CampaignUsers == nil {
		var sut []entities.SecurityCampaignUserTemplate
		// if there is no data for "campaign_users", we return template CSV data
		sut = append(sut, entities.SecurityCampaignUserTemplate{
			Email:    "test1@kb4-demo.com",
			Name:     "test name 1",
			JobTitle: "test manager 1",
		})
		sut = append(sut, entities.SecurityCampaignUserTemplate{
			Email:    "test2@kb4-demo.com",
			Name:     "test name 2",
			JobTitle: "test manager 2"})
		sut = append(sut, entities.SecurityCampaignUserTemplate{
			Email:    "test3@kb4-demo.com",
			Name:     "test name 3",
			JobTitle: "test manager 3"})
		sut = append(sut, entities.SecurityCampaignUserTemplate{
			Email:    "test4@kb4-demo.com",
			Name:     "test name 4",
			JobTitle: "test manager 4"})

		csvContent, err := gocsv.MarshalBytes(&sut)
		if err != nil {
			return nil, err
		}

		return csvContent, nil
	}

	csvContent, err := gocsv.MarshalBytes(&data.CampaignUsers)
	if err != nil {
		return nil, err
	}

	return csvContent, nil
}
