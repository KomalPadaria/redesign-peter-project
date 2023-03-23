// Package grpc contains grpc server with all necessary interceptor for logging, tracing, etc
package grpc

import (
	"strings"

	"github.com/pkg/errors"
)

// Config defines how we run server
type Config struct {
	Port int
}

// Validate config
func (c *Config) Validate() error {
	var errs []string

	if c.Port <= 0 {
		errs = append(errs, "Port shouldn't be empty")
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}
