// Package grpc contains grpc server with all necessary interceptor for logging, tracing, etc
package grpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// ServerOptions holds options for server module
type ServerOptions struct {
	fx.In

	Config Config
	Logger *zap.SugaredLogger
}

// NewServerModuleWithoutLifecycle returns new module for uber fx without lifecycle hooks
// nolint:gocritic
func NewServerModuleWithoutLifecycle(p ServerOptions) *grpc.Server {
	opts := []Option{
		WithLogger(p.Logger),
	}

	server := NewServer(opts...)

	reflection.Register(server)

	return server
}

// NewServerModule returns server module for uber fx
// nolint:gocritic
func NewServerModule(lc fx.Lifecycle, s fx.Shutdowner, p ServerOptions) (*grpc.Server, error) {
	server := NewServerModuleWithoutLifecycle(p)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.Config.Port))
			if err != nil {
				return err
			}

			p.Logger.Infow("starting GRPC server", "port", p.Config.Port)

			go func() {
				defer s.Shutdown() //nolint:errcheck

				if err := server.Serve(listener); err != nil {
					p.Logger.Errorw("serve GRPC server", "error", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Info("stopping GRPC server")
			server.GracefulStop()

			return nil
		},
	})

	return server, nil
}

// Modules for uber fx
var (
	Module = fx.Options(
		fx.Provide(
			NewServerModule,
		),
	)
	ModuleWithoutLifecycle = fx.Options(
		fx.Provide(
			NewServerModuleWithoutLifecycle,
		),
	)
)
