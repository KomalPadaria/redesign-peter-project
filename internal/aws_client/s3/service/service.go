package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/entities"
)

type Service interface {
	UploadFilesFromFileHeaders(ctx context.Context, batchUploadObjects []entities.BatchUploadObject) error
	DownloadFile(ctx context.Context, key string) ([]byte, error)
	DeleteFile(ctx context.Context, key string) error
}

func New(config config.Config) Service {

	// The session the S3 Uploader will use
	sess := aws_client.GetSession(config.AwsRegion, 3)

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	return &service{config, uploader, downloader}
}

type service struct {
	config     config.Config
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func (s service) UploadFilesFromFileHeaders(ctx context.Context, batchUploadObjects []entities.BatchUploadObject) error {
	objects := make([]s3manager.BatchUploadObject, 0)

	for _, v := range batchUploadObjects {
		file, err := v.FileHeader.Open()
		if err != nil {
			return err
		}
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		objects = append(objects, s3manager.BatchUploadObject{
			Object: &s3manager.UploadInput{
				ACL:                aws.String("private"),
				Body:               bytes.NewReader(fileBytes),
				Bucket:             aws.String(s.config.BucketName),
				Key:                aws.String(v.Key),
				ContentType:        aws.String(http.DetectContentType(fileBytes)),
				ContentDisposition: aws.String(fmt.Sprintf("attachment; filename=\"%s\"", v.FileName)),
			},
		})
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	iter := &s3manager.UploadObjectsIterator{Objects: objects}
	if err := s.uploader.UploadWithIterator(ctx, iter); err != nil {
		return err
	}

	return nil
}

func (s service) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})

	n, err := s.downloader.DownloadWithContext(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(s.config.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}

	if n <= 0 {
		return nil, fmt.Errorf("no bytes found from s3")
	}

	return buf.Bytes(), nil
}

func (s service) DeleteFile(ctx context.Context, key string) error {
	svc := s3.New(aws_client.GetSession(s.config.AwsRegion, 3))

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.config.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file, %v", err)
	}

	return nil
}
