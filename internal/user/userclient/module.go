package userclient

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	"go.uber.org/fx"
)

// ModuleClientParams for user.
type ModuleClientParams struct {
	fx.In

	DB               *sql.DB
	CommonConfig     cfg.Config
	CompanyClient    company.Client
	SalesforceClient salesforce.Client
	CognitoClient    cognito.Client
	OnboardingClient onboarding.Client
	EmailClient      ses.Client
}

// NewClientModule for user.
// nolint:gocritic
func NewClientModule(p ModuleClientParams) Client {
	repo := repository.New(p.DB)
	svc := service.New(repo, p.CompanyClient, p.SalesforceClient, p.CognitoClient, p.OnboardingClient, p.EmailClient, p.CommonConfig)
	eps := endpoints.New(svc)

	client := NewClient(eps, svc)

	return client
}

var (
	// ModuleClient for uber fx.
	ModuleClient = fx.Options(fx.Provide(NewClientModule))
)
