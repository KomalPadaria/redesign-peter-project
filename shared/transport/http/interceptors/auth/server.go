// Package auth contains interceptors for auth information
package auth

import (
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/auth"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/meta"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	v := r.Header.Get(string(auth.AuthorizationKey))
	if token, ok := auth.ExtractToken(v); ok {
		ctx = meta.WithRawToken(ctx, token)
	}

	idToken := r.Header.Get(string(auth.Access))
	ctx = meta.WithAccessToken(ctx, idToken)

	redesignToken := r.Header.Get(string(auth.RedesignTokenKey))
	if redesignToken != "" {
		ctx = meta.WithRedesignToken(ctx, redesignToken)
	}

	redesignWebhookToken := r.URL.Query().Get(string(auth.RedesignWebhookTokenKey))
	if redesignWebhookToken != "" {
		ctx = meta.WithRedesignWebhookToken(ctx, redesignWebhookToken)
	}

	h.next.ServeHTTP(w, r.WithContext(ctx))
}

// ServerHandler injects auth into context.
func ServerHandler(next http.Handler) http.Handler {
	return &authHandler{next}
}
