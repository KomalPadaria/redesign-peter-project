// Package swagger
package swagger //nolint: predeclared

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/swagger/transport/http"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// ModuleParams for user.
type ModuleParams struct {
	fx.In

	HTTPServer *httpTransport.Server
}

// NewModule for swagger.
// nolint:gocritic
func NewModule(p ModuleParams) error {

	http.RegisterTransport(p.HTTPServer)

	return nil
}

var (
	// ModuleServeSwagger for uber fx.
	ModuleServeSwagger = fx.Options(fx.Invoke(NewModule))
)
