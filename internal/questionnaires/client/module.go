package client

import (
	"database/sql"

	s3client "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ModuleClientParams for user.
type ModuleClientParams struct {
	fx.In

	DB               *gorm.DB
	SqlDB            *sql.DB
	FrameworksClient frameworksClient.Client
	SalesforceClient salesforce.Client
	Logger           *zap.SugaredLogger
	JiraClient       jira.Client
	S3client         s3client.Client
	CompanyClient    company.Client
}

// NewClientModule for user.
// nolint:gocritic
func NewClientModule(p ModuleClientParams) Client {
	repo := repository.New(p.DB, p.SqlDB)
	svc := service.New(repo, p.Logger, p.SalesforceClient, p.FrameworksClient, p.S3client, p.JiraClient, p.CompanyClient)

	client := NewClient(svc)

	return client
}

var (
	// ModuleClient for uber fx.
	ModuleClient = fx.Options(fx.Provide(NewClientModule))
)
