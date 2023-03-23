package entities

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type Framework struct {
	FrameworkUuid uuid.UUID         `json:"frameworks_uuid,omitempty" gorm:"primarykey;column:frameworks_uuid"`
	Name          string            `json:"name,omitempty" gorm:"column:name"`
	CreatedAt     nullable.NullTime `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt     nullable.NullTime `json:"updated_at,omitempty" gorm:"column:updated_at"`
	CreatedBy     uuid.UUID         `json:"created_by,omitempty" gorm:"column:created_by"`
	UpdatedBy     uuid.UUID         `json:"updated_by,omitempty" gorm:"column:updated_by"`
}

func (c *Framework) MarshalJSON() ([]byte, error) {
	type Alias Framework

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

func (m *Framework) TableName() string {
	return "frameworks"
}

type CompanyFrameworks struct {
	CompanyFrameworkUuid uuid.UUID         `json:"company_frameworks_uuid,omitempty" gorm:"primarykey;column:company_frameworks_uuid"`
	CompanyUuid          uuid.UUID         `json:"company_uuid,omitempty" gorm:"column:company_uuid"`
	FrameworksUuid       uuid.UUID         `json:"frameworks_uuid,omitempty" gorm:"column:frameworks_uuid"`
	CreatedAt            nullable.NullTime `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt            nullable.NullTime `json:"updated_at,omitempty" gorm:"column:updated_at"`
	CreatedBy            uuid.UUID         `json:"created_by,omitempty" gorm:"column:created_by"`
	UpdatedBy            uuid.UUID         `json:"updated_by,omitempty" gorm:"column:updated_by"`
}

func (m *CompanyFrameworks) TableName() string {
	return "company_frameworks"
}

type FrameworkControl struct {
	FrameworkControlUuid uuid.UUID         `json:"framework_control_uuid" gorm:"column:framework_control_uuid"`
	FrameworksUuid       uuid.UUID         `json:"-" gorm:"column:frameworks_uuid"`
	Name                 string            `json:"name,omitempty" gorm:"column:name"`
	Domain               string            `json:"domain,omitempty" gorm:"column:domain"`
	Topic                string            `json:"topic,omitempty" gorm:"column:topic"`
	BestPractices        string            `json:"best_practices,omitempty" gorm:"column:best_practices"`
	Solution             string            `json:"solution,omitempty" gorm:"column:solution"`
	Groups               pq.StringArray    `json:"-" gorm:"type:text[]"`
	CreatedAt            nullable.NullTime `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt            nullable.NullTime `json:"updated_at,omitempty" gorm:"column:updated_at"`
	CreatedBy            uuid.UUID         `json:"created_by,omitempty" gorm:"column:created_by"`
	UpdatedBy            uuid.UUID         `json:"updated_by,omitempty" gorm:"column:updated_by"`
}

func (c *FrameworkControl) MarshalJSON() ([]byte, error) {
	type Alias FrameworkControl

	o := &struct {
		*Alias
		CreatedBy string `json:"created_by"`
		UpdatedBy string `json:"updated_by"`
		Name      string `json:"name"`
		Domain    string `json:"domain"`
	}{
		Alias:     (*Alias)(c),
		CreatedBy: c.CreatedBy.String(),
		UpdatedBy: c.UpdatedBy.String(),
		Domain:    fmt.Sprintf("%s - [%s]", c.Domain, strings.Join(c.Groups, ", ")),
		Name:      fmt.Sprintf("%s - %s", c.Name, c.Topic),
	}

	if c.CreatedBy == uuid.Nil {
		o.CreatedBy = ""
	}
	if c.UpdatedBy == uuid.Nil {
		o.UpdatedBy = ""
	}

	return json.Marshal(o)
}

type ControlRemediations struct {
	ControlRemediationsUuid uuid.UUID         `json:"control_remediation_uuid" gorm:"column:control_remediation_uuid"`
	CompanyUuid             uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	FrameworksUuid          uuid.UUID         `json:"frameworks_uuid" gorm:"column:frameworks_uuid"`
	FrameworkControlUuid    uuid.UUID         `json:"framework_control_uuid" gorm:"column:framework_control_uuid"`
	Severity                string            `json:"severity,omitempty" gorm:"column:severity"`
	Comment                 string            `json:"comment,omitempty" gorm:"column:comment"`
	Status                  string            `json:"status,omitempty" gorm:"column:status"`
	CreatedAt               nullable.NullTime `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt               nullable.NullTime `json:"updated_at,omitempty" gorm:"column:updated_at"`
	CreatedBy               uuid.UUID         `json:"created_by,omitempty" gorm:"column:created_by"`
	UpdatedBy               uuid.UUID         `json:"updated_by,omitempty" gorm:"column:updated_by"`
}

// Request Types
type GetFrameworksRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetFrameworkControlRequest struct {
	CompanyUuid   uuid.UUID `json:"company_uuid"`
	UserUuid      uuid.UUID `json:"user_uuid"`
	FrameworkUuid uuid.UUID `json:"framework_uuid"`
}

type GetFrameworkStatsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetFrameworkStatsResponse struct {
	Name      string `json:"name"`
	Completed int    `json:"completed"`
	Total     int    `json:"total"`
}
