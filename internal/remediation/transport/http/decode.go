package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/remediation/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
)

type RequestBodyType interface {
	entities.ListTopRemediationRequest
}

func decodeListRemediationssRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	severity := r.URL.Query().Get("severity")

	topStr := r.URL.Query().Get("top")

	var top int
	if topStr != "" {
		top, err = strconv.Atoi(topStr)
		if err != nil {
			return nil, httpError.NewErrBadOrInvalidPathParameter("top")
		}
	}

	req := &entities.ListTopRemediationRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		Severity:    severity,
		Top:         top,
	}

	return req, nil
}
