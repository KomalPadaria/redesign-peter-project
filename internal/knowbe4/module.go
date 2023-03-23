package knowbe4

import (
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cache"
	"go.uber.org/fx"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In

	CacheClient   cache.Client
	HttpClient    *http.Client
	CompanyClient company.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	svc := service.New(p.HttpClient, p.CacheClient, p.CompanyClient)

	client := NewClient(svc)

	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
