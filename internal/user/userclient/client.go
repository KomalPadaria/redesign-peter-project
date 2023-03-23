package userclient

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/service"
)

type Client interface {
	CreateUser(ctx context.Context, user *entities.User) (uuid.UUID, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	GetUserByUuid(ctx context.Context, userUUID uuid.UUID) (*entities.User, error)
	GetContextUserCompanyInfoInternal(ctx context.Context) (*entities.GetUserCompanyInfoByUserIdResponse, error)
}

func NewClient(ep *endpoints.Endpoints, svc service.Service) Client {
	return &localClient{ep, svc}
}

type localClient struct {
	ep  *endpoints.Endpoints
	svc service.Service
}

func (l localClient) GetContextUserCompanyInfoInternal(ctx context.Context) (*entities.GetUserCompanyInfoByUserIdResponse, error) {
	return l.svc.GetContextUserCompanyInfoInternal(ctx)
}

func (l localClient) GetUserByUuid(ctx context.Context, userUUID uuid.UUID) (*entities.User, error) {
	return l.svc.GetUserByUuid(ctx, userUUID)
}

func (l localClient) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	res, err := l.ep.GetUserByUsernameEndpoint(ctx, username)

	user, _ := res.(*entities.User)

	return user, err
}

func (l localClient) CreateUser(ctx context.Context, req *entities.User) (uuid.UUID, error) {
	res, err := l.ep.CreateCompanyUserEndpoint(ctx, req)

	r, _ := res.(uuid.UUID) //nolint:errcheck

	return r, err
}
