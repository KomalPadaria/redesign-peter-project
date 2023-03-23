// Package auth provides a way to interact with auth for different transports.
package auth

import (
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/endpoint"
)

// Client exposed by auth.
type Client interface {
	SecureServiceWithCognitoEndpoint(ept goKitEndpoint.Endpoint, featureName, method string) goKitEndpoint.Endpoint
	SecureServiceWithRedesignEndpoint(ept goKitEndpoint.Endpoint) goKitEndpoint.Endpoint
	SecureServiceWithRedesignWebhookEndpoint(ept goKitEndpoint.Endpoint) goKitEndpoint.Endpoint
}

func newClient(mw endpoint.Middleware) Client {
	return &client{mw}
}

type client struct {
	mw endpoint.Middleware
}

func (c *client) SecureServiceWithRedesignWebhookEndpoint(ept goKitEndpoint.Endpoint) goKitEndpoint.Endpoint {
	return goKitEndpoint.Chain(c.mw.SecureServiceWithRedesignWebhookEndpoint())(ept)
}

func (c *client) SecureServiceWithRedesignEndpoint(ept goKitEndpoint.Endpoint) goKitEndpoint.Endpoint {
	return goKitEndpoint.Chain(c.mw.SecureServiceWithRedesignEndpoint())(ept)

}

// SecureServiceWithCognitoEndpoint wraps endpoint with middleware to retrieve user info and checks only keycloak token
func (c *client) SecureServiceWithCognitoEndpoint(ept goKitEndpoint.Endpoint, featureName, method string) goKitEndpoint.Endpoint {
	return goKitEndpoint.Chain(c.mw.SecureServiceWithCognitoEndpoint(featureName, method))(ept)
}
