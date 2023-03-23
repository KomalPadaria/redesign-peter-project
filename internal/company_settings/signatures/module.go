package signatures //nolint: predeclared

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/userclient"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
	UserClient       userclient.Client
	CompanyClient    company.Client
	OnboardingClient onboarding.Client
	SalesforceClient salesforce.Client
	CommonConfig     cfg.Config
	Logger           *zap.SugaredLogger
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.DB, p.GormDB)
	svc := service.New(repo, p.UserClient, p.CompanyClient, p.OnboardingClient, p.SalesforceClient, p.CommonConfig, p.Logger)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
