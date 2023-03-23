package penetrationtesting

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/transport/http"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/vulnerability"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for user.
type ModuleParams struct {
	fx.In

	DB                  *sql.DB
	HTTPServer          *httpTransport.Server
	APPTransport        svcTransport.Client
	AuthClient          auth.Client
	VulnerabilityClient vulnerability.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	svc := service.New(p.VulnerabilityClient)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
