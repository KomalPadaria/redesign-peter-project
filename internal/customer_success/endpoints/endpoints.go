package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/service"
)

// Endpoints represents endpoints
type Endpoints struct {
	GetServiceReviewEndpoint          endpoint.Endpoint
	UpdateServiceReviewStatusEndpoint endpoint.Endpoint
	GetSubscriptionsEndpoint          endpoint.Endpoint
	GetSubscriptionPlansEndpoint      endpoint.Endpoint
	UploadReportsEndpoint             endpoint.Endpoint
	DownloadServiceReportEndpoint     endpoint.Endpoint
	DeleteServiceReportEndpoint       endpoint.Endpoint
	GetConsultingHoursEndpoint        endpoint.Endpoint
	GetConsumedHoursEndpoint          endpoint.Endpoint
	UploadEvidenceEndpoint            endpoint.Endpoint
	DeleteEvidenceFileEndpoint        endpoint.Endpoint
	AddEvidenceFilesEndpoint          endpoint.Endpoint
	DownloadEvidenceReportEndpoint    endpoint.Endpoint
}

// New returns new endpoints
func New(svc service.Service) *Endpoints {
	return &Endpoints{
		GetServiceReviewEndpoint:          makeGetServiceReviewEndpoint(svc),
		UpdateServiceReviewStatusEndpoint: makeUpdateServiceReviewStatusEndpoint(svc),
		GetSubscriptionsEndpoint:          makeGetSubscriptionsEndpoint(svc),
		GetSubscriptionPlansEndpoint:      makeGetSubscriptionPlansEndpoint(svc),
		UploadReportsEndpoint:             makeUploadReportsEndpoint(svc),
		DownloadServiceReportEndpoint:     makeDownloadServiceReportEndpoint(svc),
		DeleteServiceReportEndpoint:       makeDeleteServiceReportEndpoint(svc),
		GetConsultingHoursEndpoint:        makeGetConsultingHoursEndpoint(svc),
		GetConsumedHoursEndpoint:          makeGetConsumedHoursEndpoint(svc),
		UploadEvidenceEndpoint:            makeUploadEvidenceEndpoint(svc),
		DeleteEvidenceFileEndpoint:        makeDeleteEvidenceFileEndpoint(svc),
		AddEvidenceFilesEndpoint:          makeAddEvidenceFilesEndpoint(svc),
		DownloadEvidenceReportEndpoint:    makeDownloadEvidenceReportEndpoint(svc),
	}
}

func makeGetSubscriptionsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetSubscriptionsRequest) //nolint:errcheck

		return svc.GetSubscriptions(ctx, req)
	}
}

func makeGetSubscriptionPlansEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetSubscriptionPlansRequest) //nolint:errcheck

		return svc.GetSubscriptionPlans(ctx, req)
	}
}

func makeUploadReportsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UploadReportsRequest) //nolint:errcheck

		err := svc.UploadReports(ctx, req)

		return nil, err
	}
}

func makeDownloadServiceReportEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DownloadReportRequest) //nolint:errcheck

		return svc.DownloadReport(ctx, req)
	}
}

func makeDeleteServiceReportEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteReportRequest) //nolint:errcheck

		err := svc.DeleteReport(ctx, req)
		return nil, err
	}
}

func makeGetConsultingHoursEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetConsultingHoursRequest) //nolint:errcheck

		return svc.GetConsultingHours(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}
func makeGetConsumedHoursEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetConsumedHoursRequest) //nolint:errcheck

		return svc.GetConsumedHours(ctx, &req.CompanyUuid, &req.UserUuid)
	}
}

func makeGetServiceReviewEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.GetServiceReviewsRequest) //nolint:errcheck

		return svc.GetServiceReview(ctx, req)
	}
}
func makeUpdateServiceReviewStatusEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UpdateServiceReviewsStatusRequest) //nolint:errcheck

		return svc.UpdateServiceReviewStatus(ctx, req)
	}
}

func makeUploadEvidenceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.UploadEvidencesRequest) //nolint:errcheck

		err := svc.UploadServiceReport(ctx, req)

		return nil, err
	}
}

func makeDeleteEvidenceFileEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DeleteEvidenceFileRequest) //nolint:errcheck

		err := svc.DeleteEvidenceFile(ctx, req)

		return nil, err
	}
}

func makeAddEvidenceFilesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.AddEvidenceFilesRequest) //nolint:errcheck

		err := svc.AddEvidenceFiles(ctx, req)

		return nil, err
	}
}

func makeDownloadEvidenceReportEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entities.DownloadEvidenceReportRequest) //nolint:errcheck

		return svc.DownloadEvidenceReport(ctx, req)
	}
}
