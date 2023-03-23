// Package http for conx.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/endpoints"
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
	registerGetAllTechInfoIpRange(server, ep.GetAllTechInfoIpRangeEndpoint, authClient, svcTransportClient)
	registerCreateTechInfoIpRange(server, ep.CreateTechInfoIpRangeEndpoint, authClient, svcTransportClient)
	registerUpdateTechInfoIpRange(server, ep.UpdateTechInfoIpRangeEndpoint, authClient, svcTransportClient)
	registerUpdateTechInfoIpRangePatch(server, ep.UpdateTechInfoIpRangePatchEndpoint, authClient, svcTransportClient)
	registerDeleteTechInfoIpRange(server, ep.DeleteTechInfoIpRangeEndpoint, authClient, svcTransportClient)
}

func registerGetAllTechInfoIpRange(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/ips"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeGetAllTechInfoIpRangeRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerCreateTechInfoIpRange(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/ip"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeCreateTechInfoIpRangeRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateTechInfoIpRange(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/ip/{ip_range_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateTechInfoIpRangeRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateTechInfoIpRangePatch(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/ip/{ip_range_id}"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeUpdateTechInfoIpRangePatchRequest, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteTechInfoIpRange(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/settings/assessments/ip/{ip_range_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "tech-info", method)
	handler := getHandler(securedEp, decodeDeleteTechInfoIpRangeRequest, atc, method)

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
