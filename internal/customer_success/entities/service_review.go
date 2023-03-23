package entities

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

const (
	ServiceReportsStatusTypeAcknowledged            = "ACKNOWLEDGED"
	ServiceReportsStatusTypeRequiredAcknowledgement = "REQUIRED_ACKNOWLEDGEMENT"

	ChangelogStatus = "Resolved"
)

type UpdateServiceReviewsStatusRequestBody struct {
	Status string `json:"status"`
}
type UpdateServiceReviewsStatusRequest struct {
	CompanyUuid      uuid.UUID                              `json:"company_uuid"`
	UserUuid         uuid.UUID                              `json:"user_uuid"`
	EvidenceUuid     uuid.UUID                              `json:"evidence_uuid"`
	PatchRequestBody *UpdateServiceReviewsStatusRequestBody `json:"patch_request_body"`
}

type GetServiceReviewsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetServiceReviewResponse struct {
	ServiceReviewUuid uuid.UUID  `json:"service_review_uuid"`
	ServiceName       string     `json:"service_name"`
	Status            string     `json:"status"`
	SubscriptionType  string     `json:"subscription_type"`
	Evidence          []Evidence `json:"evidences"`
}

func (t *GetServiceReviewResponse) MarshalJSON() ([]byte, error) {
	type Alias GetServiceReviewResponse

	o := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	for _, e := range t.Evidence {
		if e.Status == ServiceReportsStatusTypeRequiredAcknowledgement {
			o.Status = "ACTION_REQUIRED"
			break
		}
		o.Status = ServiceReportsStatusTypeAcknowledged
	}

	if len(t.Evidence) == 0 {
		o.Status = "ACTION_REQUIRED"
	}

	return json.Marshal(o)
}

type Evidence struct {
	EvidenceId     uuid.UUID         `json:"evidence_uuid "`
	CompletedOn    nullable.NullTime `json:"completed_on"`
	AcknowledgedAt nullable.NullTime `json:"acknowledged_at"`
	AcknowledgedBy string            `json:"acknowledged_by"`
	Data           []FileData        `json:"data"`
	Status         string            `json:"status"`
}

type FileData struct {
	FileId   uuid.UUID `json:"file_uuid"`
	FileName string    `json:"file_name"`
	FilePath string    `json:"file_path"`
}
