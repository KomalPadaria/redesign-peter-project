// Package http for conx.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/endpoints"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for http.
func RegisterTransport(
	server *httpTransport.Server,
	ep *endpoints.Endpoints,
	svcTransportClient svcTransport.Client,
) {
	registerUppercase(server, ep.UppercaseEndpoint, svcTransportClient)
	registerCount(server, ep.CountEndpoint, svcTransportClient)
}

func registerUppercase(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	handler := getHandler(ep, atc, decodeUppercaseRequest)
	server.Handle("POST", "/redesign/string/uppercase", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/redesign/string/uppercase", []string{"POST"})
}

func registerCount(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	handler := getHandler(ep, atc, decodeCountRequest)
	server.Handle("POST", "/redesign/string/count", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/redesign/string/count", []string{"POST"})
}

func getHandler(ep goKitEndpoint.Endpoint, atc svcTransport.Client, dec goKitHTTPTransport.DecodeRequestFunc) *goKitHTTPTransport.Server {
	return goKitHTTPTransport.NewServer(
		ep,
		dec,
		atc.EncodeAccessControlHeadersWrapper(encodeResponse, []string{"POST"}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encodeError, []string{"POST"})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)
}
