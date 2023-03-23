// Package meta contains metadata about request
package meta

import (
	"context"
)

type contextKey string

var (
	contextKeyRequestID            = contextKey("request_id")
	contextKeyUserAgent            = contextKey("user_agent")
	contextKeyUserAgentOrigin      = contextKey("user_agent_origin")
	contextKeyTransport            = contextKey("transport")
	contextKeyRawToken             = contextKey("raw_token")
	contextKeyAccessToken          = contextKey("access_token")
	contextKeyRedesignToken        = contextKey("redesign_token")
	contextKeyRedesignWebhookToken = contextKey("redesign_webhook_token")
	contextKeyClientID             = contextKey("client_id")
	contextKeyAPIKey               = contextKey("api_key")
	contextKeyClaims               = contextKey("claims")
	contextKeyUserGroups           = contextKey("user_groups")
)

func (c contextKey) String() string { return string(c) }

// RequestID extracts request id from the context
func RequestID(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyRequestID).(string); ok {
		return val
	}

	return ""
}

// WithRequestID injects request id metadata to the context
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextKeyRequestID, id)
}

// UserAgent extracts user agent from the context
func UserAgent(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyUserAgent).(string); ok {
		return val
	}

	return ""
}

// WithUserAgent injects user agent metadata to the context
func WithUserAgent(ctx context.Context, ua string) context.Context {
	return context.WithValue(ctx, contextKeyUserAgent, ua)
}

// UserAgentOrigin extracts user agent origin from the context
func UserAgentOrigin(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyUserAgentOrigin).(string); ok {
		return val
	}

	return ""
}

// WithUserAgentOrigin injects user agent origin metadata to the context
func WithUserAgentOrigin(ctx context.Context, ua string) context.Context {
	return context.WithValue(ctx, contextKeyUserAgentOrigin, ua)
}

// Transport extracts transport from the context
func Transport(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyTransport).(string); ok {
		return val
	}

	return ""
}

// WithTransport injects transport metadata to the context
func WithTransport(ctx context.Context, transport string) context.Context {
	return context.WithValue(ctx, contextKeyTransport, transport)
}

// RawToken extracts token from the context.
func RawToken(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyRawToken).(string); ok {
		return val
	}

	return ""
}

// WithRawToken injects token metadata to the context
func WithRawToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKeyRawToken, token)
}

// AccessToken extracts access token from the context.
func AccessToken(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyAccessToken).(string); ok {
		return val
	}

	return ""
}

// WithIdToken injects id token metadata to the context
func WithAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKeyAccessToken, token)
}

// RedesignToken extracts token from the context.
func RedesignToken(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyRedesignToken).(string); ok {
		return val
	}

	return ""
}

// WithRedesignToken injects token metadata to the context
func WithRedesignToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKeyRedesignToken, token)
}

// RedesignWebhookToken extracts token from the context.
func RedesignWebhookToken(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyRedesignWebhookToken).(string); ok {
		return val
	}

	return ""
}

// WithRedesignWebhookToken injects token metadata to the context
func WithRedesignWebhookToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKeyRedesignWebhookToken, token)
}

// ClientID extracts ClientId from the context.
func ClientID(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyClientID).(string); ok {
		return val
	}

	return ""
}

// WithClientID injects ClientId metadata to the context
func WithClientID(ctx context.Context, clientID string) context.Context {
	return context.WithValue(ctx, contextKeyClientID, clientID)
}

// APIKey extracts api key from the context.
func APIKey(ctx context.Context) string {
	if val, ok := ctx.Value(contextKeyAPIKey).(string); ok {
		return val
	}

	return ""
}

// WithAPIKey injects api key metadata to the context
func WithAPIKey(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKeyAPIKey, token)
}

// Claims extracts claims from the context.
func Claims(ctx context.Context) map[string]interface{} {
	if val, ok := ctx.Value(contextKeyClaims).(map[string]interface{}); ok {
		return val
	}

	return map[string]interface{}{}
}

// WithClaims injects claims metadata to the context
func WithClaims(ctx context.Context, claims map[string]interface{}) context.Context {
	return context.WithValue(ctx, contextKeyClaims, claims)
}

// UserGroups extracts user groups from the context.
func UserGroups(ctx context.Context) []string {
	if val, ok := ctx.Value(contextKeyUserGroups).([]string); ok {
		return val
	}

	return []string{}
}

// WithUserGroups injects user groups metadata to the context
func WithUserGroups(ctx context.Context, groups []string) context.Context {
	return context.WithValue(ctx, contextKeyUserGroups, groups)
}

// UserID extracts user ID from the context
func UserID(ctx context.Context) string {
	cl := Claims(ctx)

	if val, ok := cl["sub"].(string); ok {
		return val
	}

	return ""
}

// UserEmail extracts user email from the context
func UserEmail(ctx context.Context) string {
	cl := Claims(ctx)

	if val, ok := cl["email"].(string); ok {
		return val
	}

	return ""
}
