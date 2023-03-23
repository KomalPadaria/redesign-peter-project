package frameworks

import (
	"database/sql"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/service"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ModuleParams for user.
type ModuleParams struct {
	fx.In

	DB               *gorm.DB
	SqlDB            *sql.DB
	HTTPServer       *httpTransport.Server
	APPTransport     svcTransport.Client
	AuthClient       auth.Client
	SalesforceClient salesforce.Client
	FrameworksClient frameworksClient.Client
	Logger           *zap.SugaredLogger
}

// NewModule for redesign.
// nolint:gocritic
func NewModule(p ModuleParams) error {
	repo := repository.New(p.DB, p.SqlDB)
	svc := service.New(repo, p.SalesforceClient, p.Logger)
	eps := endpoints.New(svc)

	http.RegisterTransport(p.HTTPServer, eps, p.AuthClient, p.APPTransport)

	return nil
}

var (
	// ModuleHttpAPI for uber fx.
	ModuleHttpAPI = fx.Options(fx.Invoke(NewModule))
)
