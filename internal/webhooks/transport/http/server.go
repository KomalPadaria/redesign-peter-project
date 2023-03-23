// Package http for websites.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/encode"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/endpoints"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for http.
func RegisterTransport(
	server *httpTransport.Server,
	ep *endpoints.Endpoints,
	authClient auth.Client,
	svcTransportClient svcTransport.Client,
) {
	registerSFAccountWebhook(server, ep.SFAccountWebhookEndpoint, authClient, svcTransportClient)
	registerSFAccountSubscriptionsWebhook(server, ep.SFAccountSubscriptionsWebhookEndpoint, authClient, svcTransportClient)
}

func registerSFAccountWebhook(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/webhook/sf/account"
	method := "POST"
	securedEp := authClient.SecureServiceWithRedesignWebhookEndpoint(ep)
	handler := getHandler(securedEp, decodeSFAccountWebhookRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerSFAccountSubscriptionsWebhook(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/webhook/sf/account/subscriptions"
	method := "POST"
	securedEp := authClient.SecureServiceWithRedesignWebhookEndpoint(ep)
	handler := getHandler(securedEp, decodeSFAccountSubscriptionsWebhookRequest, atc, method)

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
