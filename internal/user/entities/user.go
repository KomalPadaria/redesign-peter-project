package entities

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type User struct {
	UserUuid           uuid.UUID         `json:"user_uuid,omitempty"`
	CurrentCompanyUuid uuid.UUID         `json:"current_company_uuid,omitempty"`
	Username           string            `json:"username,omitempty"`
	FirstName          string            `json:"first_name,omitempty"`
	LastName           string            `json:"last_name,omitempty"`
	Email              string            `json:"email,omitempty"`
	Phone              string            `json:"phone,omitempty"`
	IsFirstLogin       bool              `json:"is_first_login,omitempty"`
	ExternalID         string            `json:"external_id,omitempty"`
	JobTitle           string            `json:"job,omitempty"`
	Mfa                pq.StringArray    `json:"mfa" gorm:"type:text[]"`
	CreatedAt          nullable.NullTime `json:"created_at,omitempty"`
	UpdatedAt          nullable.NullTime `json:"updated_at,omitempty"`
	CreatedBy          uuid.UUID         `json:"created_by,omitempty"`
	UpdatedBy          uuid.UUID         `json:"updated_by,omitempty"`
	CompanyRole        string            `json:"role,omitempty"`
	CompanyStatus      string            `json:"status,omitempty"`
	Group              string            `json:"-"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User

	o := &struct {
		*Alias
		CurrentCompanyUuid string `json:"current_company_uuid,omitempty"`
		CreatedBy          string `json:"created_by,omitempty"`
		UpdatedBy          string `json:"updated_by,omitempty"`
	}{
		Alias:     (*Alias)(u),
		CreatedBy: u.CreatedBy.String(),
		UpdatedBy: u.UpdatedBy.String(),
	}

	if u.CurrentCompanyUuid == uuid.Nil {
		o.CurrentCompanyUuid = ""
	}
	if u.CreatedBy == uuid.Nil {
		o.CreatedBy = ""
	}
	if u.UpdatedBy == uuid.Nil {
		o.UpdatedBy = ""
	}

	return json.Marshal(o)
}

type CompanyUser struct {
	UserUuid    uuid.UUID         `json:"user_uuid,omitempty"`
	CompanyUuid uuid.UUID         `json:"company_uuid,omitempty"`
	Role        string            `json:"role,omitempty"`
	Status      string            `json:"status,omitempty"`
	CreatedAt   nullable.NullTime `json:"created_at,omitempty"`
	UpdatedAt   nullable.NullTime `json:"updated_at,omitempty"`
	CreatedBy   uuid.UUID         `json:"created_by,omitempty"`
	UpdatedBy   uuid.UUID         `json:"updated_by,omitempty"`
}

type CreateCompanyAndUserRequest struct {
	User    User                         `json:"user"`
	Company companyEntities.UsersCompany `json:"company"`
}

type CreateCompanyAndUserResponse struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetUserCompanyInfoByUserIdRequest struct {
	UserId uuid.UUID
}

type GetUserCompanyInfoByUserIdResponse struct {
	UserUuid        uuid.UUID                  `json:"user_uuid,omitempty"`
	Username        string                     `json:"username,omitempty"`
	FirstName       string                     `json:"first_name,omitempty"`
	LastName        string                     `json:"last_name,omitempty"`
	Email           string                     `json:"email,omitempty"`
	Phone           string                     `json:"phone,omitempty"`
	IsFirstLogin    bool                       `json:"is_first_login"`
	Status          string                     `json:"status"`
	Group           string                     `json:"group"`
	Job             string                     `json:"job"`
	IsMfaSmsEnabled bool                       `json:"isMfaSmsEnabled"`
	IsMfaAppEnabled bool                       `json:"isMfaAppEnabled"`
	Company         *companyEntities.Company   `json:"company,omitempty"`
	Companies       []*companyEntities.Company `json:"companies,omitempty"`
}

type GetSFUserAndCompanyInfoRequest struct {
	Token string `json:"token"`
}

type GetSFUserAndCompanyInfoResponse struct {
	User    SFUserResponse      `json:"user"`
	Company []SFCompanyResponse `json:"company"`
}

type SFUserResponse struct {
	Username   string `json:"username,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

type SFCompanyResponse struct {
	Name         string                       `json:"name,omitempty"`
	IndustryType companyEntities.IndustryType `json:"industry_type,omitempty"`
	Role         string                       `json:"role,omitempty"`
	ExternalID   string                       `json:"external_id,omitempty"`
}

type ActivateDeactivateUserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ActivateDeactivateUserResponse struct {
	Message string `json:"message,omitempty"`
}

type ConfirmUserRequest struct {
	EmailId   string `json:"email"`
	CompanyId string `json:"company_uuid"`
}

type UpdateCurrentCompanyRequest struct {
	UserUUID    uuid.UUID
	CompanyUUID uuid.UUID
}

type UpdateCurrentCompanyResponse struct {
	UserUUID    uuid.UUID
	CompanyUUID uuid.UUID
}

type GetCompanyUsersRequest struct {
	CompanyUUID uuid.UUID
	UserUUID    uuid.UUID
}

type CreateUserRequestBody struct {
	Email     string `json:"email"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserRequest struct {
	CompanyUUID           uuid.UUID
	UserUUID              uuid.UUID
	CreateUserRequestBody *CreateUserRequestBody
}

type CreateUserResponse struct {
	User
}

type UpdateCompanyUserLinkRequest struct {
	CompanyUuid uuid.UUID
	ReqUserUuid uuid.UUID
	UserUuid    uuid.UUID
	CompanyUser *CompanyUser
}

type DeleteCompanyUserLinkRequest struct {
	CompanyUuid uuid.UUID
	ReqUserUuid uuid.UUID
	UserUuid    uuid.UUID
}

type UpdateUserRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	JobTitle   string `json:"job"`
	MfaMethods []string
}

type UpdateUserRequestBody struct {
	CompanyUuid     string `json:"company_uuid"`
	UserUuid        string `json:"user_uuid"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	JobTitle        string `json:"job"`
	IsMfaSmsEnabled bool   `json:"isMfaSmsEnabled"`
	IsMfaAppEnabled bool   `json:"isMfaAppEnabled"`
}

type GetCompaniesRequest struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	Keyword  string
}

type GetCompaniesResponse struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	CompanyName string    `json:"name"`
	// Role of a user in a company (company_user table)
	CompanyRole string `json:"role"`
	// Group of a user (users table)
	UserGroup string `json:"group"`
}
