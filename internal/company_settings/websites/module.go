package websites //nolint: predeclared

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/transport/http"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for applications.
type ModuleParams struct {
	fx.In

	DB           *sql.DB
	HTTPServer   *httpTransport.Server
	APPTransport svcTransport.Client
	AuthClient   auth.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.DB)
	svc := service.New(repo)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
