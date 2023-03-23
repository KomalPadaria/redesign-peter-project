package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/endpoints"
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
	registerGetSubscriptions(server, ep.GetSubscriptionsEndpoint, authClient, svcTransportClient)
	registerGetSubscriptionPlans(server, ep.GetSubscriptionPlansEndpoint, authClient, svcTransportClient)
	registerUploadReports(server, ep.UploadReportsEndpoint, authClient, svcTransportClient)
	registerDownloadServiceReport(server, ep.DownloadServiceReportEndpoint, authClient, svcTransportClient)
	registerDeleteServiceReport(server, ep.DeleteServiceReportEndpoint, authClient, svcTransportClient)
	registerGetConsultingHours(server, ep.GetConsultingHoursEndpoint, authClient, svcTransportClient)
	registerGetServiceReview(server, ep.GetServiceReviewEndpoint, authClient, svcTransportClient)
	registerUpdateServiceReviewStatus(server, ep.UpdateServiceReviewStatusEndpoint, authClient, svcTransportClient)
	registerGetConsumedHours(server, ep.GetConsumedHoursEndpoint, authClient, svcTransportClient)

	registerUploadEvidenceReports(server, ep.UploadEvidenceEndpoint, authClient, svcTransportClient)
	registerDeleteServiceEvidenceReport(server, ep.DeleteEvidenceFileEndpoint, authClient, svcTransportClient)
	registerAddEvidenceReports(server, ep.AddEvidenceFilesEndpoint, authClient, svcTransportClient)
	registerDownloadEvidenceReport(server, ep.DownloadEvidenceReportEndpoint, authClient, svcTransportClient)
}

func registerGetSubscriptions(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/subscriptions"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetSubscriptionsRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetSubscriptionPlans(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/subscriptions/plan"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetSubscriptionsPlansRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUploadReports(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/subscriptions/{service_name}/reports"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeUploadReportsRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDownloadServiceReport(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/subscriptions/{service_name}/reports/{report_name}"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.DownloadFileResponse, []string{method})
	handler := getHandler(securedEp, decodeDownloadReportRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteServiceReport(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/subscriptions/{service_name}/reports/{report_name}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeDeleteReportRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetConsultingHours(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/consulting/hours"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetConsultingHoursRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetServiceReview(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/services-review"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetServiceReviewRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateServiceReviewStatus(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/services-review/{evidence_id}/acknowledge"
	method := "PATCH"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})

	handler := getHandler(securedEp, decodeUpdateServiceReviewStatusRequest, encoder, atc, method)
	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUploadEvidenceReports(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/services-review/{service_id}/evidence/reports"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeUploadEvidencesRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDownloadEvidenceReport(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/services-review/{service_id}/evidence/{evidence_id}/reports/{report_name}"
	method := "GET"
	//securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.DownloadFileResponse, []string{method})
	handler := getHandler(ep, decodeDownloadEvidenceReportRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerAddEvidenceReports(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/services-review/{service_id}/evidence/{evidence_id}/reports"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeAddEvidenceReportsRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteServiceEvidenceReport(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/services-review/{service_id}/evidence/{evidence_id}/reports/{report_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeDeleteServiceEvidenceReportRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetConsumedHours(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/consulting/consumed-hour"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "customer-success", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetConsumedHourRequest, encoder, atc, method)

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
