package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

const (
	techInfoExternalInfraTable = "public.tech_info_external_infra"
)

type sqlRepository struct {
	db *sql.DB
}

func (s *sqlRepository) CreateTechInfoExternalInfra(ctx context.Context, externalInfra *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error) {
	externalInfraUuid := uuid.New()
	externalInfra.TechInfoExternalInfraUuid = externalInfraUuid

	externalInfra.CreatedAt = nullable.NewNullTime(time.Now())

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(techInfoExternalInfraTable).
		Columns(
			"tech_info_external_infra_uuid",
			"company_uuid",
			"company_website_uuid",
			"ip_from",
			"ip_to",
			"env",
			"location",
			"has_permissions",
			"has_ids_ips",
			"is_whitelisted",
			"is_3rd_party_hosted",
			"created_at",
			"created_by",
		).
		Values(
			externalInfra.TechInfoExternalInfraUuid,
			externalInfra.CompanyUuid,
			externalInfra.CompanyWebsiteUuid,
			externalInfra.IpFrom,
			externalInfra.IpTo,
			externalInfra.Env,
			externalInfra.Location,
			externalInfra.HasPermissions,
			externalInfra.HasIDsIps,
			externalInfra.IsWhitelisted,
			externalInfra.Is3rdPartyHosted,
			externalInfra.CreatedAt,
			externalInfra.CreatedBy,
		).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "tech info external infra already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return nil, err
	}

	return externalInfra, nil
}

func (s *sqlRepository) GetTechInfoExternalInfraById(ctx context.Context, techInfoExternalInfraUuid *uuid.UUID) (*entities.TechInfoExternalInfra, error) {
	externalInfra := &entities.TechInfoExternalInfra{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"tech_info_external_infra_uuid",
			"company_uuid",
			"company_website_uuid",
			"ip_from",
			"ip_to",
			"env",
			"location",
			"has_permissions",
			"has_ids_ips",
			"is_whitelisted",
			"is_3rd_party_hosted",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).
		From(techInfoExternalInfraTable).
		Where(sq.Eq{"tech_info_external_infra_uuid": techInfoExternalInfraUuid}).
		RunWith(s.db).
		QueryRowContext(ctx).Scan(
		&externalInfra.TechInfoExternalInfraUuid,
		&externalInfra.CompanyUuid,
		&externalInfra.CompanyWebsiteUuid,
		&externalInfra.IpFrom,
		&externalInfra.IpTo,
		&externalInfra.Env,
		&externalInfra.Location,
		&externalInfra.HasPermissions,
		&externalInfra.HasIDsIps,
		&externalInfra.IsWhitelisted,
		&externalInfra.Is3rdPartyHosted,
		&externalInfra.CreatedAt,
		&externalInfra.UpdatedAt,
		&externalInfra.CreatedBy,
		&externalInfra.UpdatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return externalInfra, nil
}

func (s *sqlRepository) GetAllTechInfoExternalInfras(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.TechInfoExternalInfra, error) {
	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"tech_info_external_infra_uuid",
			"tech_info_external_infra.company_uuid",
			"ip_from",
			"ip_to",
			"env",
			"location",
			"has_permissions",
			"has_ids_ips",
			"is_whitelisted",
			"is_3rd_party_hosted",
			"tech_info_external_infra.created_at",
			"tech_info_external_infra.updated_at",
			"tech_info_external_infra.created_by",
			"tech_info_external_infra.updated_by",
			"company_websites.url",
			"company_websites.company_website_uuid",
		).
		From(techInfoExternalInfraTable).
		Join("company_websites USING (company_website_uuid)").
		Where(sq.Eq{"tech_info_external_infra.company_uuid": companyUuid}).
		OrderBy("tech_info_external_infra.created_at desc").
		RunWith(s.db).
		QueryContext(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	externalInfras := make([]*entities.TechInfoExternalInfra, 0)

	for rows.Next() {
		externalInfra := &entities.TechInfoExternalInfra{}
		externalInfra.CompanyWebsite = &entities.CompanyWebsite{}
		err = rows.Scan(
			&externalInfra.TechInfoExternalInfraUuid,
			&externalInfra.CompanyUuid,
			&externalInfra.IpFrom,
			&externalInfra.IpTo,
			&externalInfra.Env,
			&externalInfra.Location,
			&externalInfra.HasPermissions,
			&externalInfra.HasIDsIps,
			&externalInfra.IsWhitelisted,
			&externalInfra.Is3rdPartyHosted,
			&externalInfra.CreatedAt,
			&externalInfra.UpdatedAt,
			&externalInfra.CreatedBy,
			&externalInfra.UpdatedBy,
			&externalInfra.CompanyWebsite.URL,
			&externalInfra.CompanyWebsite.CompanyWebsiteUUID,
		)

		if err != nil {
			return nil, err
		}

		externalInfras = append(externalInfras, externalInfra)
	}

	return externalInfras, nil
}

func (s *sqlRepository) UpdateTechInfoExternalInfra(ctx context.Context, externalInfra *entities.TechInfoExternalInfra) (*entities.TechInfoExternalInfra, error) {
	oldTechInfoExternalInfra, err := s.GetTechInfoExternalInfraById(ctx, &externalInfra.TechInfoExternalInfraUuid)
	if err != nil {
		return nil, err
	}

	if oldTechInfoExternalInfra == nil {
		return nil, ErrTechInfoExternalInfraNotExists
	}

	stmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(techInfoExternalInfraTable)

	if externalInfra.CompanyUuid != oldTechInfoExternalInfra.CompanyUuid {
		stmt = stmt.Set("company_uuid", externalInfra.CompanyUuid)
	}
	if externalInfra.CompanyWebsiteUuid != oldTechInfoExternalInfra.CompanyWebsiteUuid {
		stmt = stmt.Set("company_website_uuid", externalInfra.CompanyWebsiteUuid)
	}
	if externalInfra.IpFrom != oldTechInfoExternalInfra.IpFrom {
		stmt = stmt.Set("ip_from", externalInfra.IpFrom)
	}
	if externalInfra.IpTo != oldTechInfoExternalInfra.IpTo {
		stmt = stmt.Set("ip_to", externalInfra.IpTo)
	}
	if externalInfra.Env != oldTechInfoExternalInfra.Env {
		stmt = stmt.Set("env", externalInfra.Env)
	}
	if externalInfra.Location != oldTechInfoExternalInfra.Location {
		stmt = stmt.Set("location", externalInfra.Location)
	}
	if externalInfra.HasPermissions != oldTechInfoExternalInfra.HasPermissions {
		stmt = stmt.Set("has_permissions", externalInfra.HasPermissions)
	}
	if externalInfra.HasIDsIps != oldTechInfoExternalInfra.HasIDsIps {
		stmt = stmt.Set("has_ids_ips", externalInfra.HasIDsIps)
	}
	if externalInfra.IsWhitelisted != oldTechInfoExternalInfra.IsWhitelisted {
		stmt = stmt.Set("is_whitelisted", externalInfra.IsWhitelisted)
	}
	if externalInfra.Is3rdPartyHosted != oldTechInfoExternalInfra.Is3rdPartyHosted {
		stmt = stmt.Set("is_3rd_party_hosted", externalInfra.Is3rdPartyHosted)
	}

	externalInfra.UpdatedAt = nullable.NewNullTime(time.Now())
	stmt = stmt.Set("updated_at", externalInfra.UpdatedAt)
	stmt = stmt.Set("updated_by", externalInfra.UpdatedBy)

	stmt = stmt.Where(sq.Eq{"tech_info_external_infra_uuid": externalInfra.TechInfoExternalInfraUuid})

	_, err = stmt.RunWith(s.db).ExecContext(ctx)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "tech info external infra already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		return nil, err
	}

	externalInfra.CreatedAt = oldTechInfoExternalInfra.CreatedAt
	externalInfra.CreatedBy = oldTechInfoExternalInfra.CreatedBy

	return externalInfra, nil
}

func (s *sqlRepository) DeleteTechInfoExternalInfra(ctx context.Context, techInfoExternalInfraUuid *uuid.UUID) (*uuid.UUID, error) {
	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(techInfoExternalInfraTable).
		Where(sq.Eq{"tech_info_external_infra_uuid": techInfoExternalInfraUuid}).
		RunWith(s.db).ExecContext(ctx)
	if err != nil {
		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return &uuid.Nil, err
	}

	return techInfoExternalInfraUuid, nil
}
