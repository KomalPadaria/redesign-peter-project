package meta

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/entity"
)

type contextKey string

var contextKeyUsername = contextKey("cognitoUsername")

// User extracts user from context
func User(ctx context.Context) *entity.User {
	if user, ok := ctx.Value(contextKeyUsername).(*entity.User); ok {
		return user
	}

	return nil
}

// WithUser injects user metadata to the context
func WithUser(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, contextKeyUsername, user)
}
