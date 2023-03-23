package s3client

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/service"
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

	client := NewClient(svc)

	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
