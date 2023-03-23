// Package string contains example of the service package
package company //nolint: predeclared

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/transport/http"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// ModuleParams for company.
type ModuleParams struct {
	fx.In

	DB               *sql.DB
	GormDB           *gorm.DB
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	FrameworksClient frameworksClient.Client
	OnboardingClient onboarding.Client
	SalesforceClient salesforce.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) (Client, error) {
	repo := repository.New(p.DB, p.GormDB)
	svc := service.New(repo, p.FrameworksClient, p.OnboardingClient, p.SalesforceClient)
	eps := endpoints.New(svc)

	client := NewClient(eps, svc)

	http.RegisterTransport(p.HTTPServer, eps, p.APPTransport)
	return client, nil
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
