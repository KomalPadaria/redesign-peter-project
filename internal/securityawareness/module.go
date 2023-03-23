package securityawareness //nolint: predeclared

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness/transport/http"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for applications.
type ModuleParams struct {
	fx.In

	HTTPServer    *httpTransport.Server
	APPTransport  svcTransport.Client
	AuthClient    auth.Client
	KnowBe4Client knowbe4.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	svc := service.New(p.KnowBe4Client)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
