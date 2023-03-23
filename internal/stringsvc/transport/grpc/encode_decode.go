// Package grpc contains transports
package grpc

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/pb"
)

func decodeUppercaseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.UppercaseRequest)
	req := &entities.UppercaseRequest{Input: request.GetInput()}

	return req, nil
}

func encodeUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*entities.UppercaseResponse)
	res := &pb.UppercaseResponse{Output: resp.Output}

	return res, nil
}

func decodeCountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.CountRequest)
	req := &entities.CountRequest{Input: request.GetInput()}

	return req, nil
}

func encodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*entities.CountResponse)
	res := &pb.CountResponse{Count: int32(resp.Count)}

	return res, nil
}
