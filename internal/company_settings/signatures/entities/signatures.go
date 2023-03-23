package entities

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type Signatures struct {
	SignatureUuid uuid.UUID         `json:"signature_uuid,omitempty"`
	Name          string            `json:"name,omitempty"`
	DocumentUrl   string            `json:"document_url"`
	CompanyTypes  *CompanyTypes     `json:"company_types,omitempty"`
	CreatedAt     nullable.NullTime `json:"created_at,omitempty"`
	UpdatedAt     nullable.NullTime `json:"updated_at,omitempty"`
	CreatedBy     uuid.UUID         `json:"created_by,omitempty"`
	UpdatedBy     uuid.UUID         `json:"updated_by,omitempty"`
}

type CompanySignatures struct {
	CompanySignatureUuid uuid.UUID           `json:"company_signature_uuid" gorm:"column:company_signature_uuid"`
	SignatureUuid        uuid.UUID           `json:"signature_uuid" gorm:"column:signature_uuid"`
	CompanyUuid          uuid.UUID           `json:"company_uuid" gorm:"column:company_uuid"`
	Name                 string              `json:"name" gorm:"column:name"`
	Status               string              `json:"status" gorm:"column:status"`
	SignatureData        DocusignWebhookData `json:"signature_data" gorm:"column:signature_data"`
	CreatedAt            nullable.NullTime   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt            nullable.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy            uuid.UUID           `json:"created_by" gorm:"column:created_by"`
	UpdatedBy            uuid.UUID           `json:"updated_by" gorm:"column:updated_by"`
}

func (m *CompanySignatures) TableName() string {
	return "company_signatures"
}

type GetSignaturesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetSignaturesResponse struct {
	CompanySignatureUuid uuid.UUID         `json:"company_signature_uuid,omitempty"`
	SignatureUuid        uuid.UUID         `json:"signature_uuid,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Date                 nullable.NullTime `json:"date,omitempty"`
	Status               sql.NullString    `json:"status,omitempty"`
	DocumentUrl          string            `json:"document_url"`
	DocumentViewURI      string            `json:"document_view_uri"`
}

func (c *GetSignaturesResponse) MarshalJSON() ([]byte, error) {
	type Alias GetSignaturesResponse

	o := &struct {
		*Alias
		Date                 string `json:"date"`
		Status               string `json:"status"`
		CompanySignatureUuid string `json:"company_signature_uuid,omitempty"`
	}{
		Alias: (*Alias)(c),
	}

	if !c.Date.Valid {
		o.Date = "Not Signed"
	} else {
		o.Date = c.Date.Time.Format(time.RFC3339)
	}

	if !c.Status.Valid {
		o.Status = "Not Signed"
	} else {
		o.Status = c.Status.String
	}

	if c.CompanySignatureUuid == uuid.Nil {
		o.CompanySignatureUuid = ""
	} else {
		o.CompanySignatureUuid = c.CompanySignatureUuid.String()
	}

	return json.Marshal(o)
}

type UpdateStatusRequestBody struct {
	Status string `json:"status"`
}

type UpdateStatusRequest struct {
	CompanyUuid          uuid.UUID                `json:"company_uuid"`
	UserUuid             uuid.UUID                `json:"user_uuid"`
	CompanySignatureUuid uuid.UUID                `json:"company_signature_uuid"`
	PatchRequestBody     *UpdateStatusRequestBody `json:"patch_request_body"`
}

type UpdateStatusResponse struct {
	CompanySignatureUuid uuid.UUID `json:"company_signature_uuid"`
	Status               string    `json:"status,omitempty"`
}

type ViewDocumentRequest struct {
	CompanyUuid   uuid.UUID `json:"company_uuid"`
	SignatureUuid uuid.UUID `json:"signature_uuid"`
}
