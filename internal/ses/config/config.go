package config

import (
	"strings"

	"github.com/pkg/errors"
)

// Config for ses.
type Config struct {
	AwsRegion         string
	InviteSenderEmail string
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if c.AwsRegion == "" {
		errs = append(errs, "AWS region shouldn't be empty")
	}

	if c.InviteSenderEmail == "" {
		errs = append(errs, "Invite Sender Email shouldn't be empty")
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
