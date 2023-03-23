// Package grpcweb contains grpc-web wrapper which wraps grpc server and serve it with grpc-web protocol through http server
package grpcweb

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/meta"
	grpcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/grpc"
	http2 "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Options holds options for server module
type Options struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *http2.Server
	Logger     *zap.SugaredLogger
	HTTPConfig http2.Config
	GRPCConfig grpcTransport.Config
}

// URLPrefix is a prefix in URL for grpc-web requests
const URLPrefix = "/grpc-web"

// NewModule registers GRPC web wrapper
func NewModule(lc fx.Lifecycle, s fx.Shutdowner, p Options) error { //nolint:gocritic
	wrappedGrpc := grpcweb.WrapServer(p.GRPCServer)

	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		req.URL.Path = strings.TrimPrefix(req.URL.Path, URLPrefix)
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")
		req = req.WithContext(meta.WithTransport(req.Context(), "GRPCWEB"))

		wrappedGrpc.ServeHTTP(resp, req)
	})

	methods := []string{
		http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace,
	}

	for _, m := range methods {
		p.HTTPServer.Handle(m, URLPrefix+"/{_ignore:.*}", handler)
	}

	// in order to avoid panic described here we need to stop http server before grpc server
	// https://github.com/grpc/grpc-go/issues/1384
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			p.Logger.Infow("starting HTTP server", "port", p.HTTPConfig.Port)

			go func() {
				defer s.Shutdown() //nolint:errcheck

				if err := p.HTTPServer.Serve(); err != nil {
					if err != http.ErrServerClosed {
						p.Logger.Errorw("serve HTTP server", "error", err)
					}
				}
			}()

			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.GRPCConfig.Port))
			if err != nil {
				return err
			}

			p.Logger.Infow("starting GRPC server", "port", p.GRPCConfig.Port)

			go func() {
				defer s.Shutdown() //nolint:errcheck

				if err := p.GRPCServer.Serve(listener); err != nil {
					p.Logger.Errorw("serve GRPC server", "error", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Info("stopping HTTP server")

			if err := p.HTTPServer.Stop(ctx); err != nil {
				p.Logger.Errorw("stop HTTP server", "error", err)
			}

			p.Logger.Info("stopping GRPC server")
			p.GRPCServer.GracefulStop()

			return nil
		},
	})

	return nil
}

// Module for uber fx
// !NOTE: You must use grpc.ModuleWithoutLifecycle,	http.ModuleWithoutLifecycle, not grpc.Module, http.Module
var Module = fx.Options(
	fx.Invoke(
		NewModule,
	),
)
