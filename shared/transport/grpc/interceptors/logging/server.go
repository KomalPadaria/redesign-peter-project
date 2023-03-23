// Package logging contains logs interceptors
package logging

import (
	"context"
	"path"
	"time"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/log"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/meta"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/grpc/interceptors/paths"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds logger to the context.
//
//nolint:funlen
func UnaryServerInterceptor(logger *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		service := path.Dir(info.FullMethod)[1:]
		method := path.Base(info.FullMethod)

		// do not log health checks
		if service == paths.HealthCheck {
			return handler(ctx, req)
		}

		ctxLogger := logger.With(
			"component", "server",
			"grpc.service", service,
			"grpc.method", method,
			"request_id", meta.RequestID(ctx),
			"user_agent", meta.UserAgent(ctx),
			"user_agent_origin", meta.UserAgentOrigin(ctx),
		)

		reqFields := []interface{}{}

		reqPayload, err := encodeGRPCPayload(req)
		if err != nil {
			reqFields = append(reqFields, "grpc.request.parse_error", err)
		}

		reqFields = append(reqFields, "grpc.request.content", log.Mask(reqPayload))

		ctxLogger.Infow("started unary call", reqFields...)

		// invoke rpc
		resp, rpcErr := handler(ctx, req)

		code := status.Code(rpcErr)
		logFunc := codeToLevel(code, ctxLogger)

		resFields := []interface{}{}
		resFields = append(resFields, "grpc.time_ms", durationToMilliseconds(time.Since(startTime)))

		resPayload, err := encodeGRPCPayload(resp)
		if err != nil {
			resFields = append(resFields, "grpc.response.parse_error", err)
		}

		resFields = append(resFields, "grpc.response.content", log.Mask(resPayload))

		logFunc("finished unary call with code "+code.String(), append(resFields, "error", rpcErr, "grpc.code", code.String())...)

		return resp, rpcErr
	}
}

// StreamServerInterceptor returns a new streaming server interceptors that adds logger to the context.
//
//nolint:funlen
func StreamServerInterceptor(logger *zap.SugaredLogger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()

		ctx := stream.Context()
		service := path.Dir(info.FullMethod)[1:]
		method := path.Base(info.FullMethod)

		// do not log health checks
		if service == paths.HealthCheck {
			return handler(srv, stream)
		}

		ctxLogger := logger.With(
			"component", "server",
			"grpc.service", service,
			"grpc.method", method,
			"request_id", meta.RequestID(ctx),
			"user_agent", meta.UserAgent(ctx),
			"user_agent_origin", meta.UserAgentOrigin(ctx),
		)

		ctxLogger.Infow("started stream call")

		// invoke stream
		streamErr := handler(srv, stream)

		code := status.Code(streamErr)
		logFunc := codeToLevel(code, ctxLogger)

		resFields := []interface{}{}
		resFields = append(resFields, "grpc.time_ms", durationToMilliseconds(time.Since(startTime)))

		logFunc("finished stream call with code "+code.String(), append(resFields, "error", streamErr, "grpc.code", code.String())...)

		return streamErr
	}
}
