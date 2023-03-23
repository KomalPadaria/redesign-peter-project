package jira

import (
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In

	Config        config.Config
	HttpClient    *http.Client
	CompanyClient company.Client
	Logger        *zap.SugaredLogger
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	svc, err := service.New(p.HttpClient, p.Config, p.CompanyClient, p.Logger)
	if err != nil {
		return nil, err
	}

	client := NewClient(svc)

	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
