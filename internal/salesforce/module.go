package salesforce

import (
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/service"
	"go.uber.org/fx"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In

	Config     config.Config
	HttpClient *http.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	svc := service.New(p.Config, p.HttpClient)

	client := NewClient(svc)

	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
