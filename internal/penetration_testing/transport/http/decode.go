package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/penetration_testing/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
)

type RequestBodyType interface {
	entities.GetPenetrationTestsRequest | entities.GetPenetrationTestsStatsRequest
}

func decodeGetVulnerabilitiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetPenetrationTestsRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeGetVulnerabilityStatsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetPenetrationTestsStatsRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}
