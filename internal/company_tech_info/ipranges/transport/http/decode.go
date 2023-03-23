// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.GetAllTechInfoIpRangeRequest | entities.TechInfoIpRanges | entities.UpdateTechInfoIpRangePatchRequestBody
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeGetAllTechInfoIpRangeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	req := &entities.GetAllTechInfoIpRangeRequest{
		CompanyUuid: compUUID,
		UserUuid:    userUUID,
	}

	return req, nil
}

func decodeCreateTechInfoIpRangeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	ipRange := &entities.TechInfoIpRanges{}
	err = decodeBodyFromRequest(ipRange, r)
	if err != nil {
		return nil, err
	}

	return &entities.CreateTechInfoIpRangeRequest{
		CompanyUuid:      compUUID,
		UserUuid:         userUUID,
		TechInfoIpRanges: ipRange,
	}, nil
}

func decodeUpdateTechInfoIpRangeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	ipRangeUUID, err := uuid.Parse(params["ip_range_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("ip_range_id")
	}

	ipRange := &entities.TechInfoIpRanges{}
	err = decodeBodyFromRequest(ipRange, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoIpRangeRequest{
		CompanyUuid:         compUUID,
		UserUuid:            userUUID,
		TechInfoIpRangeUuid: ipRangeUUID,
		TechInfoIpRanges:    ipRange,
	}, nil
}

func decodeUpdateTechInfoIpRangePatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	ipRangeUUID, err := uuid.Parse(params["ip_range_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("ip_range_id")
	}

	body := &entities.UpdateTechInfoIpRangePatchRequestBody{}
	err = decodeBodyFromRequest(body, r)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoIpRangePatchRequest{
		CompanyUuid:         compUUID,
		UserUuid:            userUUID,
		TechInfoIpRangeUuid: ipRangeUUID,
		PatchRequestBody:    body,
	}, nil
}

func decodeDeleteTechInfoIpRangeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	compUUID, err := uuid.Parse(params["company_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("company_id")
	}

	userUUID, err := uuid.Parse(params["user_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("user_id")
	}

	ipRangeUUID, err := uuid.Parse(params["ip_range_id"])
	if err != nil {
		return nil, httpError.NewErrBadOrInvalidPathParameter("ip_range_id")
	}

	return &entities.DeleteTechInfoIpRangeRequest{
		CompanyUuid:         compUUID,
		UserUuid:            userUUID,
		TechInfoIpRangeUuid: ipRangeUUID,
	}, nil
}
