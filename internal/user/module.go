// Package user
package user //nolint: predeclared

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for user.
type ModuleParams struct {
	fx.In

	DB               *sql.DB
	HTTPServer       *httpTransport.Server
	CommonConfig     cfg.Config
	APPTransport     svcTransport.Client
	CompanyClient    company.Client
	SalesforceClient salesforce.Client
	AuthClient       auth.Client
	CognitoClient    cognito.Client
	OnboardingClient onboarding.Client
	EmailClient      ses.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.DB)
	svc := service.New(repo, p.CompanyClient, p.SalesforceClient, p.CognitoClient, p.OnboardingClient, p.EmailClient, p.CommonConfig)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
