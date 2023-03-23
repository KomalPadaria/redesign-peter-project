package entities

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type Policy struct {
	PolicyUuid         uuid.UUID         `json:"policy_uuid" gorm:"column:policy_uuid"`
	CompanyUuid        uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	PolicyTemplateUuid nullable.NullUUID `json:"policy_template_uuid" gorm:"column:policy_template_uuid"`
	Name               string            `json:"name" gorm:"column:name"`
	Status             string            `json:"status" gorm:"column:status"`
	CreatedAt          nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy          uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy          nullable.NullUUID `json:"updated_by" gorm:"column:updated_by"`
	StatusUpdatedAt    time.Time         `json:"-" gorm:"column:status_updated_at"`
	StatusUpdatedBy    uuid.UUID         `json:"-" gorm:"column:status_updated_by"`
	StatusUpdated      entities.User     `json:"-" gorm:"foreignKey:StatusUpdatedBy;references:UserUuid"`
}

func (m *Policy) TableName() string {
	return "policies"
}

type GetAllPoliciesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Keyword     string
}

type GetAllPoliciesResponse struct {
	PolicyUUID      uuid.UUID         `json:"policy_uuid"`
	Name            string            `json:"name"`
	Version         int               `json:"version"`
	LastDraftDate   nullable.NullTime `json:"last_draft_date"`
	Status          string            `json:"status"`
	StatusUpdatedAt nullable.NullTime `json:"status_updated_at"`
	StatusUpdatedBy *UserInfo         `json:"status_updated_by"`
	Owner           *UserInfo         `json:"owner"`
	CreatedAt       time.Time         `json:"-"`
}

func (t *GetAllPoliciesResponse) MarshalJSON() ([]byte, error) {
	type Alias GetAllPoliciesResponse

	o := &struct {
		*Alias
		PolicyUUID string `json:"policy_uuid"`
	}{
		Alias:      (*Alias)(t),
		PolicyUUID: t.PolicyUUID.String(),
	}

	if !t.StatusUpdatedAt.Valid {
		t.StatusUpdatedAt = nullable.NewNullTime(t.CreatedAt)
	}

	if t.StatusUpdatedBy.UserUUID == uuid.Nil {
		t.StatusUpdatedBy = t.Owner
	}

	if t.PolicyUUID == uuid.Nil {
		o.PolicyUUID = ""
	}

	return json.Marshal(o)
}

type UserInfo struct {
	UserUUID  uuid.UUID      `json:"user_uuid"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Ico       sql.NullString `json:"ico"`
	Email     sql.NullString `json:"email"`
}

func NewUserInfo(userUUID uuid.UUID, firstName string, lastName string, ico string, email string) *UserInfo {
	return &UserInfo{
		UserUUID:  userUUID,
		FirstName: sql.NullString{String: firstName, Valid: true},
		LastName:  sql.NullString{String: lastName, Valid: true},
		Ico:       sql.NullString{String: ico, Valid: true},
		Email:     sql.NullString{String: email, Valid: true},
	}
}

func (t *UserInfo) MarshalJSON() ([]byte, error) {
	type Alias UserInfo

	o := &struct {
		*Alias
		UserUUID  string `json:"user_uuid"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Ico       string `json:"ico"`
		Email     string `json:"email"`
	}{
		Alias:     (*Alias)(t),
		UserUUID:  t.UserUUID.String(),
		FirstName: t.FirstName.String,
		LastName:  t.LastName.String,
		Ico:       t.Ico.String,
		Email:     t.Email.String,
	}

	if t.UserUUID == uuid.Nil {
		o.UserUUID = ""
	}

	return json.Marshal(o)
}

type CreatePolicyRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Policy      *Policy
}

type GetPolicyDocumentRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	PolicyUuid  uuid.UUID `json:"policy_uuid"`
	Version     int       `json:"version"`
}

type GetPolicyDocumentResponse struct {
	PolicyUUID      uuid.UUID `json:"policy_uuid"`
	Name            string    `json:"name"`
	Version         int       `json:"version"`
	Document        string    `json:"document"`
	Status          string    `json:"status"`
	StatusUpdatedAt time.Time `json:"status_updated_at"`
	StatusUpdatedBy *UserInfo `json:"status_updated_by"`
	Owner           *UserInfo `json:"owner"`
	CreatedAt       time.Time `json:"created_at"`
}
type SaveDocumentRequestBody struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type SaveDocumentRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	PolicyUUID              uuid.UUID `json:"policy_uuid"`
	SaveDocumentRequestBody *SaveDocumentRequestBody
}

type SaveDocumentResponse struct {
	PolicyUUID uuid.UUID `json:"policy_uuid"`
	Name       string    `json:"name"`
	Document   string    `json:"document"`
}

type GetDocumentResponse struct {
	Name     string
	Document []byte
}

type DeletePolicyRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	PolicyUuid  uuid.UUID `json:"policy_uuid"`
}

type UpdatePolicyDocumentStatusPatchRequestBody struct {
	Status  string `json:"status"`
	Comment string `json:"comment"`
}

type UpdatePolicyDocumentStatus struct {
	CompanyUuid      uuid.UUID                                   `json:"company_uuid"`
	UserUuid         uuid.UUID                                   `json:"user_uuid"`
	PolicyUuid       uuid.UUID                                   `json:"policy_uuid"`
	PatchRequestBody *UpdatePolicyDocumentStatusPatchRequestBody `json:"patch_request_body"`
}

type UpdatePolicyDocumentResponse struct {
	PolicyUuid uuid.UUID `json:"policy_uuid"`
	Status     string    `json:"status"`
	Comment    string    `json:"comment"`
}

type GetPoliciesStatsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetPoliciesStatsResponse struct {
	Total     int `json:"total"`
	Draft     int `json:"draft"`
	Submitted int `json:"submitted"`
	Approved  int `json:"approved"`
	Rejected  int `json:"rejected"`
}
