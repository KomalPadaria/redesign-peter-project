package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
}

func (r *sqlRepository) UpdateEvidence(ctx context.Context, evidenceUuid *uuid.UUID, updateData map[string]interface{}) error {
	result := r.gormDB.WithContext(ctx).
		Model(&companyEntities.ServiceEvidence{}).
		Where("service_evidences_uuid = ? ", evidenceUuid).
		Updates(updateData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *sqlRepository) DeleteEvidence(ctx context.Context, evidenceUuid *uuid.UUID) error {
	result := r.gormDB.WithContext(ctx).Where("service_evidences_uuid = ?", evidenceUuid).Delete(&companyEntities.ServiceEvidence{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *sqlRepository) GetEvidenceByEvidenceId(ctx context.Context, evidenceUuid *uuid.UUID) (*companyEntities.ServiceEvidence, error) {
	var evidence *companyEntities.ServiceEvidence
	result := r.gormDB.WithContext(ctx).Where("service_evidences_uuid = ?", evidenceUuid).
		Preload("AcknowledgedUser").Find(&evidence)
	if result.Error != nil {
		return nil, result.Error
	}
	return evidence, nil
}

func (r *sqlRepository) CreateEvidences(ctx context.Context, serviceReports companyEntities.ServiceEvidence) error {
	result := r.gormDB.WithContext(ctx).Create(serviceReports)
	if result.Error != nil {
		errCause := errors.Cause(result.Error)
		if ec, ok := errCause.(*pgconn.PgError); ok && ec.Code == pgerrcode.UniqueViolation {
			return nil
		}
		return result.Error
	}

	return nil
}

func (r *sqlRepository) DeleteEvidenceFile(ctx context.Context, companySubscriptionsUuid *uuid.UUID, reportName string) error {
	result := r.gormDB.WithContext(ctx).Where("company_subscriptions_uuid = ? AND display_name = ?", companySubscriptionsUuid, reportName).Delete(&companyEntities.ServiceEvidence{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/*func (r *sqlRepository) CreateCompanySubscription(ctx context.Context, cs []entities.CompanySubscription) error {
	cs.CompanySubscriptionsUuid = uuid.New()
	cs.CreatedAt = nullable.NewNullTime(time.Now())

	result := r.gormDB.WithContext(ctx).Create(cs)
	if result.Error != nil {
		return result.Error
	}

	return nil
}*/

/*
func (r *sqlRepository) GetSubscriptionByName(ctx context.Context, companyUuid *uuid.UUID, subscriptionName string) (*entities.CompanySubscription, error) {
	var sub *entities.CompanySubscription
	result := r.gormDB.WithContext(ctx).Where("company_uuid = ? AND name = ?", companyUuid, subscriptionName).Find(&sub)
	if result.Error != nil {
		return nil, result.Error
	}
	return sub, nil
}

func (r *sqlRepository) GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]entities.CompanySubscription, error) {
	var subs []entities.CompanySubscription
	result := r.gormDB.WithContext(ctx).Where("company_uuid = ?", companyUuid).Find(&subs)
	if result.Error != nil {
		return nil, result.Error
	}

	return subs, nil
}

func (r *sqlRepository) DeleteCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string) (*entities.CompanySubscription, error) {
	var cs *entities.CompanySubscription
	result := r.gormDB.WithContext(ctx).Clauses(clause.Returning{}).Where("company_uuid = ? AND sf_subscription_id = ?", companyUuid, subscriptionId).Delete(&cs)
	if result.Error != nil {
		return nil, result.Error
	}

	return cs, nil
}

func (r *sqlRepository) UpdateCompanySubscription(ctx context.Context, companyUuid *uuid.UUID, subscriptionId string, cs *entities.CompanySubscription) error {
	cs.UpdatedAt = nullable.NewNullTime(time.Now())
	result := r.gormDB.WithContext(ctx).
		Model(&entities.CompanySubscription{}).
		Where("company_uuid = ? AND sf_subscription_id = ?", companyUuid, subscriptionId).
		Updates(cs)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
*/
