package stringsvc

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/entities"
)

type Client interface {
	Uppercase(ctx context.Context, req *entities.UppercaseRequest) (*entities.UppercaseResponse, error)
	Count(ctx context.Context, req *entities.CountRequest) (*entities.CountResponse, error)
}

func NewClient(ep *endpoints.Endpoints) Client {
	return &localClient{ep}
}

type localClient struct {
	ep *endpoints.Endpoints
}

func (l localClient) Uppercase(ctx context.Context, req *entities.UppercaseRequest) (*entities.UppercaseResponse, error) {
	res, err := l.ep.UppercaseEndpoint(ctx, req)

	r, _ := res.(*entities.UppercaseResponse) //nolint:errcheck

	return r, err
}

func (l localClient) Count(ctx context.Context, req *entities.CountRequest) (*entities.CountResponse, error) {
	res, err := l.ep.CountEndpoint(ctx, req)

	r, _ := res.(*entities.CountResponse) //nolint:errcheck

	return r, err
}
