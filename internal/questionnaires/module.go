package questionnaires

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	s3client "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ModuleParams for user.
type ModuleParams struct {
	fx.In

	DB               *gorm.DB
	SqlDB            *sql.DB
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	AuthClient       auth.Client
	FrameworksClient frameworksClient.Client
	SalesforceClient salesforce.Client
	JiraClient       jira.Client
	S3client         s3client.Client
	CompanyClient    company.Client
	Logger           *zap.SugaredLogger
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.DB, p.SqlDB)
	svc := service.New(repo, p.Logger, p.SalesforceClient, p.FrameworksClient, p.S3client, p.JiraClient, p.CompanyClient)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
