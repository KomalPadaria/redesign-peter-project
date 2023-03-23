// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.TechInfoExternalInfra
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeCreateTechInfoExternalInfraRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	techInfoExternalInfra := &entities.TechInfoExternalInfra{}
	err = decodeBodyFromRequest(techInfoExternalInfra, r)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &entities.CreateTechInfoExternalInfraRequest{
		CompanyUuid:           compUUID,
		UserUuid:              userUUID,
		TechInfoExternalInfra: techInfoExternalInfra,
	}, nil
}

func decodeGetAllTechInfoExternalInfrasRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetAllTechInfoExternalInfrasRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeUpdateTechInfoExternalInfraRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	externalInfraUUID, err := uuid.Parse(params["external_infra_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("external_infra_id")
	}

	techInfoExternalInfra := &entities.TechInfoExternalInfra{}
	err = decodeBodyFromRequest(techInfoExternalInfra, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoExternalInfraRequest{
		CompanyUuid:               compUUID,
		UserUuid:                  userUUID,
		TechInfoExternalInfraUuid: externalInfraUUID,
		TechInfoExternalInfra:     techInfoExternalInfra,
	}, nil
}

func decodeDeleteTechInfoExternalInfraRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	externalInfraUUID, err := uuid.Parse(params["external_infra_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("external_infra_id")
	}

	return &entities.DeleteTechInfoExternalInfraRequest{
		CompanyUuid:               compUUID,
		UserUuid:                  userUUID,
		TechInfoExternalInfraUuid: externalInfraUUID,
	}, nil
}
