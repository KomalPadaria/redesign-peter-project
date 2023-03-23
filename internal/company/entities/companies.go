package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	CompanyTypeEngineering = "engineering"
	CompanyTypeCustomer    = "customer"
)

type Onboarding struct {
	Position  int    `json:"position"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Status    string `json:"status"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
}

type OnboardingGroup []Onboarding

// Value simply returns the JSON-encoded representation of the struct.
func (a OnboardingGroup) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the Onboarding map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (a *OnboardingGroup) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

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
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// String simply returns the joined slice elements separated by comma.
func (a IndustryType) String() string {
	return strings.Join(a, ",")
}

type UsersCompany struct {
	CompanyUuid   uuid.UUID       `json:"company_uuid,omitempty"`
	Name          string          `json:"name,omitempty"`
	Type          string          `json:"type"`
	IndustryType  IndustryType    `json:"industry_type,omitempty"`
	Role          string          `json:"role,omitempty"`
	ExternalID    string          `json:"external_id,omitempty"`
	Onboarding    OnboardingGroup `json:"onboarding,omitempty"`
	Address       Address         `json:"address,omitempty"`
	CreatedAt     time.Time       `json:"created_at,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty"`
	CreatedBy     uuid.UUID       `json:"created_by,omitempty"`
	UpdatedBy     uuid.UUID       `json:"updated_by,omitempty"`
	Knowbe4Token  string          `json:"-"`
	JiraEpicId    string          `json:"-"`
	Rapid7SiteIds []string        `json:"-"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

// Value simply returns the JSON-encoded representation of the struct.
func (a Address) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the Address map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Company struct {
	CompanyUuid   uuid.UUID             `json:"company_uuid,omitempty"`
	Name          string                `json:"name,omitempty"`
	UserRole      string                `json:"user_role"`
	Type          string                `json:"type"`
	IndustryType  *IndustryType         `json:"industry_type,omitempty"`
	Onboarding    OnboardingGroup       `json:"onboarding,omitempty"`
	Address       *Address              `json:"address,omitempty"`
	ExternalId    string                `json:"external_id,omitempty"`
	CampaignUsers SecurityCampaignUsers `json:"-"`
	Knowbe4Token  string                `json:"-"`
	JiraEpicId    string                `json:"-"`
	Rapid7SiteIds []string              `json:"-"`
}

type SecurityCampaignUserTemplate struct {
	Email    string `csv:"Email" json:"Email"`
	Name     string `csv:"Name" json:"Name"`
	JobTitle string `csv:"Job Title" json:"Job Title"`
}

// Taken from https://support.knowbe4.com/hc/en-us/articles/204873347
// There are 3 differnt templates for upload users on knowbe4 this struct represents the 3rd one with all fields possible
type SecurityCampaignUser struct {
	Email               string `csv:"Email" json:"Email"`
	Name                string `csv:"Name" json:"Name"`
	FirstName           string `csv:"First Name" json:"First Name"`
	LastName            string `csv:"Last Name" json:"Last Name"`
	Phone               string `csv:"Phone Number" json:"Phone Number"`
	Extension           string `csv:"Extension" json:"Extension"`
	Group               string `csv:"Group" json:"Group"`
	Location            string `csv:"Location" json:"Location"`
	Division            string `csv:"Division" json:"Division"`
	ManagerName         string `csv:"Manager Name" json:"Manager Name"`
	EmployeeNumber      string `csv:"Employee Number" json:"Employee Number"`
	JobTitle            string `csv:"Job Title" json:"Job Title"`
	Password            string `csv:"Password" json:"Password"`
	Mobile              string `csv:"Mobile" json:"Mobile"`
	ProvisioningManaged string `csv:"Provisioning Managed" json:"Provisioning Managed"`
	DateFormat          string `csv:"Date Format" json:"Date Format"`
	RiskBooster         string `csv:"Risk Booster" json:"Risk Booster"`
	Organization        string `csv:"Organization" json:"Organization"`
	Department          string `csv:"Department" json:"Department"`
	Language            string `csv:"Language" json:"Language"`
	Comment             string `csv:"Comment" json:"Comment"`
	EmployeeStartDate   string `csv:"Employee Start Date" json:"Employee Start Date"`
	CustomField1        string `csv:"Custom Field 1" json:"Custom Field 1"`
	CustomField2        string `csv:"Custom Field 2" json:"Custom Field 2"`
	CustomField3        string `csv:"Custom Field 3" json:"Custom Field 3"`
	CustomField4        string `csv:"Custom Field 4" json:"Custom Field 4"`
	CustomDate1         string `csv:"Custom Date 1" json:"Custom Date 1"`
	CustomDate2         string `csv:"Custom Date 2" json:"Custom Date 2"`
	TimeZone            string `csv:"Time Zone" json:"Time Zone"`
	AdminLanguage       string `csv:"Admin Language" json:"Admin Language"`
	PhishingLanguage    string `csv:"Phishing Language" json:"Phishing Language"`
	TrainingLanguage    string `csv:"Training Language" json:"Training Language"`
	EmailAlias          string `csv:"Email Alias" json:"Email Alias"`
}
type SecurityCampaignUsers []SecurityCampaignUser

// Value simply returns the JSON-encoded representation of the struct.
func (a SecurityCampaignUsers) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *SecurityCampaignUsers) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(b, &a)
}

type CompanySFFields struct {
	Knowbe4Token  string   `json:"knowbe4_token" gorm:"column:knowbe4_token"`
	JiraEpicId    string   `json:"jira_epic_id" gorm:"column:jira_epic_id"`
	Rapid7SiteIds []string `json:"rapid7_site_ids" gorm:"column:rapid7_site_ids"`
}
