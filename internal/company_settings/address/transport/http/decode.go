// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.CompanyAddress | entities.UpdateCompanyAddressPatchRequestBody
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeCreateCompanyAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	address := &entities.CompanyAddress{}
	err = decodeBodyFromRequest(address, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreateCompanyAddressRequest{
		CompanyUuid:    compUUID,
		UserUuid:       userUUID,
		CompanyAddress: address,
	}, nil
}

func decodeUpdateCompanyAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	addressUUID, err := uuid.Parse(params["address_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("address_id")
	}

	address := &entities.CompanyAddress{}
	err = decodeBodyFromRequest(address, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateCompanyAddressRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		CompanyAddressUuid: addressUUID,
		CompanyAddress:     address,
	}, nil
}

func decodeGetCompanyAddressesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetCompanyAddressRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeDeleteCompanyAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	addressUUID, err := uuid.Parse(params["address_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("address_id")
	}

	return &entities.DeleteCompanyAddressRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		CompanyAddressUuid: addressUUID,
	}, nil
}

func decodeUpdateCompanyAddressPatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	addressUUID, err := uuid.Parse(params["address_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("address_id")
	}

	body := &entities.UpdateCompanyAddressPatchRequestBody{}
	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateCompanyAddressPatchRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		CompanyAddressUuid: addressUUID,
		PatchRequestBody:   body,
	}, nil
}

func decodeGetCompanyFacilitiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	query := r.URL.Query().Get("query")
	status := r.URL.Query().Get("status")

	req := &entities.GetCompanyFacilitiesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Query:       query,
		Status:      status,
	}

	return req, nil
}
