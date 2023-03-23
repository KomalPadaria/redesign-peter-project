package cognito

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/service"
)

type Client interface {
	VerifyToken(ctx context.Context, token string) (*entities.AWSCognitoIDTokenClaims, error)
	GetClaimsFromIDToken(ctx context.Context, idToken string) (*entities.AWSCognitoIDTokenClaims, error)
	InviteUser(ctx context.Context, details entities.InviteUserDetails, resend bool) error
	UpdateUserAttributes(ctx context.Context, details map[string]string) error
}

func newClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) GetClaimsFromIDToken(ctx context.Context, idToken string) (*entities.AWSCognitoIDTokenClaims, error) {
	return l.svc.GetClaimsFromIDToken(ctx, idToken)
}

func (l *localClient) VerifyToken(ctx context.Context, token string) (*entities.AWSCognitoIDTokenClaims, error) {
	return l.svc.VerifyToken(ctx, token)
}

func (l *localClient) InviteUser(ctx context.Context, details entities.InviteUserDetails, resend bool) error {
	return l.svc.InviteUser(ctx, details, resend)
}

func (l *localClient) UpdateUserAttributes(ctx context.Context, details map[string]string) error {
	return l.svc.UpdateUserAttributes(ctx, details)
}
