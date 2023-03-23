package customer_success

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	s3client "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/config"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// ModuleParams for user.
type ModuleParams struct {
	fx.In

	GormDB           *gorm.DB
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	AuthClient       auth.Client
	SalesforceClient salesforce.Client
	CompanyClient    company.Client
	SFConfig         config.Config
	S3client         s3client.Client
	JiraClient       jira.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.GormDB)
	svc := service.New(repo, p.JiraClient, p.SalesforceClient, p.CompanyClient, p.S3client, p.SFConfig)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
