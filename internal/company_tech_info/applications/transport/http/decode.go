// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.TechInfoApplication | entities.UpdateApplicationPatchRequestBody |
		entities.ApplicationEnvCreateRequestBody | entities.ApplicationEnvUpdateRequestBody | entities.UpdateApplicationEnvPatchRequestBody
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeCreateApplicationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	application := &entities.TechInfoApplication{}
	err = decodeBodyFromRequest(application, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreateApplicationRequest{
		CompanyUuid:         compUUID,
		UserUuid:            userUUID,
		TechInfoApplication: application,
	}, nil
}

func decodeGetAllApplicationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

	req := &entities.GetAllApplicationsRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Keyword:     keyword,
	}

	return req, nil
}

func decodeUpdateApplicationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	applicationUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	application := &entities.TechInfoApplication{}
	err = decodeBodyFromRequest(application, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateApplicationRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: applicationUUID,
		TechInfoApplication:     application,
	}, nil
}

func decodeDeleteApplicationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	applicationUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	return &entities.DeleteApplicationRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: applicationUUID,
	}, nil
}

func decodeUpdateApplicationPatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	applicationUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	body := &entities.UpdateApplicationPatchRequestBody{}
	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateApplicationPatchRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: applicationUUID,
		PatchRequestBody:        body,
	}, nil
}

func decodeCreateApplicationEnvRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	appUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	env := &entities.ApplicationEnvCreateRequestBody{}
	err = decodeBodyFromRequest(env, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreateApplicationEnvRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: appUUID,
		ApplicationEnv:          env,
	}, nil
}

func decodeUpdateApplicationEnvRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	applicationUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	envUUID, err := uuid.Parse(params["env_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("env_id")
	}

	req := &entities.ApplicationEnvUpdateRequestBody{}
	err = decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateApplicationEnvRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: applicationUUID,
		ApplicationEnvUuid:      envUUID,
		ApplicationEnv:          req,
	}, nil
}

func decodeUpdateApplicationEnvPatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	applicationUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	envUUID, err := uuid.Parse(params["env_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("env_id")
	}

	req := &entities.UpdateApplicationEnvPatchRequestBody{}
	err = decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateApplicationEnvPatchRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: applicationUUID,
		ApplicationEnvUuid:      envUUID,
		PatchRequestBody:        req,
	}, nil
}

func decodeDeleteApplicationEnvRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	applicationUUID, err := uuid.Parse(params["application_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("application_id")
	}

	envUUID, err := uuid.Parse(params["env_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("env_id")
	}

	return &entities.DeleteApplicationEnvRequest{
		CompanyUuid:             compUUID,
		UserUuid:                userUUID,
		TechInfoApplicationUuid: applicationUUID,
		ApplicationEnvUuid:      envUUID,
	}, nil
}
