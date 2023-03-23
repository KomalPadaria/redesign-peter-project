package cognito

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/service"
	"go.uber.org/fx"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In

	Config config.Config
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	svc := service.New(p.Config)

	client := newClient(svc)

	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
