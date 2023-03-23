package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	stringsInternal "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/strings"
	"github.com/pkg/errors"
)

const (
	userTable        = "public.users"
	companyUserTable = "public.company_users"
)

type sqlRepository struct {
	db *sql.DB
}

func (r *sqlRepository) GetCompanyUsers(ctx context.Context, companyUUID, userUUID uuid.UUID) ([]*entities.User, error) {
	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"users.user_uuid",
			"users.last_name",
			"users.first_name",
			"users.email",
			"users.phone",
			"users.created_at",
			"company_users.role",
			"company_users.status",
		).
		From(userTable).
		Join("company_users using(user_uuid)").
		Where(sq.Eq{"company_users.company_uuid": companyUUID}).
		OrderBy("users.created_at desc").
		RunWith(r.db).
		QueryContext(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)

	for rows.Next() {
		user := &entities.User{}
		err = rows.Scan(
			&user.UserUuid,
			&user.LastName,
			&user.FirstName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.CompanyRole,
			&user.CompanyStatus,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *sqlRepository) UpdateCurrentCompany(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) error {
	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(userTable).
		Set("current_company_uuid", companyUUID).
		Where(sq.Eq{"user_uuid": userUUID}).
		RunWith(r.db).ExecContext(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *sqlRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	return r.findByColumnNameAndColumnValue(ctx, "username", username)
}

func (r *sqlRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	return r.findByColumnNameAndColumnValue(ctx, "email", email)
}

func (r *sqlRepository) ActivateOrDeactivateByEmail(ctx context.Context, email string, isActive bool) (*entities.User, error) {
	user := &entities.User{}

	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(userTable).
		Set("is_active", isActive).
		Where(sq.Eq{"email": email}).
		RunWith(r.db).ExecContext(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *sqlRepository) LinkCompanyAndUser(ctx context.Context, companyUser *entities.CompanyUser) (*entities.CompanyUser, error) {
	companyUser.CreatedAt = nullable.NewNullTime(time.Now())

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(companyUserTable).
		Columns(
			"company_uuid",
			"user_uuid",
			"role",
			"status",
			"created_at",
		).
		Values(
			companyUser.CompanyUuid,
			companyUser.UserUuid,
			companyUser.Role,
			companyUser.Status,
			companyUser.CreatedAt,
		).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, ErrCompanyUserAlreadyExists
		}
		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum company_user_role") {
				return nil, errors.WithMessage(err, "invalid input value for role")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return companyUser, nil
}

func (r *sqlRepository) FindByExternalId(ctx context.Context, externalId string) (*entities.User, error) {
	return r.findByColumnNameAndColumnValue(ctx, "external_id", externalId)
}

func (r *sqlRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	userUUID := uuid.New()

	var currentCompanyUUID *string
	if user.CurrentCompanyUuid != uuid.Nil {
		currentCompanyUUID = stringsInternal.PointerString(user.CurrentCompanyUuid.String())
	}

	user.CreatedAt = nullable.NewNullTime(time.Now())
	userGroup := user.Group
	if userGroup == "" {
		userGroup = "customer"
	}

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(userTable).
		Columns(
			"user_uuid",
			"current_company_uuid",
			"username",
			"first_name",
			"last_name",
			"email",
			"phone",
			"is_first_login",
			"external_id",
			"created_at",
			"user_group",
		).
		Values(
			userUUID,
			currentCompanyUUID,
			user.Username,
			user.FirstName,
			user.LastName,
			user.Email,
			user.Phone,
			true,
			user.ExternalID,
			user.CreatedAt,
			userGroup,
		).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "user already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint with companies table")
		}

		return nil, err
	}

	user.UserUuid = userUUID
	if currentCompanyUUID != nil && *currentCompanyUUID != "" {
		u, err := uuid.Parse(*currentCompanyUUID)
		if err != nil {
			return nil, err
		}
		user.CurrentCompanyUuid = u

	}

	return user, nil
}

func (r *sqlRepository) UpdateCompanyUserLink(ctx context.Context, cu *entities.CompanyUser) (*entities.CompanyUser, error) {
	oldCompanyUser, err := r.getCompanyUserLink(ctx, &cu.CompanyUuid, &cu.UserUuid)
	if err != nil {
		return nil, err
	}

	if oldCompanyUser == nil {
		return nil, err
	}

	stmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(companyUserTable)

	if (cu.Role != "") && (cu.Role != oldCompanyUser.Role) {
		stmt = stmt.Set("role", cu.Role)
	}

	if (cu.Status != "") && (cu.Status != oldCompanyUser.Status) {
		stmt = stmt.Set("status", strings.ToUpper(cu.Status))
	}

	cu.UpdatedAt = nullable.NewNullTime(time.Now())
	stmt = stmt.Set("updated_at", cu.UpdatedAt)
	stmt = stmt.Set("updated_by", cu.UpdatedBy)

	stmt = stmt.Where(sq.And{sq.Eq{"company_uuid": cu.CompanyUuid}, sq.Eq{"user_uuid": cu.UserUuid}})

	_, err = stmt.RunWith(r.db).ExecContext(ctx)
	if err != nil {
		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum company_users_status") {
				return nil, errors.WithMessage(err, "invalid input value for status")
			} else if strings.Contains(err.Error(), "invalid input value for enum company_user_role") {
				return nil, errors.WithMessage(err, "invalid input value for role")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	cu.CreatedAt = oldCompanyUser.CreatedAt
	cu.CreatedBy = oldCompanyUser.CreatedBy

	return cu, nil
}

func (s *sqlRepository) DeleteCompanyUserLink(ctx context.Context, companyUUID, reqUserUUID, userUUID *uuid.UUID) error {
	updatedAt := nullable.NewNullTime(time.Now())

	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(companyUserTable).
		Set("status", "INACTIVE").
		Set("updated_at", updatedAt).
		Set("updated_by", reqUserUUID).
		Where(sq.And{sq.Eq{"company_uuid": companyUUID}, sq.Eq{"user_uuid": userUUID}}).
		RunWith(s.db).ExecContext(ctx)
	if err != nil {
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}
		return err
	}

	return nil
}

func (r *sqlRepository) FindByUUID(ctx context.Context, userUUID uuid.UUID) (*entities.User, error) {
	return r.findByColumnNameAndColumnValue(ctx, "user_uuid", userUUID)
}

func (r *sqlRepository) UpdateIsFirstLoginByUserUuid(ctx context.Context, userUUID uuid.UUID, isFirstLogin bool) (*entities.User, error) {
	user := &entities.User{}

	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(userTable).
		Set("is_first_login", isFirstLogin).
		Where(sq.Eq{"user_uuid": userUUID}).
		RunWith(r.db).ExecContext(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *sqlRepository) UpdateUserCompanyStatus(ctx context.Context, userUUID uuid.UUID, status string) error {
	// only update status if not manually set to DEACTIVE
	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(companyUserTable).
		Set("status", status).
		Where(sq.Eq{"user_uuid": userUUID}).
		Where(sq.NotEq{"status": "INACTIVE"}).
		RunWith(r.db).ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqlRepository) UpdateUserDetails(ctx context.Context, userUUID uuid.UUID, details map[string]interface{}) error {
	_, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(userTable).
		SetMap(details).
		Where(sq.Eq{"user_uuid": userUUID}).
		RunWith(r.db).ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqlRepository) GetCompanyUser(ctx context.Context, userUUID, comapnyUUID uuid.UUID) (*entities.CompanyUser, error) {
	user := &entities.CompanyUser{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_uuid",
			"user_uuid",
			"role",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
			"status").
		From(companyUserTable).
		Where(sq.Eq{"user_uuid": userUUID, "company_uuid": comapnyUUID}).
		RunWith(r.db).
		QueryRowContext(ctx).Scan(
		&user.CompanyUuid,
		&user.UserUuid,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *sqlRepository) GetCompanyUserByUserId(ctx context.Context, userUUID uuid.UUID) ([]*entities.CompanyUser, error) {

	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_uuid",
			"user_uuid",
			"role",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
			"status").
		From(companyUserTable).
		Where(sq.Eq{"user_uuid": userUUID}).
		RunWith(r.db).
		QueryContext(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companyUsers := make([]*entities.CompanyUser, 0)

	for rows.Next() {
		user := &entities.CompanyUser{}
		err = rows.Scan(
			&user.CompanyUuid,
			&user.UserUuid,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.Status,
		)

		if err != nil {
			return nil, err
		}

		companyUsers = append(companyUsers, user)
	}

	return companyUsers, nil
}

func (r *sqlRepository) findByColumnNameAndColumnValue(ctx context.Context, columnName string, columnValue interface{}) (*entities.User, error) {
	user := &entities.User{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("user_uuid",
			"current_company_uuid",
			"username",
			"first_name",
			"last_name",
			"email",
			"phone",
			"is_first_login",
			"COALESCE(external_id, '')",
			"COALESCE(job, '')",
			"mfa",
			"user_group",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by").
		From(userTable).
		Where(sq.Eq{columnName: columnValue}).
		RunWith(r.db).
		QueryRowContext(ctx).Scan(
		&user.UserUuid,
		&user.CurrentCompanyUuid,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.IsFirstLogin,
		&user.ExternalID,
		&user.JobTitle,
		&user.Mfa,
		&user.Group,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *sqlRepository) getCompanyUserLink(ctx context.Context, companyUUID, userUUID *uuid.UUID) (*entities.CompanyUser, error) {
	companyUser := &entities.CompanyUser{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_uuid",
			"user_uuid",
			"role",
			"status",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).
		From(companyUserTable).
		Where(sq.And{sq.Eq{"company_uuid": companyUUID}, sq.Eq{"user_uuid": userUUID}}).
		RunWith(r.db).
		QueryRowContext(ctx).Scan(
		&companyUser.CompanyUuid,
		&companyUser.UserUuid,
		&companyUser.Role,
		&companyUser.Status,
		&companyUser.CreatedAt,
		&companyUser.UpdatedAt,
		&companyUser.CreatedBy,
		&companyUser.UpdatedBy,
	)

	if err != nil {
		if db.IsNotFoundError(err) {
			return nil, errors.WithMessage(err, "user not found")
		}
		return nil, err
	}

	return companyUser, nil
}
