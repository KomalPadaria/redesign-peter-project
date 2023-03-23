// Package auth provides a way to interact with auth for different transports.
package auth

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/userclient"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModuleParams for auth.
type ModuleParams struct {
	fx.In
	CognitoClient cognito.Client
	UserClient    userclient.Client
	Logger        *zap.SugaredLogger
}

// NewModule for auth.
// nolint:gocritic
func NewModule(p ModuleParams) Client {
	mw := endpoint.New(p.CognitoClient, p.UserClient, p.Logger)

	return newClient(mw)
}

var (
	// Module for uber fx.
	Module = fx.Options(fx.Provide(NewModule))
)
