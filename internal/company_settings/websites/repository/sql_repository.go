package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
)

const (
	websitesTable = "public.company_websites"
)

type sqlRepository struct {
	db *sql.DB
}

func (s *sqlRepository) CreateWebsite(ctx context.Context, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error) {
	websiteUuid := uuid.New()
	website.CompanyWebsiteUuid = websiteUuid

	website.CreatedAt = nullable.NewNullTime(time.Now())

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(websitesTable).
		Columns(
			"company_website_uuid",
			"company_uuid",
			"url",
			"industry_type",
			"zip",
			"country",
			"state",
			"city",
			"address1",
			"address2",
			"created_at",
			"created_by",
		).
		Values(
			website.CompanyWebsiteUuid,
			website.CompanyUuid,
			website.Url,
			website.IndustryType,
			website.Zip,
			website.Country,
			website.State,
			website.City,
			website.Address1,
			website.Address2,
			website.CreatedAt,
			website.CreatedBy,
		).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "website already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		return nil, err
	}

	return website, nil
}

func (s *sqlRepository) GetWebsiteById(ctx context.Context, websiteUuid *uuid.UUID) (*entities.CompanyWebsite, error) {
	website := &entities.CompanyWebsite{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_website_uuid",
			"company_uuid",
			"url",
			"industry_type",
			"zip",
			"country",
			"state",
			"city",
			"address1",
			"address2",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).
		From(websitesTable).
		Where(sq.Eq{"company_website_uuid": websiteUuid}).
		RunWith(s.db).
		QueryRowContext(ctx).Scan(
		&website.CompanyWebsiteUuid,
		&website.CompanyUuid,
		&website.Url,
		&website.IndustryType,
		&website.Zip,
		&website.Country,
		&website.State,
		&website.City,
		&website.Address1,
		&website.Address2,
		&website.CreatedAt,
		&website.UpdatedAt,
		&website.CreatedBy,
		&website.UpdatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return website, nil
}

func (s *sqlRepository) GetAllWebsites(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.CompanyWebsite, error) {
	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_website_uuid",
			"company_uuid",
			"url",
			"industry_type",
			"zip",
			"country",
			"state",
			"city",
			"address1",
			"address2",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).
		From(websitesTable).
		Where(sq.Eq{"company_uuid": companyUuid}).
		OrderBy("created_at desc").
		RunWith(s.db).
		QueryContext(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	websites := make([]*entities.CompanyWebsite, 0)

	for rows.Next() {
		website := &entities.CompanyWebsite{}
		err = rows.Scan(
			&website.CompanyWebsiteUuid,
			&website.CompanyUuid,
			&website.Url,
			&website.IndustryType,
			&website.Zip,
			&website.Country,
			&website.State,
			&website.City,
			&website.Address1,
			&website.Address2,
			&website.CreatedAt,
			&website.UpdatedAt,
			&website.CreatedBy,
			&website.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}

		websites = append(websites, website)
	}

	return websites, nil
}

func (s *sqlRepository) UpdateWebsite(ctx context.Context, website *entities.CompanyWebsite) (*entities.CompanyWebsite, error) {
	oldWebsite, err := s.GetWebsiteById(ctx, &website.CompanyWebsiteUuid)
	if err != nil {
		return nil, err
	}

	if oldWebsite == nil {
		return nil, ErrWebsiteNotExists
	}

	stmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(websitesTable)

	if website.CompanyWebsiteUuid != oldWebsite.CompanyWebsiteUuid {
		stmt = stmt.Set("company_website_uuid", website.CompanyWebsiteUuid)
	}
	if website.CompanyUuid != oldWebsite.CompanyUuid {
		stmt = stmt.Set("company_uuid", website.CompanyUuid)
	}
	if website.Url != oldWebsite.Url {
		stmt = stmt.Set("url", website.Url)
	}
	if website.IndustryType != oldWebsite.IndustryType {
		stmt = stmt.Set("industry_type", website.IndustryType)
	}
	if website.Zip != oldWebsite.Zip {
		stmt = stmt.Set("zip", website.Zip)
	}
	if website.Country != oldWebsite.Country {
		stmt = stmt.Set("country", website.Country)
	}
	if website.State != oldWebsite.State {
		stmt = stmt.Set("state", website.State)
	}
	if website.City != oldWebsite.City {
		stmt = stmt.Set("city", website.City)
	}
	if website.Address1 != oldWebsite.Address1 {
		stmt = stmt.Set("address1", website.Address1)
	}
	if website.Address2 != oldWebsite.Address2 {
		stmt = stmt.Set("address2", website.Address2)
	}

	website.UpdatedAt = nullable.NewNullTime(time.Now())
	stmt = stmt.Set("updated_at", website.UpdatedAt)
	stmt = stmt.Set("updated_by", website.UpdatedBy)

	stmt = stmt.Where(sq.Eq{"company_website_uuid": website.CompanyWebsiteUuid})

	_, err = stmt.RunWith(s.db).ExecContext(ctx)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "website already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return nil, err
	}

	website.CreatedAt = oldWebsite.CreatedAt
	website.CreatedBy = oldWebsite.CreatedBy

	return website, nil
}

func (s *sqlRepository) DeleteWebsite(ctx context.Context, websiteUuid *uuid.UUID) (*uuid.UUID, error) {
	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(websitesTable).
		Where(sq.Eq{"company_website_uuid": websiteUuid}).
		RunWith(s.db).ExecContext(ctx)
	if err != nil {
		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return &uuid.Nil, err
	}

	return websiteUuid, nil
}
