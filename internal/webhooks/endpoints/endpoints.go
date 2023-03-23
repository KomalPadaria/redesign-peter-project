package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/service"
)

type Endpoints struct {
	SFAccountWebhookEndpoint              endpoint.Endpoint
	SFAccountSubscriptionsWebhookEndpoint endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		SFAccountWebhookEndpoint:              makeSFAccountWebhookEndpoint(svc),
		SFAccountSubscriptionsWebhookEndpoint: makeSFAccountSubscriptionsWebhookEndpoint(svc),
	}
}

func makeSFAccountWebhookEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.AccountWebhookData) //nolint:errcheck
		err := svc.SFAccountWebhook(ctx, req)
		return nil, err
	}
}

func makeSFAccountSubscriptionsWebhookEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.AccountSubscriptionWebhookData) //nolint:errcheck
		err := svc.SFAccountSubscriptionsWebhook(ctx, *req)
		return nil, err
	}
}
