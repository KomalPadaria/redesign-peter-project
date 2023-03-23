package entities

import (
	"io"

	"github.com/google/uuid"
)

type FindCompanyByExternalIdRequest struct {
	ExternalId string `json:"external_id"`
}

type FindCompanyByExternalIdResponse struct {
	Company *UsersCompany `json:"company"`
}

type GetCompaniesByUserIdRequest struct {
	UserUuid uuid.UUID
	Keyword  string
}

type GetCompaniesByUserIdResponse struct {
	Companies []*Company
}

type GetCompanyByIdRequest struct {
	CompanyUuid uuid.UUID
}

type UpdateCompanyRequest struct {
	CompanyUuid uuid.UUID
	Data        map[string]interface{}
}

type GetCompanyInfoRequest struct {
	Token string `json:"token"`
}

type GetAllCompaniesRequest struct {
	Keyword string
}

type UploadSecurityCampaignUsersRequest struct {
	CompanyUuid uuid.UUID
	UserUuid    uuid.UUID
	File        io.ReadCloser
}

type GetCampaignUsersRequest struct {
	CompanyUuid uuid.UUID
	UserUuid    uuid.UUID
}

type FileResponse struct {
	File io.ReadCloser
}

type UpdateSFFieldsRequest struct {
	Id      string
	Name    string
	KnowBe4 string
	Rapid7  string
	Jira    string
}
