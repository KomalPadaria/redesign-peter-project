// Package http contains http client/server with all necessary interceptor for logging, tracing, etc
package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http/interceptors/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http/interceptors/logging"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http/interceptors/meta"
)

// Server defines the HTTP server
type Server struct {
	server  *http.Server
	router  *mux.Router
	options options
}

// Serve is blocking serving of HTTP requests
func (s *Server) Serve() error {
	s.registerHandlers()

	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the server from HTTP connections.
func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	s.server.SetKeepAlivesEnabled(false)

	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandlers() {
	var next http.Handler = s.router

	if s.options.logger != nil {
		next = logging.ServerHandler(s.router, s.options.logger, next)
	}

	next = auth.ServerHandler(next)
	next = meta.UserAgentServerHandler(next)
	next = meta.RequestIDServerHandler(next)

	s.server.Handler = next
}

// Handle the method and path with the handler
func (s *Server) Handle(method, path string, handler http.Handler) {
	s.router.Handle(path, handler).Methods(method)
}

// HandleWithPathPrefix path with the handler
func (s *Server) HandleWithPathPrefix(path string, handler http.Handler) {
	s.router.PathPrefix(path).Handler(handler)
}

// HandleNotFound when no route matches
func (s *Server) HandleNotFound(handler http.Handler) {
	s.router.NotFoundHandler = handler
}

// NewServer returns new HTTP server with all interceptor like tracer, logger, metrics, etc.
func NewServer(address string, opts ...Option) *Server {
	options := options{} //nolint:govet

	for _, o := range opts {
		o.apply(&options)
	}

	router := mux.NewRouter().StrictSlash(true)
	httpServer := &http.Server{
		Addr: address,
	}

	server := &Server{
		router:  router,
		server:  httpServer,
		options: options,
	}

	return server
}
