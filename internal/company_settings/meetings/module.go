package meetings //nolint: predeclared

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// ModuleParams for applications.
type ModuleParams struct {
	fx.In

	DB               *sql.DB
	GormDB           *gorm.DB
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	AuthClient       auth.Client
	CalendlyClient   calendly.Client
	OnboardingClient onboarding.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.DB, p.GormDB)
	svc := service.New(repo, p.CalendlyClient, p.OnboardingClient)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
