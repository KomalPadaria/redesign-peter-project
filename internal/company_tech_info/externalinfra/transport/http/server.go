// Package http for websites.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/endpoints"
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
	registerCreateTechInfoExternalInfra(server, ep.CreateTechInfoExternalInfraEndpoint, authClient, svcTransportClient)
	registerGetAllTechInfoExternalInfras(server, ep.GetAllTechInfoExternalInfrasEndpoint, authClient, svcTransportClient)
	registerUpdateTechInfoExternalInfra(server, ep.UpdateTechInfoExternalInfraEndpoint, authClient, svcTransportClient)
	registerDeleteTechInfoExternalInfra(server, ep.DeleteTechInfoExternalInfraEndpoint, authClient, svcTransportClient)
}

func registerCreateTechInfoExternalInfra(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/external/infra"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeCreateTechInfoExternalInfraRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetAllTechInfoExternalInfras(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/external/infra"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeGetAllTechInfoExternalInfrasRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateTechInfoExternalInfra(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/external/infra/{external_infra_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateTechInfoExternalInfraRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteTechInfoExternalInfra(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/external/infra/{external_infra_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeDeleteTechInfoExternalInfraRequest, atc, method)

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
