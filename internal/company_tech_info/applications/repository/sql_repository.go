package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	applicationsTable    = "public.tech_info_applications"
	applicationEnvsTable = "public.application_envs"
)

type sqlRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func (s *sqlRepository) DeleteApplicationEnv(ctx context.Context, applicationEnvUuid *uuid.UUID) (*uuid.UUID, error) {
	result := s.gormDB.WithContext(ctx).
		Delete(&entities.ApplicationEnv{}, "application_env_uuid = ?", applicationEnvUuid)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "application env not found"}
	}

	return applicationEnvUuid, nil
}

func (s *sqlRepository) UpdateApplicationEnvPatch(ctx context.Context, userUUID, applicationEnvUuid *uuid.UUID, req *entities.UpdateApplicationEnvPatchRequestBody) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.ApplicationEnv{}).
		WithContext(ctx).
		Where("application_env_uuid = ?", applicationEnvUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     req.Status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum application_envs_status") {
				return errors.WithMessage(result.Error, "invalid input value status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}
		return result.Error
	}

	return nil
}

func (s *sqlRepository) GetApplicationEnvById(ctx context.Context, applicationEnvUuid *uuid.UUID) (*entities.ApplicationEnv, error) {
	var applicationEnv entities.ApplicationEnv
	result := s.gormDB.WithContext(ctx).Find(&applicationEnv, "application_env_uuid = ?", applicationEnvUuid)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &appError.ErrNotFound{Message: "application env not found"}
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicationEnv, nil
}

func (s *sqlRepository) UpdateApplicationEnv(ctx context.Context, applicationEnv *entities.ApplicationEnv) (*entities.ApplicationEnv, error) {
	applicationEnv.UpdatedAt = nullable.NewNullTime(time.Now())
	result := s.gormDB.WithContext(ctx).Where("application_env_uuid = ?", applicationEnv.ApplicationEnvUuid).Updates(applicationEnv)
	if result.Error != nil {
		err := result.Error

		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "application env already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum application_env_type") {
				return nil, errors.WithMessage(err, "invalid input value for type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum hosting_provider") {
				return nil, errors.WithMessage(err, "invalid input value for hosting_provider_type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum mfa") {
				return nil, errors.WithMessage(err, "invalid input value for mfa_type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum ids_ips_solution") {
				return nil, errors.WithMessage(err, "invalid input value ids_ips_type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum application_envs_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "application env not found"}
	}

	return applicationEnv, nil
}

func (s *sqlRepository) CreateApplicationEnv(ctx context.Context, applicationEnv *entities.ApplicationEnv) (*entities.ApplicationEnv, error) {
	applicationEnv.ApplicationEnvUuid = uuid.New()
	applicationEnv.CreatedAt = nullable.NewNullTime(time.Now())

	result := s.gormDB.WithContext(ctx).Create(applicationEnv)
	if result.Error != nil {
		err := result.Error

		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "application env already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum application_env_type") {
				return nil, errors.WithMessage(err, "invalid input value for type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum hosting_provider") {
				return nil, errors.WithMessage(err, "invalid input value for hosting_provider_type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum mfa") {
				return nil, errors.WithMessage(err, "invalid input value for mfa_type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum ids_ips_solution") {
				return nil, errors.WithMessage(err, "invalid input value ids_ips_type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum application_envs_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return applicationEnv, nil
}

func (s *sqlRepository) UpdateApplicationPatch(ctx context.Context, userUUID, applicationUuid *uuid.UUID, req *entities.UpdateApplicationPatchRequestBody) error {
	updateBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(applicationsTable).
		Set("updated_by", userUUID).
		Set("updated_at", nullable.NewNullTime(time.Now()))

	if req.Status != "" {
		updateBuilder = updateBuilder.Set("status", req.Status)
	}

	sqlStr, args, err := updateBuilder.
		Where(sq.Eq{"tech_info_application_uuid": applicationUuid}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := s.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum tech_info_applications_status") {
				return errors.WithMessage(err, "invalid input value for status")
			}

			return errors.WithMessage(err, "invalid value")
		}

		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return &appError.ErrNotFound{Message: "application not found"}
	}

	return nil
}

func (s *sqlRepository) CreateApplication(ctx context.Context, application *entities.TechInfoApplication) (*entities.TechInfoApplication, error) {
	appUuid := uuid.New()
	application.TechInfoApplicationUuid = appUuid
	application.CreatedAt = nullable.NewNullTime(time.Now())
	application.Status = "Active"

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(applicationsTable).
		Columns(
			"tech_info_application_uuid",
			"company_uuid",
			"name",
			"status",
			"type",
			"created_at",
			"created_by",
		).
		Values(
			application.TechInfoApplicationUuid,
			application.CompanyUuid,
			application.Name,
			application.Status,
			application.Type,
			application.CreatedAt,
			application.CreatedBy,
		).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "application already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidLength(err) {
			return nil, errors.Errorf("The maximum length for Application Name is 250 characters")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum application_type") {
				return nil, errors.WithMessage(err, "invalid input value for type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum tech_info_applications_status") {
				return nil, errors.WithMessage(err, "invalid input value for status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return application, nil
}

func (s *sqlRepository) GetApplicationById(ctx context.Context, applicationUuid *uuid.UUID) (*entities.TechInfoApplication, error) {
	app := &entities.TechInfoApplication{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"tech_info_application_uuid",
			"company_uuid",
			"name",
			"status",
			"type",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).
		From(applicationsTable).
		Where(sq.Eq{"tech_info_application_uuid": applicationUuid}).
		RunWith(s.db).
		QueryRowContext(ctx).Scan(
		&app.TechInfoApplicationUuid,
		&app.CompanyUuid,
		&app.Name,
		&app.Status,
		&app.Type,
		&app.CreatedAt,
		&app.UpdatedAt,
		&app.CreatedBy,
		&app.UpdatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *sqlRepository) GetAllApplications(ctx context.Context, companyUuid *uuid.UUID, keyword string) ([]*entities.TechInfoApplication, error) {
	var query sq.SelectBuilder
	var sqlStr string
	var args []interface{}
	var err error

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"tech_info_applications.tech_info_application_uuid",
			"tech_info_applications.company_uuid",
			"tech_info_applications.name",
			"tech_info_applications.status",
			"tech_info_applications.type",
			"tech_info_applications.created_at",
			"tech_info_applications.updated_at",
			"tech_info_applications.created_by",
			"tech_info_applications.updated_by",

			"application_envs.application_env_uuid",
			"application_envs.company_uuid",
			"application_envs.tech_info_application_uuid",
			"application_envs.type",
			"application_envs.name",
			"application_envs.description",
			"application_envs.url",
			"application_envs.hosting_provider_type",
			"application_envs.hosting_provider",
			"application_envs.mfa_type",
			"application_envs.mfa",
			"application_envs.ids_ips_type",
			"application_envs.ids_ips",
			"application_envs.status",
			"application_envs.created_at",
		).
		From(applicationsTable).
		LeftJoin("public.application_envs ON tech_info_applications.tech_info_application_uuid = application_envs.tech_info_application_uuid").
		Where(sq.Eq{"tech_info_applications.company_uuid": companyUuid}).
		OrderBy("tech_info_applications.created_at desc")

	if keyword != "" {
		/*
			Add filter on following items:
				Application Name
				Environment Name
				Environment Type
				URL
				Host Provider
				MFA
				IDS/IPS
		*/
		sqlStr, args, err = query.Where(sq.Or{
			sq.ILike{"tech_info_applications.name": fmt.Sprintf("%%%s%%", keyword)},
			sq.ILike{"application_envs.hosting_provider": fmt.Sprintf("%%%s%%", keyword)},
			sq.ILike{"application_envs.name": fmt.Sprintf("%%%s%%", keyword)},
			sq.ILike{"application_envs.url": fmt.Sprintf("%%%s%%", keyword)},
			sq.ILike{"application_envs.mfa": fmt.Sprintf("%%%s%%", keyword)},
			sq.ILike{"application_envs.ids_ips": fmt.Sprintf("%%%s%%", keyword)},
			sq.ILike{"application_envs.type::text": fmt.Sprintf("%%%s%%", keyword)},
		}).ToSql()

		if err != nil {
			return nil, err
		}
	} else {
		sqlStr, args, err = query.ToSql()

		if err != nil {
			return nil, err
		}
	}

	rows, err := s.db.Query(sqlStr, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	applicationsMap := make(map[string]*entities.TechInfoApplication, 0)
	envsMap := make(map[string][]*entities.ApplicationEnvOnApplicationGetAllResponse)

	for rows.Next() {
		app := &entities.TechInfoApplication{}
		env := &entities.ApplicationEnvOnApplicationGetAllResponse{}

		err = rows.Scan(
			&app.TechInfoApplicationUuid,
			&app.CompanyUuid,
			&app.Name,
			&app.Status,
			&app.Type,
			&app.CreatedAt,
			&app.UpdatedAt,
			&app.CreatedBy,
			&app.UpdatedBy,

			&env.ApplicationEnvUuid,
			&env.CompanyUuid,
			&env.TechInfoApplicationUuid,
			&env.Type,
			&env.Name,
			&env.Description,
			&env.Url,
			&env.HostingProviderType,
			&env.HostingProvider,
			&env.MfaType,
			&env.Mfa,
			&env.IDsIpsType,
			&env.IDsIps,
			&env.Status,
			&env.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		applicationsMap[app.TechInfoApplicationUuid.String()] = app

		if _, ok := envsMap[env.TechInfoApplicationUuid.String()]; !ok {
			l := make([]*entities.ApplicationEnvOnApplicationGetAllResponse, 0)
			l = append(l, env)
			envsMap[env.TechInfoApplicationUuid.String()] = l
		} else {
			envsMap[env.TechInfoApplicationUuid.String()] = append(envsMap[env.TechInfoApplicationUuid.String()], env)
		}
	}

	applications := make([]*entities.TechInfoApplication, 0)

	for k, v := range applicationsMap {
		app := v
		envs := envsMap[k]
		if len(envs) > 0 {
			sort.Slice(envs, func(i, j int) bool {
				if envs[i].CreatedAt.Time != envs[j].CreatedAt.Time {
					return envs[i].CreatedAt.Time.After(envs[j].CreatedAt.Time)
				}

				return envs[i].Name.String > envs[j].Name.String
			})
		}

		app.Envs = envs
		applications = append(applications, app)
	}

	sort.Slice(applications, func(i, j int) bool {
		if applications[i].CreatedAt.Time != applications[j].CreatedAt.Time {
			return applications[i].CreatedAt.Time.After(applications[j].CreatedAt.Time)
		}

		return applications[i].Name > applications[j].Name
	})

	return applications, nil
}

func (s *sqlRepository) UpdateApplication(ctx context.Context, application *entities.TechInfoApplication) (*entities.TechInfoApplication, error) {
	oldApp, err := s.GetApplicationById(ctx, &application.TechInfoApplicationUuid)
	if err != nil {
		return nil, err
	}

	if oldApp == nil {
		return nil, &appError.ErrNotFound{Message: "application not found"}
	}

	stmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(applicationsTable)

	if application.CompanyUuid != oldApp.CompanyUuid {
		stmt = stmt.Set("company_uuid", application.CompanyUuid)
	}

	if application.Name != oldApp.Name {
		stmt = stmt.Set("name", application.Name)
	}

	if application.Type != oldApp.Type {
		stmt = stmt.Set("type", application.Type)
	}

	application.UpdatedAt = nullable.NewNullTime(time.Now())
	stmt = stmt.Set("updated_at", application.UpdatedAt)
	stmt = stmt.Set("updated_by", application.UpdatedBy)

	stmt = stmt.Where(sq.Eq{"tech_info_application_uuid": application.TechInfoApplicationUuid})

	_, err = stmt.RunWith(s.db).ExecContext(ctx)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "application already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return nil, err
	}

	application.CreatedBy = oldApp.CreatedBy
	application.CreatedAt = oldApp.CreatedAt

	return application, nil
}

func (s *sqlRepository) DeleteApplication(ctx context.Context, applicationUuid *uuid.UUID) (*uuid.UUID, error) {
	tx, txError := s.db.Begin()
	if txError != nil {
		return nil, txError
	}

	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(applicationEnvsTable).
		Where(sq.Eq{"tech_info_application_uuid": applicationUuid}).
		RunWith(tx).ExecContext(ctx)
	if err != nil {
		rtxErr := tx.Rollback()
		if rtxErr != nil {
			return nil, rtxErr
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return &uuid.Nil, err
	}

	_, err = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(applicationsTable).
		Where(sq.Eq{"tech_info_application_uuid": applicationUuid}).
		RunWith(tx).ExecContext(ctx)
	if err != nil {
		rtxErr := tx.Rollback()
		if rtxErr != nil {
			return nil, rtxErr
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return &uuid.Nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return applicationUuid, nil
}
