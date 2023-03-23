// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.CompanyWebsite
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeCreateWebsiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	website := &entities.CompanyWebsite{}
	err = decodeBodyFromRequest(website, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreateWebsiteRequest{
		CompanyUuid:    compUUID,
		UserUuid:       userUUID,
		CompanyWebsite: website,
	}, nil
}

func decodeGetAllWebsitesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetAllWebsitesRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeUpdateWebsiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	websiteUUID, err := uuid.Parse(params["website_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("website_id")
	}

	website := &entities.CompanyWebsite{}
	err = decodeBodyFromRequest(website, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateWebsiteRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		CompanyWebsiteUuid: websiteUUID,
		CompanyWebsite:     website,
	}, nil
}

func decodeDeleteWebsiteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	websiteUUID, err := uuid.Parse(params["website_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("website_id")
	}

	return &entities.DeleteWebsiteRequest{
		CompanyUuid:        compUUID,
		UserUuid:           userUUID,
		CompanyWebsiteUuid: websiteUUID,
	}, nil
}
