package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
)

type CompanyTypes []string

// Value simply returns the JSON-encoded representation of the struct.
func (a CompanyTypes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the IndustryType implement the sql.Scanner interface. This method
// simply decodes a value into the string slice.
func (a *CompanyTypes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// String simply returns the joined slice elements separated by comma.
func (a CompanyTypes) String() string {
	return strings.Join(a, ",")
}

type Meetings struct {
	MeetingsUuid uuid.UUID         `json:"meetings_uuid,omitempty" gorm:"column:meetings_uuid"`
	Name         string            `json:"name,omitempty" gorm:"column:name"`
	Description  string            `json:"description,omitempty" gorm:"column:description"`
	Duration     string            `json:"duration,omitempty" gorm:"column:duration"`
	Data         *MeetingData      `json:"data,omitempty" gorm:"column:data"`
	CompanyTypes *CompanyTypes     `json:"company_types,omitempty" gorm:"column:company_types"`
	CreatedAt    nullable.NullTime `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt    nullable.NullTime `json:"updated_at,omitempty" gorm:"column:updated_at"`
	CreatedBy    uuid.UUID         `json:"created_by,omitempty" gorm:"column:created_by"`
	UpdatedBy    uuid.UUID         `json:"updated_by,omitempty" gorm:"column:updated_by"`
}

func (c *Meetings) MarshalJSON() ([]byte, error) {
	type Alias Meetings

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

type CompanyMeetingsData struct {
	WebhookPayload *WebhookPayload          `json:"webhook_payload,omitempty"`
	ScheduledEvent *entities.ScheduledEvent `json:"scheduled_event,omitempty"`
}

// Value simply returns the JSON-encoded representation of the struct.
func (a CompanyMeetingsData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the Onboarding map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (a *CompanyMeetingsData) Scan(value interface{}) error {
	if value != nil {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("type assertion to []byte failed")
		}
		return json.Unmarshal(b, &a)
	}
	return nil
}

type CompanyMeetings struct {
	CompanyMeetingsUuid uuid.UUID            `json:"company_meetings_uuid" gorm:"column:company_meetings_uuid"`
	MeetingsUuid        uuid.UUID            `json:"meetings_uuid" gorm:"column:meetings_uuid"`
	Meeting             Meetings             `json:"-" gorm:"foreignKey:MeetingsUuid;references:MeetingsUuid"`
	CompanyUuid         uuid.UUID            `json:"company_uuid" gorm:"column:company_uuid"`
	StartAt             nullable.NullTime    `json:"start_at" gorm:"column:start_at"`
	UtcStartAt          nullable.NullTime    `json:"utc_start_at" gorm:"column:utc_start_at"`
	Host                string               `json:"host" gorm:"column:host"`
	Name                string               `json:"name" gorm:"column:name"`
	Data                *CompanyMeetingsData `json:"data" gorm:"column:data"`
	CreatedAt           nullable.NullTime    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           nullable.NullTime    `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy           uuid.UUID            `json:"created_by" gorm:"column:created_by"`
	UpdatedBy           uuid.UUID            `json:"updated_by" gorm:"column:updated_by"`
}

func (c *CompanyMeetings) TableName() string {
	return "company_meetings"
}

func (c *CompanyMeetings) MarshalJSON() ([]byte, error) {
	type Alias CompanyMeetings

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

type GetMeetingsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetCompanyMeetingsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetCompanyMeetingsResponse struct {
	CompanyMeetingsUuid uuid.UUID              `json:"company_meetings_uuid"`
	MeetingsUuid        uuid.UUID              `json:"meetings_uuid"`
	CompanyUuid         uuid.UUID              `json:"company_uuid"`
	StartAt             nullable.NullTime      `json:"start_at"`
	UtcStartAt          nullable.NullTime      `json:"utc_start_at"`
	MeetingName         sql.NullString         `json:"name"`
	Host                sql.NullString         `json:"host"`
	Data                GetCompanyMeetingsData `json:"data"`
	Duration            string                 `json:"duration"`
	Scheduled           bool                   `json:"scheduled"`
}

func (c *GetCompanyMeetingsResponse) MarshalJSON() ([]byte, error) {
	type Alias GetCompanyMeetingsResponse

	o := &struct {
		*Alias
		MeetingName         string `json:"name"`
		Host                string `json:"host"`
		CompanyMeetingsUuid string `json:"company_meetings_uuid"`
		MeetingsUuid        string `json:"meetings_uuid"`
		CompanyUuid         string `json:"company_uuid"`
	}{
		Alias:               (*Alias)(c),
		MeetingName:         c.MeetingName.String,
		Host:                c.Host.String,
		CompanyMeetingsUuid: c.CompanyMeetingsUuid.String(),
		MeetingsUuid:        c.MeetingsUuid.String(),
		CompanyUuid:         c.CompanyUuid.String(),
	}

	if c.CompanyMeetingsUuid == uuid.Nil {
		o.CompanyMeetingsUuid = ""
	}

	if c.CompanyUuid == uuid.Nil {
		o.CompanyUuid = ""
	}

	if c.MeetingsUuid == uuid.Nil {
		o.MeetingsUuid = ""
	}

	return json.Marshal(o)
}

type GetCompanyMeetingsData struct {
	SchedulingUrl sql.NullString `json:"scheduling_url"`
	CancelUrl     sql.NullString `json:"cancel_url"`
	RescheduleUrl sql.NullString `json:"reschedule_url"`
	Location      sql.NullString `json:"location"`
	Status        sql.NullString `json:"status"`
}

func (c *GetCompanyMeetingsData) MarshalJSON() ([]byte, error) {
	o := &struct {
		SchedulingUrl string `json:"scheduling_url"`
		CancelUrl     string `json:"cancel_url"`
		RescheduleUrl string `json:"reschedule_url"`
		Location      string `json:"location"`
		Status        string `json:"status"`
	}{
		SchedulingUrl: c.SchedulingUrl.String,
		CancelUrl:     c.CancelUrl.String,
		RescheduleUrl: c.RescheduleUrl.String,
		Location:      c.Location.String,
		Status:        c.Status.String,
	}

	return json.Marshal(o)
}
