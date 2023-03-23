package ipranges //nolint: predeclared

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// ModuleParams for applications.
type ModuleParams struct {
	fx.In

	DB               *gorm.DB
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	AuthClient       auth.Client
	OnboardingClient onboarding.Client
	JiraClient       jira.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	repo := repository.New(p.DB)
	svc := service.New(repo, p.OnboardingClient, p.JiraClient)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	client := NewClient(svc)

	return client, nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Provide(NewModule))
)
