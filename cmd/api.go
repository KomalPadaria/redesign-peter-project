// Package cmd contains commands
package cmd

import (
	"database/sql"
	"time"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito"
	s3client "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success"
	file_converter "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	penetrationtesting "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires"
	questionnairesClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/remediation"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/swagger"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/userclient"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/vulnerability"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cache"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/converter"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/health"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/health/check"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/log"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/module"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/grpc"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var apiCommand = &cobra.Command{
	Use:   "api",
	Short: "Serve API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return module.Run(
			fx.Provide(
				config.New(cfgFile, version),
				func(db *sql.DB) []check.Checker { return []check.Checker{check.NewSQLChecker(db)} },
			),
			db.Module,
			httpTransport.Module,
			transport.ModuleAPI,
			stringsvc.ModuleHttpAPI,
			converter.Module,
			grpc.Module,
			stringsvc.ModuleGrpcAPI,
			health.Module,
			onboarding.Module,
			ses.Module,
			company.Module,
			user.ModuleHttpAPI,
			userclient.ModuleClient,
			log.Module,
			salesforce.Module,
			auth.Module,
			cognito.Module,
			applications.ModuleHttpAPI,
			websites.ModuleHttpAPI,
			meetings.ModuleHttpAPI,
			address.ModuleHttpAPI,
			externalinfra.ModuleHttpAPI,
			wireless.ModuleHttpAPI,
			signatures.ModuleHttpAPI,
			ipranges.ModuleHttpAPI,
			policies.ModuleHttpAPI,
			vulnerability.ModuleHttpAPI,
			file_converter.ModuleHttpAPI,
			penetrationtesting.ModuleHttpAPI,
			remediation.ModuleHttpAPI,
			questionnaires.ModuleHttpAPI,
			questionnairesClient.ModuleClient,
			frameworks.ModuleHttpAPI,
			frameworksClient.ModuleClient,
			customer_success.ModuleHttpAPI,
			calendly.Module,
			cache.Module,
			securityawareness.ModuleHttpAPI,
			knowbe4.Module,
			jira.Module,
			rapid7.ModuleClient,
			swagger.ModuleServeSwagger,
			webhooks.ModuleHttpAPI,
			s3client.Module,
			fx.NopLogger,
			fx.StartTimeout(time.Second*60),
		)
	},
}

func init() {
	rootCmd.AddCommand(apiCommand)
}
