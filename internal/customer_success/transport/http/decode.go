package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.UpdateServiceReviewsStatusRequestBody
}

func decodeGetSubscriptionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	req := &entities.GetSubscriptionsRequest{
		CompanyUuid: compUUID,
	}

	return req, nil
}
func decodeGetSubscriptionsPlansRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	req := &entities.GetSubscriptionPlansRequest{
		CompanyUuid: compUUID,
	}

	return req, nil
}

func decodeUploadReportsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceName := params["service_name"]
	if strings.TrimSpace(serviceName) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_name")
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	return &entities.UploadReportsRequest{
		CompanyUuid: compUUID,
		ServiceName: serviceName,
		Files:       r.MultipartForm.File["files"],
	}, nil
}

func decodeDownloadReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceName := params["service_name"]
	if strings.TrimSpace(serviceName) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_name")
	}

	reportName := params["report_name"]
	if strings.TrimSpace(reportName) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("report_name")
	}

	return &entities.DownloadReportRequest{
		CompanyUuid: compUUID,
		ServiceName: serviceName,
		ReportName:  reportName,
	}, nil
}

func decodeDeleteReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceName := params["service_name"]
	if strings.TrimSpace(serviceName) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_name")
	}

	reportName := params["report_name"]
	if strings.TrimSpace(reportName) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("report_name")
	}

	return &entities.DeleteReportRequest{
		CompanyUuid: compUUID,
		ServiceName: serviceName,
		ReportName:  reportName,
	}, nil
}

func decodeGetConsultingHoursRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetConsultingHoursRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}
func decodeGetConsumedHourRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetConsumedHoursRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}
func decodeGetServiceReviewRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetServiceReviewsRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeUpdateServiceReviewStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	evidenceUUID, err := uuid.Parse(params["evidence_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("evidence_id")
	}

	body := &entities.UpdateServiceReviewsStatusRequestBody{}
	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateServiceReviewsStatusRequest{
		CompanyUuid:      compUUID,
		UserUuid:         userUUID,
		EvidenceUuid:     evidenceUUID,
		PatchRequestBody: body,
	}, nil
}

func decodeUploadEvidencesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceUUID, err := uuid.Parse(params["service_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_id")
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	return &entities.UploadEvidencesRequest{
		CompanyUuid: compUUID,
		ServiceUuid: serviceUUID,
		Files:       r.MultipartForm.File["files"],
	}, nil
}

func decodeDeleteServiceEvidenceReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceUUID, err := uuid.Parse(params["service_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_id")
	}

	evidenceUuid, err := uuid.Parse(params["evidence_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("evidence_id")
	}

	reportUuid, err := uuid.Parse(params["report_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("report_id")
	}

	return &entities.DeleteEvidenceFileRequest{
		CompanyUuid:  compUUID,
		ServiceUuid:  serviceUUID,
		EvidenceUuid: evidenceUuid,
		FileUuid:     reportUuid,
	}, nil
}

func decodeAddEvidenceReportsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceUUID, err := uuid.Parse(params["service_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_id")
	}

	evidenceUuid, err := uuid.Parse(params["evidence_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("evidence_id")
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	return &entities.AddEvidenceFilesRequest{
		CompanyUuid:  compUUID,
		ServiceUuid:  serviceUUID,
		EvidenceUuid: evidenceUuid,
		Files:        r.MultipartForm.File["files"],
	}, nil
}

func decodeDownloadEvidenceReportRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	serviceUUID, err := uuid.Parse(params["service_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("service_id")
	}

	evidenceUuid, err := uuid.Parse(params["evidence_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("evidence_id")
	}

	reportName := params["report_name"]
	if strings.TrimSpace(reportName) == "" {
		return nil, httpError.NewErrBadOrInvalidPathParameter("report_name")
	}

	return &entities.DownloadEvidenceReportRequest{
		CompanyUuid:  compUUID,
		ServiceUuid:  serviceUUID,
		EvidenceUuid: evidenceUuid,
		ReportName:   reportName,
	}, nil
}
