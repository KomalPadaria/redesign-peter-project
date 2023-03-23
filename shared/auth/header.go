// Package auth provides a way to interact with auth for different transports.
package auth

import (
	"strings"
)

type headerKey string
type queryPathKey string

const (
	// AuthorizationKey for bearer tokens.
	AuthorizationKey headerKey = headerKey("Authorization")
	Access           headerKey = headerKey("Access")
	RedesignTokenKey headerKey = headerKey("Redesign-Access-Token")

	RedesignWebhookTokenKey queryPathKey = queryPathKey("token")
)

// ExtractToken from a bearer token.
func ExtractToken(val string) (token string, ok bool) {
	parts := strings.Split(val, " ")
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", false
	}

	return parts[1], true
}
