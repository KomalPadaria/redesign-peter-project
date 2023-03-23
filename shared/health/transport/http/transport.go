// Package http contains health http transport
package http

import (
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/health/endpoint"

	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for health
func RegisterTransport(server *httpTransport.Server, e *endpoint.Endpoints) {
	registerCheckTransport(server, e, "/health", "GET")
	registerCheckTransport(server, e, "/", "GET")
}

func registerCheckTransport(server *httpTransport.Server, e *endpoint.Endpoints, path, method string) {
	handler := goKitHTTPTransport.NewServer(
		e.CheckEndpoint,
		goKitHTTPTransport.NopRequestDecoder,
		encodeResponse,
		goKitHTTPTransport.ServerErrorEncoder(encodeError),
	)

	server.Handle(method, path, handler)
}
