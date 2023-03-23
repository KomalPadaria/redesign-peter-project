// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.UpdateStatusRequestBody | entities.DocusignWebhookData
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeGetSignaturesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetSignaturesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeUpdateStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	companySignatureUuid, err := uuid.Parse(params["company_signature_uuid"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_signature_uuid")
	}

	body := &entities.UpdateStatusRequestBody{}
	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateStatusRequest{
		CompanyUuid:          compUUID,
		UserUuid:             userUUID,
		CompanySignatureUuid: companySignatureUuid,
		PatchRequestBody:     body,
	}, nil
}

func decodeViewDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	signatureUuid, err := uuid.Parse(params["signature_uuid"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_signature_uuid")
	}

	return &entities.ViewDocumentRequest{
		CompanyUuid:   compUUID,
		SignatureUuid: signatureUuid,
	}, nil
}

func decodeWebhookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webhookData := &entities.DocusignWebhookData{}
	err := decodeBodyFromRequest(webhookData, r)
	if err != nil {
		return nil, err
	}

	return webhookData, nil
}
