// Package grpc contains grpc server with all necessary interceptor for logging, tracing, etc
package grpc

import (
	"go.uber.org/zap"
)

// options for client
type options struct {
	logger *zap.SugaredLogger
}

// Option applies option
type Option interface{ apply(*options) }
type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

// WithLogger provides logging for every query
func WithLogger(logger *zap.SugaredLogger) Option {
	return optionFunc(func(o *options) {
		o.logger = logger
	})
}
