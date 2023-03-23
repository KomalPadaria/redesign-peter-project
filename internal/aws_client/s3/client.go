package s3client

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/service"
)

type Client interface {
	UploadFilesFromFileHeaders(ctx context.Context, batchUploadObjects []entities.BatchUploadObject) error
	DownloadFile(ctx context.Context, key string) ([]byte, error)
	DeleteFile(ctx context.Context, key string) error
}

func NewClient(svc service.Service) Client {
	return &localClient{svc}
}

type localClient struct {
	svc service.Service
}

func (l *localClient) UploadFilesFromFileHeaders(ctx context.Context, batchUploadObjects []entities.BatchUploadObject) error {
	return l.svc.UploadFilesFromFileHeaders(ctx, batchUploadObjects)
}

func (l *localClient) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	return l.svc.DownloadFile(ctx, key)
}

func (l *localClient) DeleteFile(ctx context.Context, key string) error {
	return l.svc.DeleteFile(ctx, key)
}
