package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appErr "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/entities"
)

const (
	signaturesTable        = "public.signatures"
	companySignaturesTable = "public.company_signatures"
)

type sqlRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func (s *sqlRepository) GetCompanySignature(ctx context.Context, companyUuid, signatureUuid *uuid.UUID) (*entities.CompanySignatures, error) {
	var companySignatures *entities.CompanySignatures

	result := s.gormDB.WithContext(ctx).Debug().Model(&entities.CompanySignatures{}).
		Find(&companySignatures, "company_uuid = ? AND signature_uuid = ?", companyUuid, signatureUuid)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &appErr.ErrNotFound{Message: "signed document not found"}
	}

	return companySignatures, nil
}

func (s *sqlRepository) CreateSignatures(ctx context.Context, companySignature *entities.CompanySignatures) (*entities.CompanySignatures, error) {
	now := time.Now()
	companySignature.CreatedAt = nullable.NewNullTime(now)
	companySignature.UpdatedAt = nullable.NewNullTime(now)
	companySignature.CompanySignatureUuid = uuid.New()

	err := s.gormDB.WithContext(ctx).Create(companySignature).Error
	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "company signature already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "company signature violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum signature_status") {
				return nil, errors.WithMessage(err, "invalid input value company signature status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return companySignature, nil
}

func (s *sqlRepository) UpdateStatus(ctx context.Context, userUUID, companySignatureUuid *uuid.UUID, req *entities.UpdateStatusRequestBody) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.CompanySignatures{}).
		WithContext(ctx).
		Where("company_signature_uuid = ?", companySignatureUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     req.Status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum signature_status") {
				return errors.WithMessage(result.Error, "invalid input value status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &appErr.ErrNotFound{
			Message: "company signature not found",
		}
	}
	return nil
}

func (s *sqlRepository) GetSignatures(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetSignaturesResponse, error) {
	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"signatures.signature_uuid",
			"company_signatures.company_signature_uuid",
			"signatures.name",
			"signatures.document_url",
			"company_signatures.created_at",
			"company_signatures.status").
		From(companySignaturesTable).
		RightJoin("public.signatures ON signatures.signature_uuid = company_signatures.signature_uuid AND company_signatures.company_uuid = ?", companyUUID).
		OrderBy("company_signatures.created_at desc").
		RunWith(s.db).
		QueryContext(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	signatures := make([]*entities.GetSignaturesResponse, 0)

	for rows.Next() {
		signature := &entities.GetSignaturesResponse{}
		err = rows.Scan(
			&signature.SignatureUuid,
			&signature.CompanySignatureUuid,
			&signature.Name,
			&signature.DocumentUrl,
			&signature.Date,
			&signature.Status,
		)

		if err != nil {
			return nil, err
		}

		signatures = append(signatures, signature)
	}

	return signatures, nil
}
