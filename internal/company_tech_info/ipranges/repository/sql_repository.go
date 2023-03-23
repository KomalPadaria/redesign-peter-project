package repository

import (
	"context"
	"strings"
	"time"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
}

func (s *sqlRepository) UpdateTechInfoIpRangeStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.TechInfoIpRanges{}).
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
			if strings.Contains(result.Error.Error(), "invalid input value for enum tech_info_ip_ranges_status") {
				return errors.WithMessage(result.Error, "invalid input value status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}
		return result.Error
	}

	return nil
}

func (s *sqlRepository) DeleteTechInfoIpRange(ctx context.Context, ipRangeUuid *uuid.UUID) (*uuid.UUID, error) {
	result := s.gormDB.WithContext(ctx).
		Delete(&entities.TechInfoIpRanges{}, "tech_info_ip_range_uuid = ?", ipRangeUuid)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "ip range not found"}
	}

	return ipRangeUuid, nil
}

func (s *sqlRepository) UpdateTechInfoIpRangePatch(ctx context.Context, userUUID, ipRangeUuid *uuid.UUID, req *entities.UpdateTechInfoIpRangePatchRequestBody) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.TechInfoIpRanges{}).
		WithContext(ctx).
		Where("tech_info_ip_range_uuid = ?", ipRangeUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     req.Status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum tech_info_ip_ranges_status") {
				return errors.WithMessage(result.Error, "invalid input value status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}
		return result.Error
	}

	return nil
}

func (s *sqlRepository) UpdateTechInfoIpRange(ctx context.Context, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error) {
	ipRange.UpdatedAt = nullable.NewNullTime(time.Now())

	updateData := map[string]interface{}{
		"company_facility_uuid": ipRange.CompanyFacilityUuid,
		"ip_address":            ipRange.IpAddress,
		"ip_size":               ipRange.IpSize,
		"is_external":           ipRange.IsExternal,
		"updated_at":            ipRange.UpdatedAt.Time,
		"updated_by":            ipRange.UpdatedBy,
	}
	result := s.gormDB.WithContext(ctx).Model(entities.TechInfoIpRanges{}).Where("tech_info_ip_range_uuid = ?", ipRange.TechInfoIpRangeUuid).Updates(&updateData)
	if result.Error != nil {
		err := result.Error

		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "ip range already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum tech_info_ip_ranges_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "ip range not found"}
	}

	return s.GetTechInfoIpRangeByUuid(ctx, &ipRange.TechInfoIpRangeUuid)
}

func (s *sqlRepository) CreateTechInfoIpRange(ctx context.Context, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error) {
	ipRange.TechInfoIpRangeUuid = uuid.New()
	ipRange.CreatedAt = nullable.NewNullTime(time.Now())
	ipRange.Status = "Active"

	result := s.gormDB.WithContext(ctx).Create(ipRange)
	if result.Error != nil {
		err := result.Error
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "ip range already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum tech_info_ip_ranges_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return s.GetTechInfoIpRangeByUuid(ctx, &ipRange.TechInfoIpRangeUuid)
}

func (s *sqlRepository) GetAllTechInfoIpRange(ctx context.Context, companyUuid *uuid.UUID) ([]*entities.TechInfoIpRanges, error) {
	var ips []*entities.TechInfoIpRanges

	err := s.gormDB.WithContext(ctx).Model(&entities.TechInfoIpRanges{}).
		Preload("CompanyFacility").
		Preload("CompanyFacility.CompanyAddress").
		Order("created_at desc").
		Find(&ips, "company_uuid = ?", companyUuid).Error
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (s *sqlRepository) GetTechInfoIpRangeByUuid(ctx context.Context, ipRangeUuid *uuid.UUID) (*entities.TechInfoIpRanges, error) {
	var ipRange entities.TechInfoIpRanges

	err := s.gormDB.WithContext(ctx).Model(&entities.TechInfoIpRanges{}).
		Preload("CompanyFacility").
		Preload("CompanyFacility.CompanyAddress").
		First(&ipRange, "tech_info_ip_range_uuid = ?", ipRangeUuid).Error
	if err != nil {
		return nil, err
	}

	return &ipRange, nil
}
