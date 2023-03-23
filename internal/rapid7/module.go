package rapid7

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"go.uber.org/fx"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In

	Config        config.Config
	CompanyClient company.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	_, db, err := db.New(&p.Config.DataWarehouse)
	if err != nil {
		return nil, err
	}
	repo := repository.New(db)
	svc := service.New(p.Config, repo, p.CompanyClient)

	client := NewClient(svc)

	return client, nil
}

var (
	// ModuleClient for uber fx.
	ModuleClient = fx.Options(fx.Provide(NewModule))
)
