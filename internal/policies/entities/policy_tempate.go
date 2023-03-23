package entities

import (
	"database/sql/driver"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
)

type IndustryType []string

// Value simply returns the JSON-encoded representation of the struct.
func (a IndustryType) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the IndustryType implement the sql.Scanner interface. This method
// simply decodes a value into the string slice.
func (a *IndustryType) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("IndustryType type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// String simply returns the joined slice elements separated by comma.
func (a IndustryType) String() string {
	return strings.Join(a, ",")
}

type PolicyTemplates struct {
	PolicyTemplateUuid uuid.UUID         `json:"policy_template_uuid" gorm:"column:policy_template_uuid"`
	Name               string            `json:"name" gorm:"column:name"`
	Description        string            `json:"description" gorm:"column:description"`
	Document           string            `json:"document" gorm:"column:document"`
	IndustryType       IndustryType      `json:"industry_type" gorm:"column:industry_type"`
	CreatedAt          nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy          uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy          uuid.UUID         `json:"updated_by" gorm:"column:updated_by"`
}

func (m *PolicyTemplates) TableName() string {
	return "policy_templates"
}

type GetTemplatesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	CompanyType []string  `json:"company_type"`
}

type GetTemplatesResponse struct {
	PolicyTemplateUuid uuid.UUID `json:"policy_template_uuid,omitempty"`
	Name               string    `json:"name,omitempty"`
	Description        string    `json:"description,omitempty"`
}

type CreateDocumentFromTemplateRequest struct {
	CompanyUuid        uuid.UUID `json:"company_uuid"`
	UserUuid           uuid.UUID `json:"user_uuid"`
	PolicyTemplateUuid uuid.UUID `json:"policy_template_uuid"`
}
