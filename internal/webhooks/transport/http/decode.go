// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/webhooks/entities"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

type RequestBodyType interface {
	entities.AccountWebhookData | entities.AccountSubscriptionWebhookData
}

func decodeBodyFromRequest[T RequestBodyType](req *T, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.WithMessage(httpError.ErrBadRequestBody, err.Error())
	}

	defer r.Body.Close()

	return nil
}

func decodeSFAccountWebhookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &entities.AccountWebhookData{}
	err := decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeSFAccountSubscriptionsWebhookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &entities.AccountSubscriptionWebhookData{}
	err := decodeBodyFromRequest(req, r)
	if err != nil {
		return nil, err
	}

	return req, nil
}
