// Package migrate
package migrate

import (
	"strings"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"

	"github.com/pkg/errors"
)

// Config represent main configuration of service
type Config struct{ DB db.Config }

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if err := c.DB.Validate(); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
