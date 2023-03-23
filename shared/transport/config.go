// Package transport supports transports like HTTP, gRPC and AMQP
package transport

import (
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/grpc"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
	"go.uber.org/fx"
)

// Config for transport.
type Config struct {
	fx.Out

	GRPC grpc.Config
	HTTP http.Config
}
