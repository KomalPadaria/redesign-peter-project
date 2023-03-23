package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sqlRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func (s *sqlRepository) UpdateTechInfoWirelessStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.TechInfoWireless{}).
		WithContext(ctx).
		Where("company_facility_uuid IN ?", facilityUuids).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum tech_info_wireless_status") {
				return errors.WithMessage(result.Error, "invalid input value status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}
		return result.Error
	}

	return nil
}

func (s *sqlRepository) UpdateTechInfoWirelessPatch(ctx context.Context, userUUID, wirelessUuid *uuid.UUID, req *entities.UpdateTechInfoWirelessPatchRequestBody) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.TechInfoWireless{}).
		WithContext(ctx).
		Where("tech_info_wireless_uuid = ?", wirelessUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     req.Status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum tech_info_wireless_status") {
				return errors.WithMessage(result.Error, "invalid input value status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}
		return result.Error
	}

	return nil
}

func (s *sqlRepository) CreateTechInfoWireless(ctx context.Context, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error) {
	wireless.TechInfoWirelessUuid = uuid.New()
	wireless.CreatedAt = nullable.NewNullTime(time.Now())
	wireless.Status = "Active"

	result := s.gormDB.WithContext(ctx).Create(wireless)
	if result.Error != nil {
		err := result.Error
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "tech info wireless assessment already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "No facility found")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum tech_info_wireless_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			if strings.Contains(err.Error(), "invalid input value for enum tech_info_wireless_protocol_type") {
				return nil, errors.WithMessage(err, "invalid input value protocol type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum tech_info_wireless_security_type") {
				return nil, errors.WithMessage(err, "invalid input value security type")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}
		if db.IsInvalidLength(err) {
			return nil, errors.New("SSID should be less than or equal to 50 characters")
		}

		return nil, err
	}

	return s.getWirelessByUuid(ctx, &wireless.TechInfoWirelessUuid)
}

func (s *sqlRepository) GetTechInfoWirelessById(ctx context.Context, techInfoWirelessUuid *uuid.UUID) (*entities.TechInfoWireless, error) {
	return s.getWirelessByUuid(ctx, techInfoWirelessUuid)
}

func (s *sqlRepository) GetAllTechInfoWirelesss(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.TechInfoWireless, error) {
	var wirelessNetworks []*entities.TechInfoWireless

	err := s.gormDB.WithContext(ctx).Model(&entities.TechInfoWireless{}).
		Preload("CompanyFacility").
		Preload("CompanyFacility.CompanyAddress").
		Order("created_at desc").
		Find(&wirelessNetworks, "company_uuid = ?", companyUuid).Error
	if err != nil {
		return nil, err
	}

	return wirelessNetworks, nil
}

func (s *sqlRepository) UpdateTechInfoWireless(ctx context.Context, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error) {
	wireless.UpdatedAt = nullable.NewNullTime(time.Now())
	result := s.gormDB.WithContext(ctx).Where("tech_info_wireless_uuid = ?", wireless.TechInfoWirelessUuid).Updates(wireless)
	if result.Error != nil {
		err := result.Error

		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "tech info wireless already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "No facility found")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum tech_info_wireless_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			if strings.Contains(err.Error(), "invalid input value for enum tech_info_wireless_protocol_type") {
				return nil, errors.WithMessage(err, "invalid input value protocol type")
			}

			if strings.Contains(err.Error(), "invalid input value for enum tech_info_wireless_security_type") {
				return nil, errors.WithMessage(err, "invalid input value security type")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}
		if db.IsInvalidLength(err) {
			return nil, errors.New("SSID should be less than or equal to 50 characters")
		}

		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "tech info wireless not found"}
	}

	return s.getWirelessByUuid(ctx, &wireless.TechInfoWirelessUuid)

}

func (s *sqlRepository) DeleteTechInfoWireless(ctx context.Context, techInfoWirelessUuid *uuid.UUID) (*uuid.UUID, error) {
	result := s.gormDB.WithContext(ctx).
		Delete(&entities.TechInfoWireless{}, "tech_info_wireless_uuid = ?", techInfoWirelessUuid)
	if result.Error != nil {
		if db.IsForeignKeyViolationError(result.Error) {
			return nil, errors.WithMessage(result.Error, "violates foreign key constraint")
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "tech info wireless not found"}
	}

	return techInfoWirelessUuid, nil
}

func (s *sqlRepository) getWirelessByUuid(ctx context.Context, wirelessUuid *uuid.UUID) (*entities.TechInfoWireless, error) {
	var ipRange entities.TechInfoWireless

	err := s.gormDB.WithContext(ctx).Model(&entities.TechInfoWireless{}).
		Preload("CompanyFacility").
		Preload("CompanyFacility.CompanyAddress").
		First(&ipRange, "tech_info_wireless_uuid = ?", wirelessUuid).Error
	if err != nil {
		return nil, err
	}

	return &ipRange, nil
}
