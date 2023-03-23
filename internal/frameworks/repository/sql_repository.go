package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

func (s *sqlRepository) DeleteCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error {
	frameworks, err := s.getFrameworksByNames(ctx, frameworkNames)
	if err != nil {
		return err
	}

	var frameworkUuids []uuid.UUID
	for _, f := range frameworks {
		frameworkUuids = append(frameworkUuids, f.FrameworkUuid)
	}

	result := s.gormDB.WithContext(ctx).Delete(&entities.CompanyFrameworks{}, "company_uuid = ? AND frameworks_uuid IN ?", companyUUID, frameworkUuids)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *sqlRepository) CreateCompanyFrameworksLink(ctx context.Context, companyUUID, userUUID *uuid.UUID, frameworkNames []string) error {
	frameworks, err := s.getFrameworksByNames(ctx, frameworkNames)
	if err != nil {
		return err
	}

	companyFrameworks := make([]entities.CompanyFrameworks, 0)

	for _, f := range frameworks {
		companyFramework := entities.CompanyFrameworks{
			CompanyFrameworkUuid: uuid.New(),
			CompanyUuid:          *companyUUID,
			FrameworksUuid:       f.FrameworkUuid,
			CreatedAt:            nullable.NewNullTime(time.Now()),
		}

		if userUUID != nil {
			companyFramework.CreatedBy = *userUUID
		}

		companyFrameworks = append(companyFrameworks, companyFramework)
	}

	result := s.gormDB.WithContext(ctx).Create(&companyFrameworks)
	if result.Error != nil {
		err := result.Error
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}

		return err
	}

	return nil
}

func (s *sqlRepository) GetFrameworks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.Framework, error) {
	var frameworks []*entities.Framework

	err := s.gormDB.WithContext(ctx).Model(&entities.Framework{}).
		Order("created_at desc").
		Joins("left join public.company_frameworks ON company_frameworks.frameworks_uuid = frameworks.frameworks_uuid").
		Find(&frameworks, "company_uuid = ?", companyUuid).Error
	if err != nil {
		return nil, err
	}

	return frameworks, nil
}

func (s *sqlRepository) GetFrameworkControls(ctx context.Context, companyUuid, userUuid, frameworkUuid *uuid.UUID) ([]*entities.FrameworkControl, error) {
	var framework_controls []*entities.FrameworkControl
	err := s.gormDB.WithContext(ctx).Model(&entities.FrameworkControl{}).
		Order("created_at desc").
		// Joins("join public.control_remediations ON framework_controls.framework_control_uuid = control_remediations.framework_control_uuid AND framework_controls.frameworks_uuid = control_remediations.frameworks_uuid").
		Where("frameworks_uuid = ?", frameworkUuid).
		Find(&framework_controls).Error

	if err != nil {
		return nil, err
	}

	return framework_controls, nil
}

func (s *sqlRepository) GetFrameworkStats(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetFrameworkStatsResponse, error) {
	type qResult struct {
		Name      string `json:"name" gorm:"column:name"`
		Total     int    `json:"total" gorm:"column:total"`
		Completed int    `json:"completed" gorm:"column:completed"`
	}

	var rs []*qResult
	err := s.gormDB.WithContext(ctx).Debug().Table("frameworks_questionnaires fq").
		Select("f.name,count(q.questionnaires_uuid) as total,count(qa.questionnaire_answers_uuid) as completed").
		Joins("left join frameworks f on f.frameworks_uuid = fq.frameworks_uuid").
		Joins("left join questionnaires q on q.questionnaires_uuid = fq.questionnaires_uuid").
		Joins("left join company_frameworks cf on cf.frameworks_uuid = fq.frameworks_uuid").
		Joins("left join questionnaire_answers qa on qa.questionnaires_uuid = q.questionnaires_uuid and qa.company_uuid = ?", companyUuid).
		Where("cf.company_uuid = ?", companyUuid).
		Group("f.name").
		Order("f.name").
		Find(&rs).Error
	if err != nil {
		return nil, err
	}
	var controlStats []*entities.GetFrameworkStatsResponse

	for _, r := range rs {
		controlStats = append(controlStats, &entities.GetFrameworkStatsResponse{
			Name:      r.Name,
			Completed: r.Completed,
			Total:     r.Total,
		})
	}
	return controlStats, nil
}

func (s *sqlRepository) CreateFrameworkControlRemediations(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	// Load all MPA & CIS controls added at time of ENV creation
	controls, err := s.getFrameworkControls(ctx, companyUUID)
	if err != nil {
		return err
	}

	controlRemediations := make([]entities.ControlRemediations, 0)

	// By default all control remediations are in "pending" state
	for _, c := range controls {
		remediation := entities.ControlRemediations{
			ControlRemediationsUuid: uuid.New(),
			CompanyUuid:             *companyUUID,
			FrameworksUuid:          c.FrameworksUuid,
			FrameworkControlUuid:    c.FrameworkControlUuid,
			Status:                  "pending",
			CreatedAt:               nullable.NewNullTime(time.Now()),
			CreatedBy:               *userUUID,
		}

		controlRemediations = append(controlRemediations, remediation)
	}

	result := s.gormDB.WithContext(ctx).Create(&controlRemediations)
	if result.Error != nil {
		err := result.Error
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}

		return err
	}

	return nil
}

func (s *sqlRepository) GetCompanyByUUID(ctx context.Context, companyUuid *uuid.UUID) (*companyEntities.Company, error) {
	company := &companyEntities.Company{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("company_uuid",
			"name",
			"to_json(industry_type)",
			"onboarding",
			"address",
			"COALESCE(external_id, '')",
			"COALESCE(knowbe4_token, '')",
			"COALESCE(jira_epic_id, '')",
			"campaign_users",
			"rapid7_site_ids",
		).
		From("public.companies").
		Where(sq.Eq{"company_uuid": companyUuid}).
		RunWith(s.sqlDB).
		QueryRowContext(ctx).Scan(
		&company.CompanyUuid,
		&company.Name,
		&company.IndustryType,
		&company.Onboarding,
		&company.Address,
		&company.ExternalId,
		&company.Knowbe4Token,
		&company.JiraEpicId,
		&company.CampaignUsers,
		pq.Array(&company.Rapid7SiteIds),
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (r *sqlRepository) CreateCompanySubscription(ctx context.Context, cs *companyEntities.CompanySubscription) error {
	cs.CompanySubscriptionsUuid = uuid.New()
	cs.CreatedAt = nullable.NewNullTime(time.Now())

	result := r.gormDB.WithContext(ctx).Create(cs)
	if result.Error != nil {
		errCause := errors.Cause(result.Error)
		if ec, ok := errCause.(*pgconn.PgError); ok && ec.Code == pgerrcode.UniqueViolation {
			return nil
		}

		return result.Error
	}

	return nil
}

func (s *sqlRepository) getFrameworkControls(ctx context.Context, companyUUID *uuid.UUID) ([]*entities.FrameworkControl, error) {
	var frameworkControls []*entities.FrameworkControl

	subQuery := s.gormDB.Select("frameworks_uuid").Where("company_uuid = ?", companyUUID).Table("company_frameworks")

	err := s.gormDB.WithContext(ctx).Model(&entities.FrameworkControl{}).
		Where("frameworks_uuid in (?)", subQuery).
		Find(&frameworkControls).Error
	if err != nil {
		return nil, err
	}

	return frameworkControls, nil
}

func (s *sqlRepository) getFrameworksByNames(ctx context.Context, frameworkNames []string) ([]*entities.Framework, error) {
	var frameworks []*entities.Framework

	err := s.gormDB.WithContext(ctx).
		Model(&entities.Framework{}).
		Where("name IN ?", frameworkNames).
		Find(&frameworks).Error
	if err != nil {
		return nil, err
	}

	return frameworks, nil
}
