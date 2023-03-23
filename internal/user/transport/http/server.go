// Package http for conx.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/encode"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/endpoints"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for http.
func RegisterTransport(
	server *httpTransport.Server,
	ep *endpoints.Endpoints,
	authClient auth.Client,
	svcTransportClient svcTransport.Client,

) {
	registerCreateCompanyAndUser(server, ep.CreateCompanyUserEndpoint, svcTransportClient)
	registerGetContextUserCompany(server, ep.GetContextUserCompanyInfoEndpoint, authClient, svcTransportClient)
	registerGetSFUserAndCompanyInfo(server, ep.GetSFUserAndCompanyInfoEndpoint, svcTransportClient)
	registerActivateUserByEmail(server, ep.ActivateUserByEmailEndpoint, svcTransportClient)
	registerGetContextUserCompanyInternal(server, ep.GetContextUserCompanyInfoInternalEndpoint, authClient, svcTransportClient)
	registerGetCompanyUsers(server, ep.GetCompanyUsersEndpoint, authClient, svcTransportClient)
	registerCreateUser(server, ep.CreateUserEndpoint, authClient, svcTransportClient)
	registerUpdateCompanyUserLink(server, ep.UpdateCompanyUserLinkEndpoint, authClient, svcTransportClient)
	registerDeleteCompanyUserLink(server, ep.DeleteCompanyUserLinkEndpoint, authClient, svcTransportClient)
	registerResendUserInvite(server, ep.ResendInviteEndpoint, authClient, svcTransportClient)
	registerSwitchCompany(server, ep.SwitchCompanyEndpoint, authClient, svcTransportClient)
	registerUpdateUser(server, ep.UpdateUserEndpoint, authClient, svcTransportClient)
	registerListCompanies(server, ep.ListCompaniesForUserEndpoint, authClient, svcTransportClient)
}

func registerCreateCompanyAndUser(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	method := "POST"
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeCreateCompanyAndUserRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("POST", "/users", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/users", []string{method})
}

func registerGetContextUserCompany(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "GET"
	path := "/users/me"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "all", method),
		decodeGetContextUserCompanyRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("GET", path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateUser(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "PUT"
	path := "/users/me"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "all", method),
		decodeUpdateUserRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("PUT", path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetContextUserCompanyInternal(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "GET"
	path := "/internal/users/me"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithRedesignEndpoint(ep),
		decodeGetContextUserCompanyRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("GET", path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerActivateUserByEmail(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	method := "POST"
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeActivateUserByExternalIdRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("POST", "/users/activate", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/users/activate", []string{method})
}

func registerGetSFUserAndCompanyInfo(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	method := "POST"
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeGetSFUserAndCompanyInfoRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("POST", "/sf/users", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/sf/users", []string{method})
}

func registerGetCompanyUsers(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "GET"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "account-management", method),
		decodeGetCompanyUsersRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("GET", "/companies/{company_id}/users/{user_id}/settings/users", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/companies/{company_id}/users/{user_id}/settings/users", []string{method})
}

func registerCreateUser(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "POST"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "account-management", method),
		decodeCreateUserRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("POST", "/companies/{company_id}/users/{user_id}/settings/users", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/companies/{company_id}/users/{user_id}/settings/users", []string{method})
}

func registerUpdateCompanyUserLink(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "POST"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "account-management", method),
		decodeUpdateCompanyUserLinkRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("PUT", "/companies/{company_id}/users/{req_user_id}/settings/users/{user_id}", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/companies/{company_id}/users/{req_user_id}/settings/users/{user_id}", []string{method})
}

func registerDeleteCompanyUserLink(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "GET"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "account-management", method),
		decodeDecodeCompanyUserLinkRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("DELETE", "/companies/{company_id}/users/{req_user_id}/settings/users/{user_id}", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/companies/{company_id}/users/{req_user_id}/settings/users/{user_id}", []string{"DELETE"})
}

func registerResendUserInvite(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "POST"
	path := "/companies/{company_id}/users/{user_id}/settings/users/invite"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "account-management", method),
		decodeResendUserInviteRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("POST", path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerSwitchCompany(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "PATCH"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "account-management", method),
		decodeSwitchCompanyRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{"PATCH"}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("PATCH", "/users/{user_id}/companies/{company_id}/switch", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/users/{user_id}/companies/{company_id}/switch", []string{"PATCH"})
}

func registerListCompanies(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	method := "GET"
	handler := goKitHTTPTransport.NewServer(
		authClient.SecureServiceWithCognitoEndpoint(ep, "all", method),
		decodeListCompaniesRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("GET", "/users/{user_id}/companies", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/users/{user_id}/companies", []string{method})
}
