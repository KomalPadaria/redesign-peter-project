package http

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goKitHTTPTransport "github.com/go-kit/kit/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/endpoints"
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
	registerGetCategories(server, ep.GetQuestionnairesEndpoint, authClient, svcTransportClient)
	registerGetQuestionnairesByCategory(server, ep.GetQuestionnairesByCategoryEndpoint, authClient, svcTransportClient)
	registerPostAnswer(server, ep.PostAnswerEndpoint, authClient, svcTransportClient)
	registerAddEngineerFeedback(server, ep.EngineerFeedbackEndpoint, authClient, svcTransportClient)
	registerSubmitQuestionnaires(server, ep.SubmitQuestionnaires, authClient, svcTransportClient)
	registerAddAnswerWithEvidence(server, ep.AddAnswerWithEvidenceEndpoint, authClient, svcTransportClient)
	registerUpdateAnswerWithEvidence(server, ep.UpdateAnswerWithEvidenceEndpoint, authClient, svcTransportClient)
	registerDownloadAnswerEvidence(server, ep.DownloadEvidenceEndpoint, authClient, svcTransportClient)
	registerDeleteAnswerEvidence(server, ep.DeleteEvidenceEndpoint, authClient, svcTransportClient)
}

func registerGetCategories(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/categories"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetQuestionnairesRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerGetQuestionnairesByCategory(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/categories/{category}/questionnaires"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeGetQuestionnairesByCategoryRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerPostAnswer(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/answers"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeAddAnswerRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerAddEngineerFeedback(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/{questionnaire_id}/answers/{answer_id}/feedback"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeAddEngineerFeedbackRequest, encoder, atc, method)
	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerSubmitQuestionnaires(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/submit"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeSubmitQuestionnairesRequest, encoder, atc, method)
	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerAddAnswerWithEvidence(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/{questionnaire_id}/answer"
	method := "POST"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeUploadEvidencesRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerUpdateAnswerWithEvidence(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/{questionnaire_id}/answers/{answer_id}"
	method := "PUT"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeUpdateAnswerRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDownloadAnswerEvidence(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/{questionnaire_id}/answers/{answer_id}/files/{file_id}"
	method := "GET"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.DownloadFileResponse, []string{method})
	handler := getHandler(securedEp, decodeDownloadEvidenceRequest, encoder, atc, method)

	server.Handle(method, path, handler)
	atc.RegisterAccessControlOptionsHandler(server, path, []string{method})
}

func registerDeleteAnswerEvidence(server *httpTransport.Server, ep goKitEndpoint.Endpoint, authClient auth.Client, atc svcTransport.Client) {
	path := "/companies/{company_id}/users/{user_id}/questionnaires/{questionnaire_id}/answers/{answer_id}/files/{file_id}"
	method := "DELETE"
	securedEp := authClient.SecureServiceWithCognitoEndpoint(ep, "gap-analysis", method)
	encoder := atc.EncodeAccessControlHeadersWrapper(encode.Response, []string{method})
	handler := getHandler(securedEp, decodeDeleteEvidenceRequest, encoder, atc, method)

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
