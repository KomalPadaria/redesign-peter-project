// Package grpc contains transports
package grpc

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/endpoints"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	uppercaseHandler grpctransport.Handler
	countHandler     grpctransport.Handler
	pb.UnimplementedStringAPIServer
}

// Register constructor
func Register(srv *grpc.Server, ep *endpoints.Endpoints) {
	pb.RegisterStringAPIServer(srv, MakeGRPCHandler(ep))
}

// MakeGRPCHandler constructor
func MakeGRPCHandler(ep *endpoints.Endpoints) pb.StringAPIServer {
	return &server{
		uppercaseHandler: grpctransport.NewServer(
			ep.UppercaseEndpoint,
			decodeUppercaseRequest,
			encodeUppercaseResponse,
		),
		countHandler: grpctransport.NewServer(
			ep.CountEndpoint,
			decodeCountRequest,
			encodeCountResponse,
		),
	}
}

func (svc *server) Uppercase(ctx context.Context, req *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	_, resp, err := svc.uppercaseHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, errorEncoder(err)
	}

	r := resp.(*pb.UppercaseResponse)

	return r, nil
}

func (svc *server) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	_, resp, err := svc.countHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, errorEncoder(err)
	}

	r := resp.(*pb.CountResponse)

	return r, nil
}

// errorEncoder returns error with appropriate GRPC status code
func errorEncoder(err error) error {
	statusErr, ok := status.FromError(err)
	if ok {
		return statusErr.Err()
	}

	return status.Error(codes.Internal, err.Error())
}
