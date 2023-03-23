package onboarding

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cache"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// ModuleParams contain dependencies for module
type ModuleParams struct {
	fx.In
	DB *gorm.DB
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	// set cache to never expire
	cache := cache.New(0, 0)
	repo := repository.New(p.DB)
	svc := service.New(repo, cache)

	client := newClient(svc)

	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
