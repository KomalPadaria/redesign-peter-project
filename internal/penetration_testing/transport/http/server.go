package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/endpoints"
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
	registerGetPenetrationTests(server, ep.GetPenetrationTestsEndpoint, authClient, svcTransportClient)
	registerGetPenetrationTestStats(server, ep.GetPenetrationTestStatsEndpoint, authClient, svcTransportClient)
}

func registerGetPenetrationTests(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/penetration/topremediation"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "penetration-test", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetVulnerabilitiesRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetPenetrationTestStats(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/penetration/stats"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "penetration-test", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetVulnerabilityStatsRequest, encoder, atc, method)

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
