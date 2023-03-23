package calendly

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/service"
)

type Client interface {
	GetScheduledEvent(ctx context.Context, eventID string) (*entities.ScheduledEvent, error)
	GetUserInfo(ctx context.Context) (*entities.UserMe, error)
	GetEventTypes(ctx context.Context, userURI string) ([]entities.EventType, error)
	GetEventType(ctx context.Context, eventTypeUUID string) (*entities.EventType, error)
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l localClient) GetUserInfo(ctx context.Context) (*entities.UserMe, error) {
	return l.svc.GetUserInfo(ctx)
}

func (l localClient) GetEventTypes(ctx context.Context, userURI string) ([]entities.EventType, error) {
	return l.svc.GetEventTypes(ctx, userURI)
}

func (l localClient) GetEventType(ctx context.Context, eventTypeUUID string) (*entities.EventType, error) {
	return l.svc.GetEventType(ctx, eventTypeUUID)
}

func (l localClient) GetScheduledEvent(ctx context.Context, eventID string) (*entities.ScheduledEvent, error) {
	return l.svc.GetScheduledEvent(ctx, eventID)
}
