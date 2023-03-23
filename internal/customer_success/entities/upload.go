package entities

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UploadReportsRequest struct {
	CompanyUuid uuid.UUID
	ServiceName string
	Files       []*multipart.FileHeader
}

type DownloadReportRequest struct {
	CompanyUuid uuid.UUID
	ServiceName string
	ReportName  string
}
type DeleteReportRequest struct {
	CompanyUuid uuid.UUID
	ServiceName string
	ReportName  string
}

type UploadEvidencesRequest struct {
	CompanyUuid uuid.UUID
	ServiceUuid uuid.UUID
	Files       []*multipart.FileHeader
}

type DeleteEvidenceFileRequest struct {
	CompanyUuid  uuid.UUID
	ServiceUuid  uuid.UUID
	EvidenceUuid uuid.UUID
	FileUuid     uuid.UUID
}

type AddEvidenceFilesRequest struct {
	CompanyUuid  uuid.UUID
	ServiceUuid  uuid.UUID
	EvidenceUuid uuid.UUID
	Files        []*multipart.FileHeader
}

type DownloadEvidenceReportRequest struct {
	CompanyUuid  uuid.UUID
	ServiceUuid  uuid.UUID
	EvidenceUuid uuid.UUID
	ReportName   string
}
