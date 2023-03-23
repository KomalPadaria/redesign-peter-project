package config

import (
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
	ApiHost     string
	AccessToken string
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if c.ApiHost == "" {
		errs = append(errs, "Calendly API host shouldn't be empty")
	}

	if c.AccessToken == "" {
		errs = append(errs, "Calendly access token shouldn't be empty")
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
