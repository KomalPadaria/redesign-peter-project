// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/entities"
)

type RequestBodyType interface {
	entities.TechInfoWireless | entities.UpdateTechInfoWirelessPatchRequestBody
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, "Bad request body")
	}

	defer r.Body.Close()

	return nil
}

func decodeCreateTechInfoWirelessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	techInfoWireless := &entities.TechInfoWireless{}
	err = decodeBodyFromRequest(techInfoWireless, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreateTechInfoWirelessRequest{
		CompanyUuid:      compUUID,
		UserUuid:         userUUID,
		TechInfoWireless: techInfoWireless,
	}, nil
}

func decodeGetAllTechInfoWirelesssRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetAllTechInfoWirelesssRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeUpdateTechInfoWirelessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	wirelessUUID, err := uuid.Parse(params["wireless_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("wireless_id")
	}

	techInfoWireless := &entities.TechInfoWireless{}
	err = decodeBodyFromRequest(techInfoWireless, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoWirelessRequest{
		CompanyUuid:          compUUID,
		UserUuid:             userUUID,
		TechInfoWirelessUuid: wirelessUUID,
		TechInfoWireless:     techInfoWireless,
	}, nil
}

func decodeDeleteTechInfoWirelessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	wirelessUUID, err := uuid.Parse(params["wireless_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("wireless_id")
	}

	return &entities.DeleteTechInfoWirelessRequest{
		CompanyUuid:          compUUID,
		UserUuid:             userUUID,
		TechInfoWirelessUuid: wirelessUUID,
	}, nil
}

func decodeUpdateTechInfoWirelessPatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	wirelessUUID, err := uuid.Parse(params["wireless_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("wireless_id")
	}

	body := &entities.UpdateTechInfoWirelessPatchRequestBody{}
	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoWirelessPatchRequest{
		CompanyUuid:          compUUID,
		UserUuid:             userUUID,
		TechInfoWirelessUuid: wirelessUUID,
		PatchRequestBody:     body,
	}, nil
}
