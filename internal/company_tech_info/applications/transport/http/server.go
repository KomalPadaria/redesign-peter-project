// Package http for conx.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/endpoints"
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
	registerCreateApplication(server, ep.CreateApplicationEndpoint, authClient, svcTransportClient)
	registerUpdateApplication(server, ep.UpdateApplicationEndpoint, authClient, svcTransportClient)
	registerDeleteApplication(server, ep.DeleteApplicationEndpoint, authClient, svcTransportClient)
	registerGetAllApplications(server, ep.GetAllApplicationsEndpoint, authClient, svcTransportClient)
	registerUpdateApplicationPatch(server, ep.UpdateApplicationPatchEndpoint, authClient, svcTransportClient)

	registerCreateApplicationEnv(server, ep.CreateApplicationEnvEndpoint, authClient, svcTransportClient)
	registerUpdateApplicationEnv(server, ep.UpdateApplicationEnvEndpoint, authClient, svcTransportClient)
	registerUpdateApplicationEnvPatch(server, ep.UpdateApplicationEnvPatchEndpoint, authClient, svcTransportClient)
	registerDeleteApplicationEnv(server, ep.DeleteApplicationEnvEndpoint, authClient, svcTransportClient)
}

func registerCreateApplication(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeCreateApplicationRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetAllApplications(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/applications"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeGetAllApplicationsRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateApplication(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateApplicationRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateApplicationPatch(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateApplicationPatchRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteApplication(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeDeleteApplicationRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCreateApplicationEnv(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}/env"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeCreateApplicationEnvRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateApplicationEnv(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}/env/{env_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateApplicationEnvRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateApplicationEnvPatch(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}/env/{env_id}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateApplicationEnvPatchRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteApplicationEnv(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/application/{application_id}/env/{env_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeDeleteApplicationEnvRequest, atc, method)

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
