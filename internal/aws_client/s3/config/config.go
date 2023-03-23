package config

import (
	"strings"

	"github.com/pkg/errors"
)

// Config for auth.
type Config struct {
	AwsRegion  string
	BucketName string
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if c.AwsRegion == "" {
		errs = append(errs, "AWS region shouldn't be empty")
	}

	if c.BucketName == "" {
		errs = append(errs, "AWS BucketName shouldn't be empty")
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
