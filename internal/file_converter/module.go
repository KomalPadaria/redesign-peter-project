package policies //nolint: predeclared

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/transport/http"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/converter"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for applications.
type ModuleParams struct {
	fx.In

	File         *converter.File
	HTTPServer   *httpTransport.Server
	APPTransport svcTransport.Client
	AuthClient   auth.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	// repo := file_converter.New(p.File)
	svc := service.New(*p.File)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
