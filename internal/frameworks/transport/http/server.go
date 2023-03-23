package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/endpoints"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/encode"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for http.
func RegisterTransport(
	server *httpTransport.Server,
	ep *endpoints.Endpoints,
	authClient auth.Client,
	svcTransportClient svcTransport.Client,
) {
	registerGetFrameworks(server, ep.GetFrameworksEndpoint, authClient, svcTransportClient)
	registerGetFrameworkControls(server, ep.GetFrameworkControlsEndpoint, authClient, svcTransportClient)
	registerGetFrameworkStats(server, ep.GetFrameworkStatsEndpoint, authClient, svcTransportClient)
}

func registerGetFrameworks(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/frameworks"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetFrameworksRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetFrameworkControls(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/frameworks/{framework_id}/controls"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetFrameworkControlRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetFrameworkStats(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/frameworks/stats"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetFrameworkStatsRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func getHandler(ep goKitEndpoint.Endpoint, dec goKitHTTPTransport.DecodeRequestFunc, enc goKitHTTPTransport.EncodeResponseFunc, atc svcTransport.Client, method string) *goKitHTTPTransport.Server {
	return goKitHTTPTransport.NewServer(
		ep,
		dec,
		enc,
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)
}
