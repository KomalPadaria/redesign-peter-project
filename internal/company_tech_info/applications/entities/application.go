package entities

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type TechInfoApplication struct {
	TechInfoApplicationUuid uuid.UUID                                    `json:"tech_info_application_uuid"`
	CompanyUuid             uuid.UUID                                    `json:"company_uuid"`
	Name                    string                                       `json:"name"`
	Type                    string                                       `json:"type"`
	Status                  string                                       `json:"status"`
	Envs                    []*ApplicationEnvOnApplicationGetAllResponse `json:"envs"`
	CreatedAt               nullable.NullTime                            `json:"created_at"`
	UpdatedAt               nullable.NullTime                            `json:"updated_at"`
	CreatedBy               uuid.UUID                                    `json:"created_by"`
	UpdatedBy               uuid.UUID                                    `json:"updated_by"`
}

func (a *TechInfoApplication) MarshalJSON() ([]byte, error) {
	type Alias TechInfoApplication

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
	if a.Envs == nil {
		a.Envs = make([]*ApplicationEnvOnApplicationGetAllResponse, 0)
	}

	return json.Marshal(o)
}

type CreateApplicationRequest struct {
	CompanyUuid         uuid.UUID `json:"company_uuid"`
	UserUuid            uuid.UUID `json:"user_uuid"`
	TechInfoApplication *TechInfoApplication
}

type UpdateApplicationRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
	TechInfoApplication     *TechInfoApplication
}

type GetAllApplicationsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Keyword     string
}

type GetApplicationByIdRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
}

type DeleteApplicationRequest struct {
	CompanyUuid             uuid.UUID `json:"company_uuid"`
	UserUuid                uuid.UUID `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
}

type CompanyWebsite struct {
	CompanyWebsiteUUID string `json:"company_website_uuid,omitempty"`
	URL                string `json:"url,omitempty"`
}

type UpdateApplicationPatchRequestBody struct {
	Status string `json:"status"`
}

type UpdateApplicationPatchRequest struct {
	CompanyUuid             uuid.UUID                          `json:"company_uuid"`
	UserUuid                uuid.UUID                          `json:"user_uuid"`
	TechInfoApplicationUuid uuid.UUID                          `json:"tech_info_application_uuid"`
	PatchRequestBody        *UpdateApplicationPatchRequestBody `json:"patch_request_body"`
}

type UpdateApplicationPatchResponse struct {
	TechInfoApplicationUuid uuid.UUID `json:"tech_info_application_uuid"`
	Status                  string    `json:"status,omitempty"`
}

type TechInfoApplicationCreateUpdateResponse struct {
	TechInfoApplicationUuid uuid.UUID         `json:"tech_info_application_uuid"`
	Name                    string            `json:"name"`
	Type                    string            `json:"type"`
	CreatedAt               nullable.NullTime `json:"created_at"`
	UpdatedAt               nullable.NullTime `json:"updated_at"`
	CreatedBy               uuid.UUID         `json:"created_by"`
	UpdatedBy               uuid.UUID         `json:"updated_by"`
}

func (a *TechInfoApplicationCreateUpdateResponse) MarshalJSON() ([]byte, error) {
	type Alias TechInfoApplicationCreateUpdateResponse

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
