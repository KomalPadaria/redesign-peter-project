package http

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.GetCompanyInfoRequest | entities.UploadSecurityCampaignUsersRequest
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeGetCompanyInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &entities.GetCompanyInfoRequest{}
	err := decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	splitted := strings.Split(req.Token, "?")

	company_id := splitted[0]
	companyUuidStr, err := base64.StdEncoding.DecodeString(company_id)
	if err != nil {
		return nil, err
	}
	companyUuid := uuid.MustParse(string(companyUuidStr))

	return companyUuid, nil
}

func decodeUploadSecurityCampaignUsers(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	return &entities.UploadSecurityCampaignUsersRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
		File:        r.Body,
	}, nil
}

func decodeGetSecurityCampaignUsers(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	return &entities.GetCampaignUsersRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}, nil
}
