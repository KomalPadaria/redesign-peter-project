package ses

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses/service"
)

type Client interface {
	SendEmail(ctx context.Context, subject, body, recipient string) error
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) SendEmail(ctx context.Context, subject, body, recipient string) error {
	return l.svc.SendEmail(ctx, subject, body, recipient)
}
