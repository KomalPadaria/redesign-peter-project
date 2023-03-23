package salesforce

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/service"
)

type Client interface {
	NewSession(ctx context.Context) (*service.SFSession, error)
	GetAccountByID(ctx context.Context, session *service.SFSession, accountId string) (*entities.Account, error)
	GetContactByID(ctx context.Context, session *service.SFSession, contactId string) (*entities.Contact, error)
	GetSubscriptionsByID(ctx context.Context, session *service.SFSession, subscriptionsID string) (*entities.Subscription, error)
	GetSubscriptionsByIDs(ctx context.Context, session *service.SFSession, subscriptionsIDs []string) ([]entities.Subscription, error)
	GetSubscriptionsByAccountID(ctx context.Context, session *service.SFSession, accountId string) ([]entities.Subscription, error)
	GetContractsByAccountId(ctx context.Context, session *service.SFSession, accountId string) ([]entities.Contract, error)
	UpdateJiraEpicIdForAccount(ctx context.Context, session *service.SFSession, accountId, jiraEpicId string) (*entities.Account, error)
	GetConfig() config.Config
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l localClient) GetSubscriptionsByIDs(ctx context.Context, session *service.SFSession, subscriptionsIDs []string) ([]entities.Subscription, error) {
	return l.svc.GetSubscriptionsByIDs(ctx, session, subscriptionsIDs)
}

func (l localClient) UpdateJiraEpicIdForAccount(ctx context.Context, session *service.SFSession, accountId, jiraEpicId string) (*entities.Account, error) {
	return l.svc.UpdateJiraEpicIdForAccount(ctx, session, accountId, jiraEpicId)
}

func (l localClient) GetConfig() config.Config {
	return l.svc.GetConfig()
}

func (l localClient) GetSubscriptionsByID(ctx context.Context, session *service.SFSession, subscriptionsID string) (*entities.Subscription, error) {
	return l.svc.GetSubscriptionsByID(ctx, session, subscriptionsID)
}

func (l localClient) GetContractsByAccountId(ctx context.Context, session *service.SFSession, accountId string) ([]entities.Contract, error) {
	return l.svc.GetContractsByAccountId(ctx, session, accountId)
}

func (l localClient) GetSubscriptionsByAccountID(ctx context.Context, session *service.SFSession, accountId string) ([]entities.Subscription, error) {
	return l.svc.GetSubscriptionsByAccountID(ctx, session, accountId)
}

func (l localClient) GetAccountByID(ctx context.Context, session *service.SFSession, accountId string) (*entities.Account, error) {
	return l.svc.GetAccountByID(ctx, session, accountId)
}

func (l localClient) GetContactByID(ctx context.Context, session *service.SFSession, contactId string) (*entities.Contact, error) {
	return l.svc.GetContactByID(ctx, session, contactId)
}

func (l localClient) NewSession(ctx context.Context) (*service.SFSession, error) {
	return l.svc.NewSession(ctx)
}
