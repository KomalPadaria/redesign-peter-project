package config

import (
	"strings"

	cognitoCfg "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/config"
	s3 "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/config"
	calendly "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/config"
	jira "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/config"
	rapid7Config "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/rapid7/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/config"
	sesCfg "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses/config"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/log"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

// Config is represented main configuration of service
type Config struct {
	fx.Out

	Common                    cfg.Config
	Transport                 transport.Config
	Logger                    log.Config
	DB                        db.Config
	Salesforce                config.Config
	Cognito                   cognitoCfg.Config
	Calendly                  calendly.Config
	Rapid7                    rapid7Config.Config
	Ses                       sesCfg.Config
	AccessControlAllowOrigins svcTransport.AccessControlAllowOrigins
	Jira                      jira.Config
	S3                        s3.Config
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	validatables := []cfg.Validatable{
		&c.Common, &c.Transport.GRPC, &c.Logger, &c.Salesforce, &c.Calendly, &c.Rapid7, &c.Ses, &c.Jira,
	}

	if err := cfg.ValidateConfigs(validatables...); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}

// New config.
func New(path, version string) func() (Config, error) {
	return func() (Config, error) {
		cfgFile := Config{}
		cfgFile.Common.Version = version

		err := cfg.Init("config", path, &cfgFile)
		if err != nil {
			return cfgFile, err
		}

		return cfgFile, nil
	}
}
