package config

import (
	"strings"

	"github.com/pkg/errors"
)

// Config is a common service config
type Config struct {
	Username   string
	ApiToken   string
	ProjectKey string
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if c.Username == "" {
		errs = append(errs, "JIRA Username shouldn't be empty")
	}

	if c.ApiToken == "" {
		errs = append(errs, "JIRA ApiToken shouldn't be empty")
	}

	if c.ProjectKey == "" {
		errs = append(errs, "JIRA ProjectKey shouldn't be empty")
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
