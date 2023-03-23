package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	questionnairesEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	companiesTable = "public.companies"
)

type sqlRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
}

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
	result := r.gormDB.WithContext(ctx).Where("company_uuid = ?", companyUuid).
		Preload("ServiceEvidence").
		Preload("ServiceEvidence.AcknowledgedUser").
		Find(&subs)
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

func (r *sqlRepository) CreateCompanySubscription(ctx context.Context, cs *entities.CompanySubscription) error {
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

func (r *sqlRepository) FindByExternalId(ctx context.Context, externalId string) (*entities.UsersCompany, error) {
	userCompany := &entities.UsersCompany{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("company_uuid",
			"name",
			"type",
			"to_json(industry_type)",
			"COALESCE(external_id, '')").
		From(companiesTable).
		Where(sq.And{sq.Eq{"external_id": externalId}, sq.NotEq{"type": entities.CompanyTypeEngineering}}).
		RunWith(r.db).
		QueryRowContext(ctx).Scan(
		&userCompany.CompanyUuid,
		&userCompany.Name,
		&userCompany.Type,
		&userCompany.IndustryType,
		&userCompany.ExternalID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return userCompany, nil
}

func (r *sqlRepository) CreateCompany(ctx context.Context, company *entities.UsersCompany) (uuid.UUID, error) {
	companyUUID := uuid.New()
	var trimmedRapid7SiteIds []string
	for _, v := range company.Rapid7SiteIds {
		trimmedRapid7SiteIds = append(trimmedRapid7SiteIds, strings.TrimSpace(v))
	}

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(companiesTable).
		Columns(
			"company_uuid",
			"name",
			"industry_type",
			"onboarding",
			"address",
			"external_id",
			"knowbe4_token",
			"jira_epic_id",
			"rapid7_site_ids",
			"created_at",
			"updated_at",
		).
		Values(
			companyUUID,
			company.Name,
			pq.Array(company.IndustryType),
			company.Onboarding,
			company.Address,
			company.ExternalID,
			strings.TrimSpace(company.Knowbe4Token),
			strings.TrimSpace(company.JiraEpicId),
			trimmedRapid7SiteIds,
			"NOW()",
			"NOW()").ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	_, err = r.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return uuid.Nil, ErrCompanyAlreadyExists
		}
		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum industry_type") {
				return uuid.Nil, errors.WithMessage(err, "invalid input value industry_type")
			}
			return uuid.Nil, errors.WithMessage(err, "invalid value")
		}

		return uuid.Nil, err
	}

	return companyUUID, nil
}

func (r *sqlRepository) GetUserCompaniesByUserUuid(ctx context.Context, userUUID uuid.UUID, keyword string) ([]*entities.UsersCompany, error) {
	var query sq.SelectBuilder
	var sqlStr string
	var args []interface{}
	var err error

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_uuid",
			"name",
			"type",
			"to_json(industry_type)",
			"address",
			"onboarding",
			"role",
			"COALESCE(external_id, '')").
		From(companiesTable).
		Join("public.company_users USING (company_uuid)").
		Where(sq.Eq{"user_uuid": userUUID})

	if keyword != "" {
		sqlStr, args, err = query.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", keyword)}).ToSql()
		if err != nil {
			return nil, err
		}
	} else {
		sqlStr, args, err = query.ToSql()

		if err != nil {
			return nil, err
		}
	}

	rows, err := r.db.Query(sqlStr, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := make([]*entities.UsersCompany, 0)

	for rows.Next() {
		company := &entities.UsersCompany{}
		err = rows.Scan(
			&company.CompanyUuid,
			&company.Name,
			&company.Type,
			&company.IndustryType,
			&company.Address,
			&company.Onboarding,
			&company.Role,
			&company.ExternalID,
		)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (r *sqlRepository) FindByUUID(ctx context.Context, companyUuid *uuid.UUID) (*entities.Company, error) {
	company := &entities.Company{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("company_uuid",
			"name",
			"type",
			"to_json(industry_type)",
			"onboarding",
			"address",
			"COALESCE(external_id, '')",
			"COALESCE(knowbe4_token, '')",
			"COALESCE(jira_epic_id, '')",
			"campaign_users",
			"rapid7_site_ids",
		).
		From(companiesTable).
		Where(sq.Eq{"company_uuid": companyUuid}).
		RunWith(r.db).
		QueryRowContext(ctx).Scan(
		&company.CompanyUuid,
		&company.Name,
		&company.Type,
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

func (r *sqlRepository) UpdateCompany(ctx context.Context, updateData *entities.UpdateCompanyRequest) error {

	sql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(companiesTable).
		SetMap(updateData.Data).
		Where(sq.Eq{"company_uuid": updateData.CompanyUuid}).RunWith(r.db)

	_, err := sql.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *sqlRepository) GetAllCompanies(ctx context.Context, keyword string) ([]*entities.Company, error) {
	var query sq.SelectBuilder
	var sqlStr string
	var args []interface{}
	var err error

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("company_uuid",
			"name",
			"type",
			"to_json(industry_type)",
			"onboarding",
			"address",
			"COALESCE(external_id, '')",
			"COALESCE(knowbe4_token, '')",
			"COALESCE(jira_epic_id, '')",
			"rapid7_site_ids",
		).
		From(companiesTable).
		OrderBy("created_at desc")

	if keyword != "" {
		sqlStr, args, err = query.Where(sq.ILike{"name": fmt.Sprintf("%%%s%%", keyword)}).ToSql()
		if err != nil {
			return nil, err
		}
	} else {
		sqlStr, args, err = query.ToSql()

		if err != nil {
			return nil, err
		}
	}

	rows, err := r.db.Query(sqlStr, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	companies := make([]*entities.Company, 0)

	for rows.Next() {
		company := &entities.Company{}
		err = rows.Scan(
			&company.CompanyUuid,
			&company.Name,
			&company.Type,
			&company.IndustryType,
			&company.Onboarding,
			&company.Address,
			&company.ExternalId,
			&company.Knowbe4Token,
			&company.JiraEpicId,
			pq.Array(&company.Rapid7SiteIds),
		)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}
	return companies, nil
}

func (s *sqlRepository) CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	categories, err := s.getCategories(ctx)
	if err != nil {
		return err
	}
	companyQuestionnaires := make([]questionnairesEntities.CompanyQuestionnaires, 0)

	for _, c := range categories {
		cq := questionnairesEntities.CompanyQuestionnaires{
			CompanyQuestionnairesUuid: uuid.New(),
			CompanyUuid:               *companyUUID,
			Category:                  c,
			CreatedAt:                 nullable.NewNullTime(time.Now()),
			CreatedBy:                 *userUUID,
		}

		companyQuestionnaires = append(companyQuestionnaires, cq)
	}

	result := s.gormDB.WithContext(ctx).Create(&companyQuestionnaires)
	if result.Error != nil {
		err := result.Error
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}

		return err
	}

	return nil
}

func (s *sqlRepository) getCategories(ctx context.Context) ([]string, error) {
	var categories []string

	err := s.gormDB.WithContext(ctx).
		Model(&questionnairesEntities.Questionnaires{}).
		Distinct().
		Pluck("category", &categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
