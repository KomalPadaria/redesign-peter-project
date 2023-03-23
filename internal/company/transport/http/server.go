package http

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"strconv"

	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/endpoints"
	svcTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/encode"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for http.
func RegisterTransport(
	server *httpTransport.Server,
	ep *endpoints.Endpoints,
	svcTransportClient svcTransport.Client,

) {
	registerGetCompanyInfo(server, ep.GetCompanyInfoEndpoint, svcTransportClient)
	registerUploadSecurityCampaignUsers(server, ep.UploadSecurityCampaignUsersEndpoint, svcTransportClient)
	registerGetSecurityCampaignUsers(server, ep.GetSecurityCampaignUsersEndpoint, svcTransportClient)
}

func registerGetCompanyInfo(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeGetCompanyInfoRequest,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{"POST"}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{"POST"})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle("POST", "/company/info", handler)
	atc.RegisterAccessControlOptionsHandler(server, "/company/info", []string{"POST"})
}

func registerUploadSecurityCampaignUsers(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/security/awareness/users"
	method := "POST"
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeUploadSecurityCampaignUsers,
		atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetSecurityCampaignUsers(server *httpTransport.Server, ep goKitEndpoint.Endpoint, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/security/awareness/users"
	method := "GET"
	handler := goKitHTTPTransport.NewServer(
		ep,
		decodeGetSecurityCampaignUsers,
		atc.EncodeAccessControlHeadersWrapper(encodeCSVResponse, []string{method}),
		goKitHTTPTransport.ServerErrorEncoder(atc.EncodeErrorControlHeadersWrapper(encode.Error, []string{method})),
		goKitHTTPTransport.ServerErrorHandler(atc.LogErrorHandler()),
	)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func encodeCSVResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	ioi := response.([]byte)

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("sample_campaign_users.csv"))
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("content-length", fmt.Sprintf("%d", binary.Size(ioi)))
	_, err := io.Copy(w, bytes.NewReader(ioi))
	if err != nil {
		return err
	}
	return nil
}
