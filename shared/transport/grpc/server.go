// Package grpc contains grpc server with all necessary interceptor for logging, tracing, etc
package grpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/grpc/interceptors/logging"
	"google.golang.org/grpc"
)

// NewServer returns new GRPC servers with all interceptor like tracer, logger, metrics, etc.
func NewServer(opts ...Option) *grpc.Server {
	defaultOptions := options{}

	for _, o := range opts {
		o.apply(&defaultOptions)
	}

	var dialOpts []grpc.ServerOption

	var srvUnaryInterceptors []grpc.UnaryServerInterceptor

	var srvStreamInterceptors []grpc.StreamServerInterceptor

	if defaultOptions.logger != nil {
		srvUnaryInterceptors = append(srvUnaryInterceptors, logging.UnaryServerInterceptor(defaultOptions.logger))
		srvStreamInterceptors = append(srvStreamInterceptors, logging.StreamServerInterceptor(defaultOptions.logger))
	}

	dialOpts = append(dialOpts,
		grpc_middleware.WithUnaryServerChain(srvUnaryInterceptors...),
		grpc_middleware.WithStreamServerChain(srvStreamInterceptors...),
	)

	return grpc.NewServer(dialOpts...)
}
