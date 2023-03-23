package entities

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type CompanyFacility struct {
	CompanyFacilityUuid uuid.UUID           `json:"company_facility_uuid"`
	CompanyAddressUuid  uuid.UUID           `json:"company_address_uuid"`
	CompanyUuid         uuid.UUID           `json:"company_uuid"`
	Name                nullable.NullString `json:"name"`
	Type                nullable.NullString `json:"type"`
	Status              nullable.NullString `json:"status"`
	CreatedAt           nullable.NullTime   `json:"created_at"`
	UpdatedAt           nullable.NullTime   `json:"updated_at"`
	CreatedBy           uuid.UUID           `json:"created_by"`
	UpdatedBy           uuid.UUID           `json:"updated_by"`
}

func (c *CompanyFacility) MarshalJSON() ([]byte, error) {
	type Alias CompanyFacility

	o := &struct {
		*Alias
		CreatedBy          string `json:"created_by,omitempty"`
		UpdatedBy          string `json:"updated_by,omitempty"`
		CreatedAt          string `json:"created_at,omitempty"`
		UpdatedAt          string `json:"updated_at,omitempty"`
		CompanyAddressUuid string `json:"company_address_uuid,omitempty"`
		CompanyUuid        string `json:"company_uuid,omitempty"`
		Name               string `json:"name,omitempty"`
	}{
		Alias:              (*Alias)(c),
		CreatedBy:          "",
		UpdatedBy:          "",
		CompanyAddressUuid: "",
		CompanyUuid:        "",
	}

	if c.Type.Valid && c.Type.String == "Others" {
		o.Name = c.Name.String
	} else {
		o.Name = c.Type.String
	}

	return json.Marshal(o)
}

type CompanyFacilityW struct {
	CompanyFacilityUuid uuid.UUID         `json:"company_facility_uuid"`
	CompanyAddressUuid  uuid.UUID         `json:"company_address_uuid"`
	CompanyUuid         uuid.UUID         `json:"company_uuid"`
	Name                sql.NullString    `json:"name"`
	Type                sql.NullString    `json:"type"`
	Status              sql.NullString    `json:"status"`
	CreatedAt           nullable.NullTime `json:"created_at"`
	UpdatedAt           nullable.NullTime `json:"updated_at"`
	CreatedBy           uuid.UUID         `json:"created_by"`
	UpdatedBy           uuid.UUID         `json:"updated_by"`
}

func (c *CompanyFacilityW) MarshalJSON() ([]byte, error) {
	type Alias CompanyFacilityW

	o := &struct {
		*Alias
		CreatedBy          string `json:"created_by,omitempty"`
		UpdatedBy          string `json:"updated_by,omitempty"`
		CreatedAt          string `json:"created_at,omitempty"`
		UpdatedAt          string `json:"updated_at,omitempty"`
		CompanyAddressUuid string `json:"company_address_uuid,omitempty"`
		CompanyUuid        string `json:"company_uuid,omitempty"`
		Name               string `json:"name,omitempty"`
		Type               string `json:"type"`
		Status             string `json:"status"`
	}{
		Alias:              (*Alias)(c),
		CreatedBy:          "",
		UpdatedBy:          "",
		CompanyAddressUuid: "",
		CompanyUuid:        "",
		Type:               c.Type.String,
		Status:             c.Status.String,
	}

	if c.Type.String == "Others" {
		o.Name = c.Name.String
	} else {
		o.Name = c.Type.String
	}

	return json.Marshal(o)
}
