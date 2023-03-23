package entities

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type CompanyAddress struct {
	CompanyAddressUuid uuid.UUID          `json:"company_address_uuid"`
	CompanyUuid        uuid.UUID          `json:"company_uuid"`
	Zip                string             `json:"zip"`
	Country            string             `json:"country"`
	State              string             `json:"state"`
	City               string             `json:"city"`
	Address1           string             `json:"address1"`
	Address2           string             `json:"address2"`
	Status             string             `json:"status"`
	Facilities         []*CompanyFacility `json:"facilities"`
	CreatedAt          nullable.NullTime  `json:"created_at"`
	UpdatedAt          nullable.NullTime  `json:"updated_at"`
	CreatedBy          uuid.UUID          `json:"created_by"`
	UpdatedBy          uuid.UUID          `json:"updated_by"`
}

func (c *CompanyAddress) MarshalJSON() ([]byte, error) {
	type Alias CompanyAddress

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

	if c.Facilities == nil {
		c.Facilities = make([]*CompanyFacility, 0)
	}

	return json.Marshal(o)
}

type CompanyAddressW struct {
	CompanyAddressUuid uuid.UUID           `json:"company_address_uuid"`
	CompanyUuid        uuid.UUID           `json:"company_uuid"`
	Zip                string              `json:"zip"`
	Country            string              `json:"country"`
	State              string              `json:"state"`
	City               string              `json:"city"`
	Address1           string              `json:"address1"`
	Address2           string              `json:"address2"`
	Status             string              `json:"status"`
	Facilities         []*CompanyFacilityW `json:"facilities"`
	CreatedAt          nullable.NullTime   `json:"created_at"`
	UpdatedAt          nullable.NullTime   `json:"updated_at"`
	CreatedBy          uuid.UUID           `json:"created_by"`
	UpdatedBy          uuid.UUID           `json:"updated_by"`
}

func (c *CompanyAddressW) MarshalJSON() ([]byte, error) {
	type Alias CompanyAddressW

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

	if c.Facilities == nil {
		c.Facilities = make([]*CompanyFacilityW, 0)
	}

	return json.Marshal(o)
}

type GetCompanyAddressRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type CreateCompanyAddressRequest struct {
	CompanyUuid    uuid.UUID       `json:"company_uuid"`
	UserUuid       uuid.UUID       `json:"user_uuid"`
	CompanyAddress *CompanyAddress `json:"company_address"`
}

type UpdateCompanyAddressRequest struct {
	CompanyUuid        uuid.UUID       `json:"company_uuid"`
	UserUuid           uuid.UUID       `json:"user_uuid"`
	CompanyAddressUuid uuid.UUID       `json:"company_address_uuid"`
	CompanyAddress     *CompanyAddress `json:"company_address"`
}

type UpdateCompanyAddressPatchRequestBody struct {
	Status string `json:"status"`
}

type UpdateCompanyAddressPatchRequest struct {
	CompanyUuid        uuid.UUID                             `json:"company_uuid"`
	UserUuid           uuid.UUID                             `json:"user_uuid"`
	CompanyAddressUuid uuid.UUID                             `json:"company_address_uuid"`
	PatchRequestBody   *UpdateCompanyAddressPatchRequestBody `json:"patch_request_body"`
}

type UpdateCompanyAddressPatchResponse struct {
	CompanyAddressUuid uuid.UUID `json:"company_address_uuid"`
	Status             string    `json:"status"`
}

type DeleteCompanyAddressRequest struct {
	CompanyUuid        uuid.UUID `json:"company_uuid"`
	UserUuid           uuid.UUID `json:"user_uuid"`
	CompanyAddressUuid uuid.UUID `json:"company_address_uuid"`
}

type GetCompanyFacilitiesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Query       string    `json:"query"`
	Status      string    `json:"status"`
}
