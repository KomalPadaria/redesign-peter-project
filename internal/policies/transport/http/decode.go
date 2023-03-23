// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.Policy | entities.UpdatePolicyDocumentStatusPatchRequestBody | entities.SaveDocumentRequestBody
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeGetAllPoliciesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	keyword := r.URL.Query().Get("keyword")

	req := &entities.GetAllPoliciesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Keyword:     keyword,
	}

	return req, nil
}

func decodeCreatePolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policy := &entities.Policy{}
	err = decodeBodyFromRequest(policy, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreatePolicyRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Policy:      policy,
	}, nil
}

func decodeGetPolicyDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyUUID, err := uuid.Parse(params["policy_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("policy_id")
	}

	versionStr := r.URL.Query().Get("version")

	var version int
	if versionStr != "" {
		version, err = strconv.Atoi(versionStr)
		if err != nil {
			return nil, httpError.NewErrBadOrInvalidPathParameter("version")
		}
	}

	req := &entities.GetPolicyDocumentRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		PolicyUuid:  policyUUID,
		Version:     version,
	}

	return req, nil
}

func decodeSaveDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyUUID, err := uuid.Parse(params["policy_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("policy_id")
	}

	req := &entities.SaveDocumentRequestBody{}
	err = decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return &entities.SaveDocumentRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		PolicyUUID:              policyUUID,
		SaveDocumentRequestBody: req,
	}, nil
}

func decodeGetDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyUUID, err := uuid.Parse(params["policy_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("policy_id")
	}

	versionStr := r.URL.Query().Get("version")

	var version int
	if versionStr != "" {
		version, err = strconv.Atoi(versionStr)
		if err != nil {
			return nil, httpError.NewErrBadOrInvalidPathParameter("version")
		}
	}

	req := &entities.GetPolicyDocumentRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		PolicyUuid:  policyUUID,
		Version:     version,
	}

	return req, nil
}

func decodeGetPolicyHistoriesByPolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyUUID, err := uuid.Parse(params["policy_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("policy_id")
	}

	return &entities.GetPolicyHistoriesByPolicyRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		PolicyUUID:  policyUUID,
	}, nil
}

func decodeDeletePolicyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyUUID, err := uuid.Parse(params["policy_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("policy_id")
	}

	req := &entities.DeletePolicyRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		PolicyUuid:  policyUUID,
	}
	return req, nil
}

func decodeUpdatePolicyDocumentStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyUUID, err := uuid.Parse(params["policy_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("policy_id")
	}

	body := &entities.UpdatePolicyDocumentStatusPatchRequestBody{}

	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	req := &entities.UpdatePolicyDocumentStatus{
		CompanyUuid:      compUUID,
		UserUuid:         userUUID,
		PolicyUuid:       policyUUID,
		PatchRequestBody: body,
	}

	return req, nil
}

func decodeGetPoliciesStatsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetPoliciesStatsRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeGetTemplatesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	var companyTypes []string
	for k, v := range r.URL.Query() {
		if k == "company_type" {
			companyTypes = v
		}
	}

	req := &entities.GetTemplatesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		CompanyType: companyTypes,
	}

	return req, nil
}

func decodeCreateDocumentFromTemplateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	policyTemplateUuid, err := uuid.Parse(params["template_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("template_id")
	}

	req := &entities.CreateDocumentFromTemplateRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		PolicyTemplateUuid: policyTemplateUuid,
	}

	return req, nil
}
