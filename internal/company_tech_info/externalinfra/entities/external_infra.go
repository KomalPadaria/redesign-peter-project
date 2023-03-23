package entities

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type TechInfoExternalInfra struct {
	TechInfoExternalInfraUuid uuid.UUID         `json:"tech_info_external_infra_uuid,omitempty"`
	CompanyUuid               uuid.UUID         `json:"company_uuid,omitempty"`
	CompanyWebsiteUuid        uuid.UUID         `json:"company_website_uuid,omitempty"`
	CompanyWebsite            *CompanyWebsite   `json:"company_website,omitempty"`
	IpFrom                    string            `json:"ip_from,omitempty"`
	IpTo                      string            `json:"ip_to,omitempty"`
	Env                       string            `json:"env,omitempty"`
	Location                  string            `json:"location,omitempty"`
	HasPermissions            bool              `json:"has_permissions"`
	HasIDsIps                 bool              `json:"has_ids_ips"`
	IsWhitelisted             bool              `json:"is_whitelisted"`
	Is3rdPartyHosted          bool              `json:"is_3rd_party_hosted"`
	CreatedAt                 nullable.NullTime `json:"created_at,omitempty"`
	UpdatedAt                 nullable.NullTime `json:"updated_at,omitempty"`
	CreatedBy                 uuid.UUID         `json:"created_by,omitempty"`
	UpdatedBy                 uuid.UUID         `json:"updated_by,omitempty"`
}

func (t *TechInfoExternalInfra) MarshalJSON() ([]byte, error) {
	type Alias TechInfoExternalInfra

	o := &struct {
		*Alias
		CreatedBy          string `json:"created_by"`
		UpdatedBy          string `json:"updated_by"`
		CompanyWebsiteUuid string `json:"company_website_uuid,omitempty"`
	}{
		Alias:     (*Alias)(t),
		CreatedBy: t.CreatedBy.String(),
		UpdatedBy: t.UpdatedBy.String(),
	}

	if t.CreatedBy == uuid.Nil {
		o.CreatedBy = ""
	}
	if t.UpdatedBy == uuid.Nil {
		o.UpdatedBy = ""
	}
	if t.CompanyWebsiteUuid != uuid.Nil {
		o.CompanyWebsiteUuid = t.CompanyWebsiteUuid.String()
	}

	return json.Marshal(o)
}

type CreateTechInfoExternalInfraRequest struct {
	CompanyUuid           uuid.UUID              `json:"company_uuid"`
	UserUuid              uuid.UUID              `json:"user_uuid"`
	TechInfoExternalInfra *TechInfoExternalInfra `json:"tech_info_external_infra"`
}

type UpdateTechInfoExternalInfraRequest struct {
	CompanyUuid               uuid.UUID              `json:"company_uuid"`
	UserUuid                  uuid.UUID              `json:"user_uuid"`
	TechInfoExternalInfraUuid uuid.UUID              `json:"tech_info_external_infra_uuid"`
	TechInfoExternalInfra     *TechInfoExternalInfra `json:"tech_info_external_infra"`
}

type GetAllTechInfoExternalInfrasRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetTechInfoExternalInfraByIdRequest struct {
	CompanyUuid               uuid.UUID `json:"company_uuid"`
	UserUuid                  uuid.UUID `json:"user_uuid"`
	TechInfoExternalInfraUuid uuid.UUID `json:"tech_info_external_infra_uuid"`
}

type DeleteTechInfoExternalInfraRequest struct {
	CompanyUuid               uuid.UUID `json:"company_uuid"`
	UserUuid                  uuid.UUID `json:"user_uuid"`
	TechInfoExternalInfraUuid uuid.UUID `json:"tech_info_external_infra_uuid"`
}

type CompanyWebsite struct {
	CompanyWebsiteUUID string `json:"company_website_uuid,omitempty"`
	URL                string `json:"url,omitempty"`
}
