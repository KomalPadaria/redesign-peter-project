package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type CompanySubscription struct {
	CompanySubscriptionsUuid uuid.UUID         `json:"company_subscriptions_uuid" gorm:"column:company_subscriptions_uuid"`
	CompanyUuid              uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	Name                     string            `json:"name" gorm:"column:name"`
	SfSubscriptionID         string            `json:"sf_subscription_id" gorm:"column:sf_subscription_id"`
	SfProductID              string            `json:"sf_product_id" gorm:"column:sf_product_id"`
	Type                     string            `json:"type" gorm:"column:type"`
	SubType                  string            `json:"sub_type" gorm:"column:sub_type"`
	Status                   string            `json:"status" gorm:"column:status"`
	StartDate                time.Time         `json:"start_date" gorm:"column:start_date"`
	EndDate                  time.Time         `json:"end_date" gorm:"column:end_date"`
	CreatedAt                nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy                uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy                uuid.UUID         `json:"updated_by" gorm:"column:updated_by"`
	ServiceEvidence          []ServiceEvidence `json:"evidences" gorm:"foreignKey:CompanySubscriptionsUuid;references:CompanySubscriptionsUuid"`
}

func (m *CompanySubscription) TableName() string {
	return "company_subscriptions"
}
