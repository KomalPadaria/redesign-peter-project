package entities

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type ApplicationEnvOnApplicationGetAllResponse struct {
	ApplicationEnvUuid      uuid.UUID         `json:"application_env_uuid"`
	CompanyUuid             uuid.UUID         `json:"company_uuid"`
	TechInfoApplicationUuid uuid.UUID         `json:"tech_info_application_uuid"`
	Type                    sql.NullString    `json:"type"`
	Description             sql.NullString    `json:"description"`
	Url                     sql.NullString    `json:"url"`
	HostingProviderType     sql.NullString    `json:"hosting_provider_type"`
	HostingProvider         sql.NullString    `json:"hosting_provider"`
	MfaType                 sql.NullString    `json:"mfa_type"`
	Mfa                     sql.NullString    `json:"mfa"`
	IDsIpsType              sql.NullString    `json:"ids_ips_type"`
	IDsIps                  sql.NullString    `json:"ids_ips"`
	Status                  sql.NullString    `json:"status"`
	Name                    sql.NullString    `json:"name"`
	CreatedAt               nullable.NullTime `json:"-"`
}

func (c *ApplicationEnvOnApplicationGetAllResponse) MarshalJSON() ([]byte, error) {
	type Alias ApplicationEnvOnApplicationGetAllResponse

	o := &struct {
		*Alias
		TechInfoApplicationUuid string `json:"tech_info_application_uuid,omitempty"`
		CompanyUuid             string `json:"company_uuid,omitempty"`
		Type                    string `json:"type"`
		Description             string `json:"description"`
		Url                     string `json:"url"`
		HostingProviderType     string `json:"hosting_provider_type"`
		HostingProvider         string `json:"hosting_provider"`
		MfaType                 string `json:"mfa_type"`
		Mfa                     string `json:"mfa"`
		IDsIpsType              string `json:"ids_ips_type"`
		IDsIps                  string `json:"ids_ips"`
		Status                  string `json:"status"`
		Name                    string `json:"name"`
	}{
		Alias:                   (*Alias)(c),
		TechInfoApplicationUuid: "",
		CompanyUuid:             "",
		Type:                    c.Type.String,
		Description:             c.Description.String,
		Url:                     c.Url.String,
		HostingProviderType:     c.HostingProviderType.String,
		HostingProvider:         c.HostingProvider.String,
		MfaType:                 c.MfaType.String,
		Mfa:                     c.Mfa.String,
		IDsIpsType:              c.IDsIpsType.String,
		IDsIps:                  c.IDsIps.String,
		Status:                  c.Status.String,
		Name:                    c.Name.String,
	}

	return json.Marshal(o)
}

type ApplicationEnv struct {
	ApplicationEnvUuid      uuid.UUID         `json:"application_env_uuid" gorm:"column:application_env_uuid"`
	CompanyUuid             uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	TechInfoApplicationUuid uuid.UUID         `json:"tech_info_application_uuid" gorm:"column:tech_info_application_uuid"`
	Name                    string            `json:"name" gorm:"column:name"`
	Type                    string            `json:"type" gorm:"column:type"`
	Description             string            `json:"description" gorm:"column:description"`
	Url                     string            `json:"url" gorm:"column:url"`
	HostingProviderType     string            `json:"hosting_provider_type" gorm:"column:hosting_provider_type"`
	HostingProvider         string            `json:"hosting_provider" gorm:"column:hosting_provider"`
	MfaType                 string            `json:"mfa_type" gorm:"column:mfa_type"`
	Mfa                     string            `json:"mfa" gorm:"column:mfa"`
	IDsIpsType              string            `json:"ids_ips_type" gorm:"column:ids_ips_type"`
	IDsIps                  string            `json:"ids_ips" gorm:"column:ids_ips"`
	Status                  string            `json:"status" gorm:"column:status"`
	CreatedAt               nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt               nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy               uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy               uuid.UUID         `json:"updated_by" gorm:"column:updated_by"`
}

func (a *ApplicationEnv) MarshalJSON() ([]byte, error) {
	type Alias ApplicationEnv

	o := &struct {
		*Alias
		CreatedBy string `json:"created_by"`
		UpdatedBy string `json:"updated_by"`
	}{
		Alias:     (*Alias)(a),
		CreatedBy: a.CreatedBy.String(),
		UpdatedBy: a.UpdatedBy.String(),
	}

	if a.CreatedBy == uuid.Nil {
		o.CreatedBy = ""
	}
	if a.UpdatedBy == uuid.Nil {
		o.UpdatedBy = ""
	}

	return json.Marshal(o)
}

type ApplicationEnvCreateRequestBody struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	Description         string `json:"description"`
	Url                 string `json:"url"`
	HostingProviderType string `json:"hosting_provider_type"`
	HostingProvider     string `json:"hosting_provider"`
	MfaType             string `json:"mfa_type"`
	Mfa                 string `json:"mfa"`
	IDsIpsType          string `json:"ids_ips_type"`
	IDsIps              string `json:"ids_ips"`
}

type ApplicationEnvUpdateRequestBody struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	Description         string `json:"description"`
	Url                 string `json:"url"`
	HostingProviderType string `json:"hosting_provider_type"`
	HostingProvider     string `json:"hosting_provider"`
	MfaType             string `json:"mfa_type"`
	Mfa                 string `json:"mfa"`
	IDsIpsType          string `json:"ids_ips_type"`
	IDsIps              string `json:"ids_ips"`
	Status              string `json:"status"`
}

type CreateApplicationEnvRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
	ApplicationEnv          *ApplicationEnvCreateRequestBody
}

type UpdateApplicationEnvRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
	ApplicationEnvUuid      uuid.UUID `json:"application_env_uuid"`
	ApplicationEnv          *ApplicationEnvUpdateRequestBody
}

type UpdateApplicationEnvPatchRequestBody struct {
	Status string `json:"status"`
}

type UpdateApplicationEnvPatchRequest struct {
	CompanyUuid             uuid.UUID                             `json:"company_uuid"`
	UserUuid                uuid.UUID                             `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID                             `json:"tech_info_application_uuid"`
	ApplicationEnvUuid      uuid.UUID                             `json:"application_env_uuid"`
	PatchRequestBody        *UpdateApplicationEnvPatchRequestBody `json:"patch_request_body"`
}

type UpdateApplicationEnvPatchResponse struct {
	ApplicationEnvUuid uuid.UUID `json:"application_env_uuid"`
	Status             string    `json:"status,omitempty"`
}

type DeleteApplicationEnvRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
	ApplicationEnvUuid      uuid.UUID `json:"application_env_uuid"`
}
