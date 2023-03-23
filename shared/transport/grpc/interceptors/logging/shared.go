// Package logging contains logs interceptors
package logging

import (
	"reflect"
	"time"

	"github.com/golang/protobuf/jsonpb" // nolint:staticcheck
	"github.com/golang/protobuf/proto"  // nolint:staticcheck

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

func codeToLevel(code codes.Code, logger *zap.SugaredLogger) func(msg string, keysAndValues ...interface{}) {
	switch code {
	case codes.OK, codes.Canceled, codes.InvalidArgument, codes.NotFound, codes.AlreadyExists, codes.Unauthenticated:
		return logger.Infow
	case codes.DeadlineExceeded, codes.PermissionDenied, codes.ResourceExhausted, codes.FailedPrecondition, codes.Aborted, codes.OutOfRange, codes.Unavailable:
		return logger.Warnw
	case codes.Unknown, codes.Unimplemented, codes.Internal, codes.DataLoss:
		return logger.Errorw
	default:
		return logger.Errorw
	}
}

func durationToMilliseconds(duration time.Duration) float32 {
	return float32(duration.Nanoseconds()/1000) / 1000
}

func encodeGRPCPayload(payload interface{}) (string, error) {
	p, isProto := payload.(proto.Message)
	v := reflect.ValueOf(p)

	if isProto && p != nil && (v.Kind() != reflect.Ptr || (v.Kind() == reflect.Ptr && !v.IsNil())) {
		m := &jsonpb.Marshaler{}
		return m.MarshalToString(p)
	}

	return "", nil
}
