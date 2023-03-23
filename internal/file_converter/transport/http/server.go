// Package http for conx.
package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/endpoints"
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
	registerDocx2Html(server, ep.ConvertDocxFile, authClient, svcTransportClient)
	registerHtml2Docx(server, ep.ConvertHtmlFile, authClient, svcTransportClient)
}

func registerDocx2Html(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/converter/docx2html"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	handler := getHandler(securedEp, decodeDocx2Html, atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}), atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerHtml2Docx(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/converter/html2docx"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "policies-procedures", method)
	handler := getHandler(securedEp, decodeHtml2Docx, encodeHtml2DocxResponse, atc, method)

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
