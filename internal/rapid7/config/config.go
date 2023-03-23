package config

import (
	"strings"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/pkg/errors"
)

type Config struct {
	DataWarehouse db.Config
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
