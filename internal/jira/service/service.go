package service

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/client"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	baseURL = "https://redesigngroup.atlassian.net"
)

type Service interface {
	CreateEpic(ctx context.Context, req *entities.CreateIssueRequest) (*entities.CreateIssueResponse, error)
	AddComment(ctx context.Context, companyUuid *uuid.UUID, comment *entities.Comment) error
	GetIssuesByEpicId(ctx context.Context, epicId string) ([]entities.Issues, error)
}

func New(httpClient *http.Client, config config.Config, companyClient company.Client, logger *zap.SugaredLogger) (Service, error) {
	hc := client.New(baseURL, httpClient, client.WithExternalCall(true))

	return &service{hc, config, companyClient, logger}, nil
}

type service struct {
	httpClient    client.Client
	config        config.Config
	companyClient company.Client
	logger        *zap.SugaredLogger
}

func (s *service) CreateEpic(ctx context.Context, req *entities.CreateIssueRequest) (*entities.CreateIssueResponse, error) {
	apiEndpoint := "/rest/api/3/issue"

	issue := map[string]interface{}{
		"fields": map[string]interface{}{
			"summary":           req.Summary,
			"customfield_10003": req.Name,
			"project": map[string]interface{}{
				"key": s.config.ProjectKey,
			},
			"issuetype": map[string]interface{}{
				"name": "Epic",
			},
		},
	}

	var response *entities.CreateIssueResponse
	err := s.httpClient.Post(ctx, apiEndpoint, s.jiraHeaders(), issue, &response)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "JIRA Error"))
		return nil, errors.Wrap(err, "JIRA Error")
	}

	return response, nil
}

func (s *service) GetIssuesByEpicId(ctx context.Context, epicId string) ([]entities.Issues, error) {
	// cf[10006] is jira custom field identifier for 'Epic Link'
	apiEndpoint := fmt.Sprintf("rest/api/3/search?jql=cf[10006]=%s&fields=summary,assignee,worklog,status&expand=changelog", epicId)

	var sr entities.SearchResult
	err := s.httpClient.Get(ctx, apiEndpoint, s.jiraHeaders(), nil, &sr)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "JIRA Error"))
		return nil, errors.Wrap(err, "JIRA Error")
	}

	return sr.Issues, nil
}

func (s *service) AddComment(ctx context.Context, companyUuid *uuid.UUID, comment *entities.Comment) error {
	company, err := s.getCompanyByUuid(ctx, companyUuid)
	if err != nil {
		return err
	}

	if company.JiraEpicId == "" {
		s.logger.Warn(fmt.Sprintf("JIRA Epic ID not found for the company '%s'", company.Name))
		return nil
	}

	apiEndpoint := fmt.Sprintf("/rest/api/3/issue/%s/comment", company.JiraEpicId)

	err = s.httpClient.Post(ctx, apiEndpoint, s.jiraHeaders(), comment, nil)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "JIRA Error"))
		return errors.Wrap(err, "JIRA Error")
	}

	return nil
}

func (s *service) jiraHeaders() map[string]string {
	data := fmt.Sprintf("%s:%s", s.config.Username, s.config.ApiToken)
	token := b64.StdEncoding.EncodeToString([]byte(data))

	return map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", token),
	}
}

func (s *service) getCompanyByUuid(ctx context.Context, companyUuid *uuid.UUID) (*companyEntities.Company, error) {
	return s.companyClient.FindByUUID(ctx, &companyEntities.GetCompanyByIdRequest{
		CompanyUuid: *companyUuid,
	})
}
