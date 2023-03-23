// Package http for conx.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/endpoints"
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
	registerGetAllPolicies(server, ep.GetAllPoliciesEndpoint, authClient, svcTransportClient)
	registerCreatePolicy(server, ep.CreatePolicyEndpoint, authClient, svcTransportClient)
	registerGetPolicyDocument(server, ep.GetPolicyDocumentEndpoint, authClient, svcTransportClient)
	registerSaveDocument(server, ep.SaveDocumentEndpoint, authClient, svcTransportClient)
	registerGetDocument(server, ep.GetDocumentEndpoint, authClient, svcTransportClient)
	registerGetPolicyHistoriesByPolicy(server, ep.GetPolicyHistoriesByPolicyEndpoint, authClient, svcTransportClient)
	registerUpdatePolicyDocumentPatch(server, ep.UpdatePolicyDocumentStatusEndpoint, authClient, svcTransportClient)
	registerDeletePolicy(server, ep.DeletePolicyEndpoint, authClient, svcTransportClient)
	registerGetStats(server, ep.GetPoliciesStatsEndpoint, authClient, svcTransportClient)
	registerGetTemplates(server, ep.GetTemplatesEndpoint, authClient, svcTransportClient)
	registerCreateDocumentFromTemplate(server, ep.CreateDocumentFromTemplateEndpoint, authClient, svcTransportClient)
}

func registerGetAllPolicies(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policies"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetAllPoliciesRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCreatePolicy(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeCreatePolicyRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetPolicyDocument(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{policy_id}"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetPolicyDocumentRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerSaveDocument(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{policy_id}/document"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeSaveDocumentRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetDocument(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{policy_id}/document"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encodeGetDocumentResponse, []string{method})
	handler := getHandler(securedEp, decodeGetDocumentRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetPolicyHistoriesByPolicy(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{policy_id}/history"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetPolicyHistoriesByPolicyRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdatePolicyDocumentPatch(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{policy_id}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeUpdatePolicyDocumentStatusRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeletePolicy(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{policy_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeDeletePolicyRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetStats(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policies/stats"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetPoliciesStatsRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetTemplates(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/policies/templates"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetTemplatesRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCreateDocumentFromTemplate(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/policy/{template_id}"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeCreateDocumentFromTemplateRequest, encoder, atc, method)

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
