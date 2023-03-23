package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type User struct {
	UserUuid  uuid.UUID `json:"user_uuid,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
}

func (m *User) TableName() string {
	return "users"
}

type ServiceEvidence struct {
	ServiceEvidencesUuid     uuid.UUID               `json:"service_evidences_uuid" gorm:"column:service_evidences_uuid"`
	CompanySubscriptionsUuid uuid.UUID               `json:"company_subscriptions_uuid" gorm:"column:company_subscriptions_uuid"`
	CompletedOn              nullable.NullTime       `json:"completed_on" gorm:"column:completed_on"`
	AcknowledgedAt           nullable.NullTime       `json:"acknowledged_at" gorm:"column:acknowledged_at"`
	AcknowledgedBy           nullable.NullUUID       `json:"acknowledged_by" gorm:"column:acknowledged_by"`
	AcknowledgedUser         User                    `json:"-" gorm:"foreignKey:AcknowledgedBy;references:UserUuid"`
	Data                     ServiceEvidenceDataList `json:"data" gorm:"column:data"`
	Status                   string                  `json:"status" gorm:"column:status"`
}

func (m *ServiceEvidence) TableName() string {
	return "service_evidences"
}

type ServiceEvidenceData struct {
	FileName  string    `json:"file_name"`
	FileUuid  uuid.UUID `json:"file_uuid"`
	FileExt   string    `json:"file_ext"`
	CreatedAt time.Time `json:"created_at"`
}

type ServiceEvidenceDataList []ServiceEvidenceData

// Value simply returns the JSON-encoded representation of the struct.
func (m ServiceEvidenceDataList) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan makes the Onboarding map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (m *ServiceEvidenceDataList) Scan(value interface{}) error {
	if value != nil {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("type assertion to []byte failed")
		}
		return json.Unmarshal(b, &m)
	}
	return nil
}
