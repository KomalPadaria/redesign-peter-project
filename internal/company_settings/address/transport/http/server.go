// Package http for address.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/endpoints"
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
	registerGetCompanyAddresses(server, ep.GetCompanyAddressesEndpoint, authClient, svcTransportClient)
	registerCreateCompanyAddress(server, ep.CreateCompanyAddressEndpoint, authClient, svcTransportClient)
	registerUpdateCompanyAddress(server, ep.UpdateCompanyAddressEndpoint, authClient, svcTransportClient)
	registerDeleteCompanyAddress(server, ep.DeleteCompanyAddressEndpoint, authClient, svcTransportClient)
	registerUpdateCompanyAddressPatch(server, ep.UpdateCompanyAddressPatchEndpoint, authClient, svcTransportClient)

	registerGetCompanyFacilities(server, ep.GetFacilitiesEndpoint, authClient, svcTransportClient)
}

func registerGetCompanyAddresses(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/addresses"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeGetCompanyAddressesRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCreateCompanyAddress(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/address"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeCreateCompanyAddressRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateCompanyAddress(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/address/{address_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeUpdateCompanyAddressRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteCompanyAddress(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/address/{address_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeDeleteCompanyAddressRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateCompanyAddressPatch(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/address/{address_id}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeUpdateCompanyAddressPatchRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetCompanyFacilities(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/facilities"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "company-settings", method)
	handler := getHandler(securedEp, decodeGetCompanyFacilitiesRequest, atc, method)

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
