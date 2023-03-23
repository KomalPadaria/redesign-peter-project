package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type PolicyHistory struct {
	PolicyHistoryUuid uuid.UUID         `json:"policy_history_uuid" gorm:"column:policy_history_uuid"`
	PolicyUuid        uuid.UUID         `json:"policy_uuid" gorm:"column:policy_uuid"`
	Policy            Policy            `json:"policy" gorm:"foreignKey:PolicyUuid;references:PolicyUuid"`
	Document          string            `json:"document" gorm:"column:document"`
	Comment           string            `json:"comment" gorm:"column:comment"`
	Version           int               `json:"version" gorm:"column:version"`
	CreatedAt         nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy         uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	Created           entities.User     `json:"-" gorm:"foreignKey:CreatedBy;references:UserUuid"`
	UpdatedBy         nullable.NullUUID `json:"updated_by" gorm:"column:updated_by"`
}

func (m *PolicyHistory) TableName() string {
	return "policy_histories"
}

type GetPolicyHistoriesByPolicyRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	PolicyUUID  uuid.UUID `json:"policy_uuid"`
}

type GetPolicyHistoriesByPolicyResponse struct {
	PolicyHistoryUUID uuid.UUID `json:"policy_history_uuid"`
	Version           int       `json:"version"`
	CreatedAt         time.Time `json:"created_at"`
	Owner             *UserInfo `json:"owner"`
}
