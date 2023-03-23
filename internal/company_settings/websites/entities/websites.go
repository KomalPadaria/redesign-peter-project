package entities

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
)

type CompanyWebsite struct {
	CompanyWebsiteUuid uuid.UUID         `json:"company_website_uuid,omitempty"`
	CompanyUuid        uuid.UUID         `json:"company_uuid,omitempty"`
	Url                string            `json:"url,omitempty"`
	IndustryType       string            `json:"industry_type,omitempty"`
	Zip                string            `json:"zip,omitempty"`
	Country            string            `json:"country,omitempty"`
	State              string            `json:"state,omitempty"`
	City               string            `json:"city,omitempty"`
	Address1           string            `json:"address1,omitempty"`
	Address2           string            `json:"address2,omitempty"`
	CreatedAt          nullable.NullTime `json:"created_at,omitempty"`
	UpdatedAt          nullable.NullTime `json:"updated_at,omitempty"`
	CreatedBy          uuid.UUID         `json:"created_by,omitempty"`
	UpdatedBy          uuid.UUID         `json:"updated_by,omitempty"`
}

func (c *CompanyWebsite) MarshalJSON() ([]byte, error) {
	type Alias CompanyWebsite

	o := &struct {
		*Alias
		CreatedBy string `json:"created_by"`
		UpdatedBy string `json:"updated_by"`
	}{
		Alias:     (*Alias)(c),
		CreatedBy: c.CreatedBy.String(),
		UpdatedBy: c.UpdatedBy.String(),
	}

	if c.CreatedBy == uuid.Nil {
		o.CreatedBy = ""
	}
	if c.UpdatedBy == uuid.Nil {
		o.UpdatedBy = ""
	}

	return json.Marshal(o)
}

type IndustryType string

// Value simply returns the JSON-encoded representation of the struct.
func (a IndustryType) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the IndustryType implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (a *IndustryType) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// String simply returns the joined slice elements separated by comma.
func (a IndustryType) String() string {
	return string(a)
}

type CreateWebsiteRequest struct {
	CompanyUuid    uuid.UUID       `json:"company_uuid"`
	UserUuid       uuid.UUID       `json:"user_uuid"`
	CompanyWebsite *CompanyWebsite `json:"company_website"`
}

type UpdateWebsiteRequest struct {
	CompanyUuid        uuid.UUID       `json:"company_uuid"`
	UserUuid           uuid.UUID       `json:"user_uuid"`
	CompanyWebsiteUuid uuid.UUID       `json:"company_website_uuid"`
	CompanyWebsite     *CompanyWebsite `json:"company_website"`
}

type GetAllWebsitesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetWebsiteByIdRequest struct {
	CompanyUuid        uuid.UUID `json:"company_uuid"`
	UserUuid           uuid.UUID `json:"user_uuid"`
	CompanyWebsiteUuid uuid.UUID `json:"company_website_uuid"`
}

type DeleteWebsiteRequest struct {
	CompanyUuid        uuid.UUID `json:"company_uuid"`
	UserUuid           uuid.UUID `json:"user_uuid"`
	CompanyWebsiteUuid uuid.UUID `json:"company_website_uuid"`
}
