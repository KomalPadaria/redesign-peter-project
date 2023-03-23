// Package string contains example of the service package
package stringsvc //nolint: predeclared

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/transport/grpc"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/transport/http"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for stringsvc.
type ModuleParams struct {
	fx.In

	HTTPServer   *httpTransport.Server
	APPTransport svcTransport.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	svc := service.New()
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.APPTransport)

	client := NewClient(eps)

	return client, nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))

	ModuleGrpcAPI = fx.Options(
		fx.Provide(
			service.New,
			endpoints.New,
		),
		fx.Invoke(
			grpc.Register,
		),
	)
)
