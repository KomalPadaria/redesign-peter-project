package ses

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses/service"
	"go.uber.org/fx"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In

	Config config.Config
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) Client {
	svc := service.New(p.Config)

	client := NewClient(svc)

	return client
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
