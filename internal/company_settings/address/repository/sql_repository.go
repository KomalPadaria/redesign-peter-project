package repository

import (
	"context"
	"database/sql"
	"sort"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/address/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	tErrors "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	companyAddressTable    = "public.company_address"
	companyFacilitiesTable = "public.company_facilities"
)

type sqlRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func (s *sqlRepository) GetFacilitiesByAddress(ctx context.Context, addressUuid *uuid.UUID) ([]*CompanyFacility, error) {
	var facilities []*CompanyFacility

	err := s.gormDB.WithContext(ctx).Model(&CompanyFacility{}).
		Where("company_address_uuid = ?", addressUuid).
		Find(&facilities).Error
	if err != nil {
		return nil, err
	}

	return facilities, nil
}

func (s *sqlRepository) GetFacilities(ctx context.Context, companyUUID *uuid.UUID, query, status string) ([]*CompanyFacility, error) {
	var facilities []*CompanyFacility

	tx := s.gormDB.WithContext(ctx).Model(&CompanyFacility{}).
		Preload("CompanyAddress")

	tx.Where("company_uuid = ?", companyUUID)

	if status != "" {
		tx.Where("status = ?", status)
	}

	if query != "" {
		tx.Where("name ILIKE ? OR type::text ILIKE ? ", query, query)
	}

	err := tx.Order("created_at desc").Find(&facilities).Error
	if err != nil {
		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum company_facility_status") {
				return nil, errors.WithMessage(err, "invalid input value for status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}
		return nil, err
	}

	return facilities, nil
}

func (s *sqlRepository) UpdateCompanyAddressPatch(ctx context.Context, userUUID, addressUuid *uuid.UUID, req *entities.UpdateCompanyAddressPatchRequestBody) error {
	result := s.gormDB.WithContext(ctx).Model(&CompanyAddress{}).
		WithContext(ctx).
		Where("company_address_uuid = ?", addressUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     req.Status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum company_address_status") {
				return errors.WithMessage(result.Error, "invalid input value for status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}

		return result.Error
	}

	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		return &tErrors.ErrNotFound{Message: "company address not found"}
	}

	return nil
}

func (s *sqlRepository) UpdateCompanyFacilityStatusByAddress(ctx context.Context, userUUID, addressUuid *uuid.UUID, status string) error {
	result := s.gormDB.WithContext(ctx).Model(&CompanyFacility{}).
		WithContext(ctx).
		Where("company_address_uuid = ?", addressUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"status":     status,
			})
	if result.Error != nil {
		if db.IsInvalidValueError(result.Error) {
			if strings.Contains(result.Error.Error(), "invalid input value for enum company_facility_status") {
				return errors.WithMessage(result.Error, "invalid input value for status")
			}

			return errors.WithMessage(result.Error, "invalid value")
		}

		return result.Error
	}

	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		return &tErrors.ErrNotFound{Message: "company address not found"}
	}

	return nil
}

func (s *sqlRepository) GetAddressById(ctx context.Context, addressUUID *uuid.UUID) (*entities.CompanyAddress, error) {
	address := &entities.CompanyAddress{}

	rows, err := s.gormDB.WithContext(ctx).Model(&CompanyAddress{}).Select(
		"company_address.company_address_uuid",
		"company_address.company_uuid",
		"company_address.zip",
		"company_address.country",
		"company_address.state",
		"company_address.city",
		"company_address.address1",
		"company_address.address2",
		"company_address.status",
		"company_address.created_at",
		"company_address.updated_at",
		"company_address.created_by",
		"company_address.updated_by",

		"company_facilities.company_facility_uuid",
		"company_facilities.company_address_uuid",
		"company_facilities.name",
		"company_facilities.type",
		"company_facilities.status",
	).Omit("company_facilities.name").Joins("left join public.company_facilities ON company_facilities.company_address_uuid = company_address.company_address_uuid").
		Where("company_address.company_address_uuid", addressUUID).
		Order("company_address.created_at desc").Rows()

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	facilities := make([]*entities.CompanyFacility, 0)

	for rows.Next() {
		facility := &entities.CompanyFacility{}

		err = rows.Scan(
			&address.CompanyAddressUuid,
			&address.CompanyUuid,
			&address.Zip,
			&address.Country,
			&address.State,
			&address.City,
			&address.Address1,
			&address.Address2,
			&address.Status,
			&address.CreatedAt,
			&address.UpdatedAt,
			&address.CreatedBy,
			&address.UpdatedBy,

			&facility.CompanyFacilityUuid,
			&facility.CompanyAddressUuid,
			&facility.Name,
			&facility.Type,
			&facility.Status,
		)
		if err != nil {
			return nil, err
		}

		facilities = append(facilities, facility)
	}

	address.Facilities = facilities

	return address, nil
}

func (s *sqlRepository) UpdateAddress(ctx context.Context, address *entities.CompanyAddress) (*entities.CompanyAddress, error) {
	oldAddress, err := s.GetAddressById(ctx, &address.CompanyAddressUuid)
	if err != nil {
		return nil, err
	}

	if oldAddress == nil {
		return nil, &tErrors.ErrNotFound{Message: "company address not found"}
	}

	address.UpdatedAt = nullable.NewNullTime(time.Now())

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(companyAddressTable).
		SetMap(map[string]interface{}{
			"zip":        address.Zip,
			"country":    address.Country,
			"state":      address.State,
			"city":       address.City,
			"address1":   address.Address1,
			"address2":   address.Address2,
			"status":     address.Status,
			"updated_at": address.UpdatedAt,
			"updated_by": address.UpdatedBy,
		}).
		Where(sq.Eq{"company_address_uuid": address.CompanyAddressUuid}).
		ToSql()
	if err != nil {
		return nil, err
	}

	tx, tErr := s.db.Begin()
	if tErr != nil {
		return nil, tErr
	}

	_, err = tx.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		trErr := tx.Rollback()
		if trErr != nil {
			return nil, trErr
		}

		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "address already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum company_address_status") {
				return nil, errors.WithMessage(err, "invalid input value for status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	// delete old facilities
	for _, oldFacility := range oldAddress.Facilities {
		found := false
		for _, newFacility := range address.Facilities {
			if newFacility.Name == oldFacility.Name && newFacility.Type == oldFacility.Type {
				found = true
			}
		}
		// if facility doesn't exists in request then delete it
		if !found {
			_, err = s.deleteFacility(ctx, tx, &oldFacility.CompanyFacilityUuid)
			if err != nil {
				trErr := tx.Rollback()
				if trErr != nil {
					return nil, trErr
				}
				return nil, err
			}
		}
	}

	// add new facility
	for _, newFacility := range address.Facilities {
		// check if facililty already exists
		found := false
		for _, oldFacility := range oldAddress.Facilities {
			if newFacility.Name == oldFacility.Name && newFacility.Type == oldFacility.Type {
				found = true
			}
		}
		// create new facility if one doesn't exist
		if !found {
			newFacility.CompanyUuid = oldAddress.CompanyUuid
			newFacility.CompanyAddressUuid = oldAddress.CompanyAddressUuid
			_, err = s.createFacility(ctx, tx, newFacility)
			if err != nil {
				trErr := tx.Rollback()
				if trErr != nil {
					return nil, trErr
				}
				return nil, err
			}
		}
	}

	txCErr := tx.Commit()
	if txCErr != nil {
		return nil, txCErr
	}

	return address, nil
}

func (s *sqlRepository) CreateAddress(ctx context.Context, address *entities.CompanyAddress) (*entities.CompanyAddress, error) {
	address.CompanyAddressUuid = uuid.New()

	address.CreatedAt = nullable.NewNullTime(time.Now())

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(companyAddressTable).
		Columns(
			"company_address_uuid",
			"company_uuid",
			"zip",
			"country",
			"state",
			"city",
			"address1",
			"address2",
			"status",
			"created_at",
			"created_by",
		).
		Values(
			address.CompanyAddressUuid,
			address.CompanyUuid,
			address.Zip,
			address.Country,
			address.State,
			address.City,
			address.Address1,
			address.Address2,
			address.Status,
			address.CreatedAt,
			address.CreatedBy,
		).ToSql()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			return nil, rErr
		}
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "address already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum company_address_status") {
				return nil, errors.WithMessage(err, "invalid input value for status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	for _, facility := range address.Facilities {
		facility.CompanyAddressUuid = address.CompanyAddressUuid
		facility.CreatedAt = nullable.NewNullTime(time.Now())

		_, fErr := s.createFacility(ctx, tx, facility)
		if fErr != nil {
			rErr := tx.Rollback()
			if rErr != nil {
				return nil, rErr
			}
			return nil, fErr
		}
	}

	cErr := tx.Commit()
	if cErr != nil {
		return nil, cErr
	}

	return address, nil
}

func (s *sqlRepository) GetAddresses(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.CompanyAddressW, error) {
	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"company_address.company_address_uuid",
			"company_address.company_uuid",
			"company_address.zip",
			"company_address.country",
			"company_address.state",
			"company_address.city",
			"company_address.address1",
			"company_address.address2",
			"company_address.status",
			"company_address.created_at",
			"company_address.updated_at",
			"company_address.created_by",
			"company_address.updated_by",

			"company_facilities.company_facility_uuid",
			"company_facilities.company_address_uuid",
			"company_facilities.name",
			"company_facilities.type",
			"company_facilities.status",
		).
		From(companyAddressTable).
		LeftJoin("public.company_facilities ON company_facilities.company_address_uuid = company_address.company_address_uuid").
		Where(sq.Eq{"company_address.company_uuid": companyUUID}).
		OrderBy("company_address.created_at desc").OrderBy("company_facilities.created_at asc").
		RunWith(s.db).
		QueryContext(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addressesMap := make(map[string]*entities.CompanyAddressW, 0)
	facilitiesMap := make(map[string][]*entities.CompanyFacilityW)

	for rows.Next() {
		address := &entities.CompanyAddressW{}
		facility := &entities.CompanyFacilityW{}

		err = rows.Scan(
			&address.CompanyAddressUuid,
			&address.CompanyUuid,
			&address.Zip,
			&address.Country,
			&address.State,
			&address.City,
			&address.Address1,
			&address.Address2,
			&address.Status,
			&address.CreatedAt,
			&address.UpdatedAt,
			&address.CreatedBy,
			&address.UpdatedBy,

			&facility.CompanyFacilityUuid,
			&facility.CompanyAddressUuid,
			&facility.Name,
			&facility.Type,
			&facility.Status,
		)
		if err != nil {
			return nil, err
		}

		addressesMap[address.CompanyAddressUuid.String()] = address

		if _, ok := facilitiesMap[facility.CompanyAddressUuid.String()]; !ok {
			l := make([]*entities.CompanyFacilityW, 0)
			l = append(l, facility)
			facilitiesMap[facility.CompanyAddressUuid.String()] = l
		} else {
			facilitiesMap[facility.CompanyAddressUuid.String()] = append(facilitiesMap[facility.CompanyAddressUuid.String()], facility)
		}
	}

	addresses := make([]*entities.CompanyAddressW, 0)

	for k, v := range addressesMap {
		address := v
		facilities := facilitiesMap[k]
		// sort.Slice(facilities, func(i, j int) bool {
		// 	if facilities[i].CreatedAt.Time != facilities[j].CreatedAt.Time {
		// 		return facilities[i].CreatedAt.Time.After(facilities[j].CreatedAt.Time)
		// 	}

		// 	return facilities[i].Name.String > facilities[j].Name.String
		// })

		address.Facilities = facilities
		addresses = append(addresses, address)
	}

	sort.Slice(addresses, func(i, j int) bool {
		return addresses[i].CreatedAt.Time.After(addresses[j].CreatedAt.Time)
	})

	return addresses, nil
}

func (s *sqlRepository) DeleteAddress(ctx context.Context, addressUuid *uuid.UUID) (*uuid.UUID, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	//TODO handle error in tx.Commit()
	defer tx.Commit() //nolint:errcheck

	err = s.deleteFacilitiesByAddressUuid(ctx, tx, addressUuid)
	if err != nil {
		return nil, err
	}

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(companyAddressTable).
		Where(sq.Eq{"company_address_uuid": addressUuid}).ToSql()
	if err != nil {
		return nil, err
	}

	if tx == nil {
		_, err = s.db.ExecContext(ctx, sqlStr, args...)
	} else {
		_, err = tx.ExecContext(ctx, sqlStr, args...)
	}

	if err != nil {
		trErr := tx.Rollback()
		if trErr != nil {
			return nil, trErr
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return &uuid.Nil, err
	}

	return addressUuid, nil
}

func (s *sqlRepository) deleteFacilitiesByAddressUuid(ctx context.Context, tx *sql.Tx, facilityUuid *uuid.UUID) error {
	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(companyFacilitiesTable).
		Where(sq.Eq{"company_address_uuid": facilityUuid}).ToSql()
	if err != nil {
		return err
	}

	if tx == nil {
		_, err = s.db.ExecContext(ctx, sqlStr, args...)
	} else {
		_, err = tx.ExecContext(ctx, sqlStr, args...)
	}

	if err != nil {
		if tx != nil {
			trErr := tx.Rollback()
			if trErr != nil {
				return trErr
			}
		}
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}
		return err
	}

	return nil
}

func (s *sqlRepository) deleteFacility(ctx context.Context, tx *sql.Tx, facilityUuid *uuid.UUID) (*uuid.UUID, error) {
	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(companyFacilitiesTable).
		Where(sq.Eq{"company_facility_uuid": facilityUuid}).ToSql()
	if err != nil {
		return nil, err
	}

	if tx == nil {
		_, err = s.db.ExecContext(ctx, sqlStr, args...)
	} else {
		_, err = tx.ExecContext(ctx, sqlStr, args...)
	}

	if err != nil {
		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}
		return &uuid.Nil, err
	}

	return facilityUuid, nil
}

func (s *sqlRepository) createFacility(ctx context.Context, tx *sql.Tx, facility *entities.CompanyFacility) (*entities.CompanyFacility, error) {
	facility.CompanyFacilityUuid = uuid.New()

	facility.CreatedAt = nullable.NewNullTime(time.Now())

	if facility.Type.Valid && facility.Type.String != "Others" {
		facility.Name = facility.Type
	}

	sqlStr, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(companyFacilitiesTable).
		Columns(
			"company_facility_uuid",
			"company_address_uuid",
			"company_uuid",
			"name",
			"type",
			"status",
			"created_at",
			"created_by",
		).
		Values(
			facility.CompanyFacilityUuid,
			facility.CompanyAddressUuid,
			facility.CompanyUuid,
			facility.Name,
			facility.Type,
			facility.Status,
			facility.CreatedAt,
			facility.CreatedBy,
		).ToSql()
	if err != nil {
		return nil, err
	}

	if tx == nil {
		_, err = s.db.ExecContext(ctx, sqlStr, args...)
	} else {
		_, err = tx.ExecContext(ctx, sqlStr, args...)
	}

	if err != nil {
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "facility already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum company_facility_status") {
				return nil, errors.WithMessage(err, "invalid input value for facility status")
			} else if strings.Contains(err.Error(), "invalid input value for enum company_facility_types") {
				return nil, errors.WithMessage(err, "invalid input value for facility type")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return facility, nil
}
