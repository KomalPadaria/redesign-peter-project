package endpoints

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	CreateCompanyUserEndpoint                 endpoint.Endpoint
	GetContextUserCompanyInfoEndpoint         endpoint.Endpoint
	GetContextUserCompanyInfoInternalEndpoint endpoint.Endpoint
	GetSFUserAndCompanyInfoEndpoint           endpoint.Endpoint
	ActivateUserByEmailEndpoint               endpoint.Endpoint
	GetUserByUsernameEndpoint                 endpoint.Endpoint
	GetCompanyUsersEndpoint                   endpoint.Endpoint
	CreateUserEndpoint                        endpoint.Endpoint
	UpdateCompanyUserLinkEndpoint             endpoint.Endpoint
	DeleteCompanyUserLinkEndpoint             endpoint.Endpoint
	ResendInviteEndpoint                      endpoint.Endpoint
	SwitchCompanyEndpoint                     endpoint.Endpoint
	UpdateUserEndpoint                        endpoint.Endpoint
	ListCompaniesForUserEndpoint              endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateCompanyUserEndpoint:                 makeCreateCompanyUserEndpoint(svc),
		GetContextUserCompanyInfoEndpoint:         makeGetContextUserCompanyInfoEndpoint(svc),
		GetContextUserCompanyInfoInternalEndpoint: makeGetUserCompanyInfoByUsernameEndpoint(svc),
		GetSFUserAndCompanyInfoEndpoint:           makeGetSFUserAndCompanyInfoEndpoint(svc),
		ActivateUserByEmailEndpoint:               makeActivateUserByEmailIdEndpoint(svc),
		GetUserByUsernameEndpoint:                 makeGetUserByUsernameEndpoint(svc),
		GetCompanyUsersEndpoint:                   makeGetCompanyUsersEndpoint(svc),
		CreateUserEndpoint:                        makeCreateUserEndpoint(svc),
		UpdateCompanyUserLinkEndpoint:             makeUpdateCompanyUserLinkEndpoint(svc),
		DeleteCompanyUserLinkEndpoint:             makeDeleteCompanyUserLinkEndpoint(svc),
		ResendInviteEndpoint:                      makeResendUserInviteEndpoint(svc),
		SwitchCompanyEndpoint:                     makeSwitchCompanyEndpoint(svc),
		UpdateUserEndpoint:                        makeUpdateUserEndpoint(svc),
		ListCompaniesForUserEndpoint:              makeListCompaniesForUserEndpoint(svc),
	}
}

func makeCreateCompanyUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateCompanyAndUserRequest) //nolint:errcheck

		return svc.CreateCompanyUser(ctx, req)
	}
}

func makeGetContextUserCompanyInfoEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return svc.GetContextUserCompanyInfo(ctx)
	}
}

func makeGetUserCompanyInfoByUsernameEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return svc.GetContextUserCompanyInfoInternal(ctx)
	}
}

func makeGetSFUserAndCompanyInfoEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetSFUserAndCompanyInfoRequest) //nolint:errcheck

		return svc.GetSFUserAndCompanyInfo(ctx, req)
	}
}

func makeActivateUserByEmailIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.ActivateDeactivateUserRequest) //nolint:errcheck

		return svc.ActivateUserByEmail(ctx, req)
	}
}

func makeGetUserByUsernameEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return svc.GetUserByUsername(ctx, fmt.Sprintf("%v", req))
	}
}

func makeGetCompanyUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompanyUsersRequest) //nolint:errcheck

		return svc.GetCompanyUsers(ctx, req.CompanyUUID, req.UserUUID)
	}
}

func makeCreateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateUserRequest) //nolint:errcheck

		return svc.CreateUser(ctx, req.CompanyUUID, req.UserUUID, req.CreateUserRequestBody)
	}
}

func makeUpdateCompanyUserLinkEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateCompanyUserLinkRequest) //nolint:errcheck

		req.CompanyUser.UpdatedBy = req.ReqUserUuid
		req.CompanyUser.CompanyUuid = req.CompanyUuid
		req.CompanyUser.UserUuid = req.UserUuid

		return svc.UpdateCompanyUserLink(ctx, req.CompanyUser)
	}
}

func makeDeleteCompanyUserLinkEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteCompanyUserLinkRequest) //nolint:errcheck

		err := svc.DeleteCompanyUserLink(ctx, &req.CompanyUuid, &req.ReqUserUuid, &req.UserUuid)

		return nil, err
	}
}

func makeResendUserInviteEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.CreateUserRequest) //nolint:errcheck

		err := svc.ResendUserInvite(ctx, req.CompanyUUID, req.UserUUID, req.CreateUserRequestBody)
		return nil, err
	}
}

func makeSwitchCompanyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateCurrentCompanyRequest) //nolint:errcheck

		err := svc.UpdateCurrentCompany(ctx, req)
		return nil, err
	}
}

func makeUpdateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateUserRequest) //nolint:errcheck

		return svc.UpdateUserDetails(ctx, req)
	}
}

func makeListCompaniesForUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetCompaniesRequest) //nolint:errcheck

		return svc.ListCompaniesForUser(ctx, req.UserUuid, req.Keyword)
	}
}
