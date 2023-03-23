package client

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ModuleClientParams for user.
type ModuleClientParams struct {
	fx.In

	DB               *gorm.DB
	SqlDB            *sql.DB
	SalesforceClient salesforce.Client
	Logger           *zap.SugaredLogger
}

// NewClientModule for user.
// nolint:gocritic
func NewClientModule(p ModuleClientParams) Client {
	repo := repository.New(p.DB, p.SqlDB)
	svc := service.New(repo, p.SalesforceClient, p.Logger)

	client := NewClient(svc)

	return client
}

var (
	// ModuleClient for uber fx.
	ModuleClient = fx.Options(fx.Provide(NewClientModule))
)
