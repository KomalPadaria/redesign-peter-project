package webhooks //nolint: predeclared

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/transport/http"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModuleParams for applications.
type ModuleParams struct {
	fx.In

	Logger           *zap.SugaredLogger
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	AuthClient       auth.Client
	CompanyClient    company.Client
	JiraClient       jira.Client
	SalesforceClient salesforce.Client
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	svc := service.New(p.Logger, p.CompanyClient, p.JiraClient, p.SalesforceClient)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
