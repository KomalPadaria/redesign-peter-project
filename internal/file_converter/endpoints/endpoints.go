package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/service"
)

type Endpoints struct {
	ConvertDocxFile endpoint.Endpoint
	ConvertHtmlFile endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		ConvertDocxFile: makeConvertDocx(svc),
		ConvertHtmlFile: makeConvertHtml(svc),
	}
}

func makeConvertDocx(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UploadedFile) //nolint:errcheck
		return svc.ConvertDocxToHtml(ctx, req)
	}
}

func makeConvertHtml(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UploadedFile) //nolint:errcheck
		return svc.ConvertHtmlToDocx(ctx, req)
	}
}
