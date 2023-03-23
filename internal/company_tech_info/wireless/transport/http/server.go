// Package http for websites.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/endpoints"
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
	registerCreateTechInfoWireless(server, ep.CreateTechInfoWirelessEndpoint, authClient, svcTransportClient)
	registerGetAllTechInfoWirelesss(server, ep.GetAllTechInfoWirelesssEndpoint, authClient, svcTransportClient)
	registerUpdateTechInfoWireless(server, ep.UpdateTechInfoWirelessEndpoint, authClient, svcTransportClient)
	registerDeleteTechInfoWireless(server, ep.DeleteTechInfoWirelessEndpoint, authClient, svcTransportClient)
	registerUpdateTechInfoWirelessPatch(server, ep.UpdateTechInfoWirelessPatchEndpoint, authClient, svcTransportClient)
}

func registerCreateTechInfoWireless(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/wireless"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeCreateTechInfoWirelessRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetAllTechInfoWirelesss(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/wireless"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeGetAllTechInfoWirelesssRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateTechInfoWireless(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/wireless/{wireless_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateTechInfoWirelessRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteTechInfoWireless(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/wireless/{wireless_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeDeleteTechInfoWirelessRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateTechInfoWirelessPatch(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/wireless/{wireless_id}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateTechInfoWirelessPatchRequest, atc, method)

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
