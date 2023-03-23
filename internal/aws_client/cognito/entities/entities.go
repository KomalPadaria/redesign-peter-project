package entities

import (
	"github.com/dgrijalva/jwt-go"
)

type AWSCognitoAccessTokenClaims struct {
	jwt.StandardClaims

	ClientId  string `json:"client_id,omitempty"`
	OriginJti string `json:"origin_jti,omitempty"`
	EventId   string `json:"event_id,omitempty"`
	TokenUse  string `json:"token_use,omitempty"`
	Scope     string `json:"scope,omitempty"`
	AuthTime  int    `json:"auth_time,omitempty"`
	Username  string `json:"username,omitempty"`
}

type AWSCognitoIDTokenClaims struct {
	jwt.StandardClaims

	OriginJti string `json:"origin_jti,omitempty"`
	AuthTime  int64  `json:"auth_time,omitempty"`
	TokenUse  string `json:"token_use,omitempty"`

	RedesignUserFirstName       string `json:"redesign_user:first_name,omitempty"`
	RedesignUserUserUuid        string `json:"redesign_user:user_uuid,omitempty"`
	RedesignUserPhone           string `json:"redesign_user:phone,omitempty"`
	RedesignUserUsername        string `json:"redesign_user:username,omitempty"`
	RedesignUserEmail           string `json:"redesign_user:email,omitempty"`
	RedesignUserIsFirstLogin    string `json:"redesign_user:is_first_login,omitempty"`
	RedesignUserGroup           string `json:"redesign_user:user_group,omitempty"`
	RedesignCompanyName         string `json:"redesign_company:name,omitempty"`
	RedesignCompanyType         string `json:"redesign_company:type,omitempty"`
	RedesignCompanyCompanyUuid  string `json:"redesign_company:company_uuid,omitempty"`
	RedesignCompanyUserRole     string `json:"redesign_company:user_role,omitempty"`
	RedesignCompanyExternalId   string `json:"redesign_company:external_id,omitempty"`
	RedesignCompanyIndustryType string `json:"redesign_company:industry_type,omitempty"`
	RedesignUserLastName        string `json:"redesign_user:last_name,omitempty"`

	Email            string `json:"email,omitempty"`
	EmailVerified    bool   `json:"email_verified,omitempty"`
	CognitoUsername  string `json:"cognito:username,omitempty"`
	CustomLastName   string `json:"custom:last_name,omitempty"`
	EventId          string `json:"event_id,omitempty"`
	CustomFirstName  string `json:"custom:first_name,omitempty"`
	CustomExternalId string `json:"custom:external_id,omitempty"`
}

type InviteUserDetails struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string

	CompanyName       string
	IndustryType      string
	CompanyExternalId string
}
