package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"

	"github.com/eko/gocache/lib/v4/store"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cache"
	"github.com/pkg/errors"
)

const (
	baseURL = "https://us.api.knowbe4.com/v1"
)

type Service interface {
	GetAllPhishingSecurityTests(ctx context.Context, companyUuid uuid.UUID) ([]entities.SecurityTest, error)
	GetAllRecipientResults(ctx context.Context, companyUuid uuid.UUID, pstId int) ([]entities.RecipientResult, error)
	GetAllTrainingEnrollments(ctx context.Context, companyUuid uuid.UUID) ([]entities.TrainingEnrollment, error)
}

func New(httpClient *http.Client, cacheClient cache.Client, companyClient company.Client) Service {
	hc := client.New(baseURL, httpClient, client.WithExternalCall(true))
	return &service{hc, cacheClient, companyClient}
}

type service struct {
	httpClient    client.Client
	cacheClient   cache.Client
	companyClient company.Client
}

func (s *service) GetAllTrainingEnrollments(ctx context.Context, companyUuid uuid.UUID) ([]entities.TrainingEnrollment, error) {
	var enrollments []entities.TrainingEnrollment

	err := s.getData(ctx, "training/enrollments", &enrollments, companyUuid)
	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

func (s *service) GetAllRecipientResults(ctx context.Context, companyUuid uuid.UUID, pstId int) ([]entities.RecipientResult, error) {
	var recipientResults []entities.RecipientResult

	err := s.getData(ctx, fmt.Sprintf("phishing/security_tests/%d/recipients", pstId), &recipientResults, companyUuid)
	if err != nil {
		return nil, err
	}

	return recipientResults, nil
}

func (s *service) GetAllPhishingSecurityTests(ctx context.Context, companyUuid uuid.UUID) ([]entities.SecurityTest, error) {
	var securityTests []entities.SecurityTest

	err := s.getData(ctx, "phishing/security_tests", &securityTests, companyUuid)
	if err != nil {
		return nil, err
	}

	return securityTests, nil
}

func (s *service) knowBe4Headers(token string) map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"Accept":        "application/json",
	}
}
func (s *service) getData(ctx context.Context, path string, out any, companyUuid uuid.UUID) error {
	var accessToken string
	company, err := s.companyClient.FindByUUID(ctx, &companyEntities.GetCompanyByIdRequest{CompanyUuid: companyUuid})
	if err != nil {
		return err
	}
	accessToken = company.Knowbe4Token
	if accessToken == "" {
		return errors.New(fmt.Sprintf("knkowbe4 access token not found for the company %s", companyUuid.String()))
	}

	return s.getWithCache(ctx, path, out, accessToken)
}

func (s *service) getWithCache(ctx context.Context, path string, out any, accessToken string) error {
	catchKey := path + ":" + accessToken
	val, err := s.cacheClient.Get(ctx, catchKey)
	if err != nil {
		if err.Error() != store.NOT_FOUND_ERR {
			return errors.WithMessage(err, "cache error")
		}
	}

	if val != "" {
		return json.Unmarshal([]byte(val), out)
	}

	err = s.httpClient.Get(ctx, path, s.knowBe4Headers(accessToken), nil, out)
	if err != nil {
		return s.handleError(err)
	}

	b, err := json.Marshal(out)
	if err != nil {
		return err
	}

	err = s.cacheClient.Set(ctx, catchKey, string(b))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) handleError(err error) error {
	e := client.InvalidResponseError(err)
	if e != nil {
		errCode := e.HTTPStatusCode
		errMsg := e.Description

		var knowBe4Err entities.KnowBe4Error
		mErr := json.Unmarshal([]byte(errMsg), &knowBe4Err)
		if mErr != nil {
			return mErr
		}
		return &client.ErrInvalidResponse{HTTPStatusCode: errCode, Code: errCode, Description: "knowbe4: " + knowBe4Err.Message}
	}
	return e
}
