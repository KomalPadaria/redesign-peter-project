// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.CreateCompanyAndUserRequest | entities.GetSFUserAndCompanyInfoRequest | entities.ActivateDeactivateUserRequest | entities.CreateUserRequestBody | entities.CompanyUser | entities.ConfirmUserRequest | entities.UpdateUserRequestBody
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeCreateCompanyAndUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &entities.CreateCompanyAndUserRequest{}
	err := decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetContextUserCompanyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeGetSFUserAndCompanyInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &entities.GetSFUserAndCompanyInfoRequest{}
	err := decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeActivateUserByExternalIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &entities.ActivateDeactivateUserRequest{}
	err := decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetCompanyUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetCompanyUsersRequest{
		CompanyUUID: compUUID,
		UserUUID:    userUUID,
	}

	return req, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	reqBody := &entities.CreateUserRequestBody{}
	err = decodeBodyFromRequest(reqBody, r)
	if err != nil {
		return nil, err
	}

	req := &entities.CreateUserRequest{
		CompanyUUID:           compUUID,
		UserUUID:              userUUID,
		CreateUserRequestBody: reqBody,
	}

	return req, nil
}

func decodeUpdateCompanyUserLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	reqUserUUID, err := uuid.Parse(params["req_user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("req_user_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	reqBody := &entities.CompanyUser{}
	err = decodeBodyFromRequest(reqBody, r)
	if err != nil {
		return nil, err
	}

	req := &entities.UpdateCompanyUserLinkRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		ReqUserUuid: reqUserUUID,
		CompanyUser: reqBody,
	}

	return req, nil
}

func decodeDecodeCompanyUserLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	reqUserUUID, err := uuid.Parse(params["req_user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("req_user_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.DeleteCompanyUserLinkRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		ReqUserUuid: reqUserUUID,
	}

	return req, nil
}

func decodeResendUserInviteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	reqBody := &entities.CreateUserRequestBody{}
	err = decodeBodyFromRequest(reqBody, r)
	if err != nil {
		return nil, err
	}

	req := &entities.CreateUserRequest{
		CompanyUUID:           compUUID,
		UserUUID:              userUUID,
		CreateUserRequestBody: reqBody,
	}

	return req, nil
}

func decodeSwitchCompanyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	req := &entities.UpdateCurrentCompanyRequest{
		UserUUID:    userUUID,
		CompanyUUID: compUUID,
	}

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	reqBody := &entities.UpdateUserRequestBody{}
	err := decodeBodyFromRequest(reqBody, r)
	if err != nil {
		return nil, err
	}
	var mfa_methods []string

	if reqBody.IsMfaAppEnabled {
		mfa_methods = append(mfa_methods, "app")
	}
	if reqBody.IsMfaSmsEnabled {
		mfa_methods = append(mfa_methods, "sms")
	}

	req := &entities.UpdateUserRequest{
		FirstName:  reqBody.FirstName,
		LastName:   reqBody.LastName,
		Phone:      reqBody.Phone,
		JobTitle:   reqBody.JobTitle,
		MfaMethods: mfa_methods,
	}

	return req, nil
}

func decodeListCompaniesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	keyword := r.URL.Query().Get("keyword")

	req := &entities.GetCompaniesRequest{
		UserUuid: userUUID,
		Keyword:  keyword,
	}

	return req, nil
}
