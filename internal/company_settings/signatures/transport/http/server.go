// Package http for websites.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/endpoints"
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
	registerGetAddresses(server, ep.GetSignaturesEndpoint, authClient, svcTransportClient)
	registerUpdateStatus(server, ep.UpdateStatusEndpoint, authClient, svcTransportClient)
	registerViewDocument(server, ep.ViewSignaturesDocumentEndpoint, authClient, svcTransportClient)
	registerWebhook(server, ep.WebhookEndpoint, authClient, svcTransportClient)
}

func registerGetAddresses(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/signatures"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeGetSignaturesRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateStatus(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/signatures/{company_signature_uuid}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeUpdateStatusRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerViewDocument(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/signatures/{signature_uuid}/view"
	method := "GET"
	//TODO make this endpoint secured
	//securedEp := authClient.SecureServiceWithCognitoEndpoint(ep)
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeViewDocumentRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.ViewFileResponse, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)
	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerWebhook(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/docusign/webhook"
	method := "POST"
	handler := getHandler(ep, decodeWebhookRequest, atc, method)

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
