// Package http for websites.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/endpoints"
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
	registerGetMeetings(server, ep.GetMeetingsEndpoint, authClient, svcTransportClient)
	registerGetCompanyMeetings(server, ep.GetCompanyMeetingsEndpoint, authClient, svcTransportClient)
	registerCalendlyWebhook(server, ep.CalendlyWebhookEndpoint, authClient, svcTransportClient)
	registerCreateMeetingFromCalendly(server, ep.CreateMeetingFromCalendlyEndpoint, authClient, svcTransportClient)
}

func registerGetMeetings(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/meetings"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeGetMeetingsRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetCompanyMeetings(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/meetings"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeGetCompanyMeetingsRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCalendlyWebhook(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/settings/meetings/webhook"
	method := "POST"
	handler := getHandler(ep, decodeCalendlyWebhookRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCreateMeetingFromCalendly(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/internal/meetings/create"
	method := "GET"
	handler := getHandler(authClient.SecureServiceWithRedesignEndpoint(ep), decodeCreateMeetingFromCalendlysRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func getHandler(ep goKitEndpoint.Endpoint, dec goKitHTTPTransport.DecodeRequestFunc, atc svcTransport.Client, method string) *goKitHTTPTransport.Server {
	return goKitHTTPTransport.NewServer(
		ep,
		dec,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)
}
