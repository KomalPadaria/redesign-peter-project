package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
}

func (s *sqlRepository) GetTemplateByUuid(ctx context.Context, policyTemplateUuid *uuid.UUID) (*entities.PolicyTemplates, error) {
	var pt *entities.PolicyTemplates

	result := s.gormDB.WithContext(ctx).Model(&entities.PolicyTemplates{}).
		Select("policy_template_uuid", "name", "description", "document").
		Find(&pt, "policy_template_uuid = ?", policyTemplateUuid)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "policy template not found"}
	}

	return pt, nil
}

func (s *sqlRepository) GetTemplates(ctx context.Context, companyTypes []string) ([]*entities.PolicyTemplates, error) {
	var pts []*entities.PolicyTemplates

	tx := s.gormDB.WithContext(ctx).Model(&entities.PolicyTemplates{}).
		Select("policy_template_uuid", "name", "description").
		Order("created_at desc").
		Order("name")

	for _, ct := range companyTypes {
		tx.Or("? = any(industry_type)", ct)
	}

	err := tx.Find(&pts).Error
	if err != nil {
		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum industry_type") {
				return nil, errors.WithMessage(err, "invalid input value company_type")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return pts, nil
}

func (s *sqlRepository) GetPolicyHistoriesByPolicyUuid(ctx context.Context, policyUuid *uuid.UUID) ([]*entities.PolicyHistory, error) {
	var phs []*entities.PolicyHistory

	err := s.gormDB.WithContext(ctx).Model(&entities.PolicyHistory{}).
		Preload("Created").
		Order("created_at desc").
		Find(&phs, "policy_uuid = ?", policyUuid).Error
	if err != nil {
		return nil, err
	}

	return phs, nil
}

func (s *sqlRepository) SaveDocument(ctx context.Context, req *entities.SaveDocumentRequest) (*entities.PolicyHistory, error) {
	policy := &entities.Policy{
		PolicyUuid: req.PolicyUUID,
		Name:       req.SaveDocumentRequestBody.Name,
	}

	err := s.updatePolicyByUuid(ctx, &req.UserUuid, &req.PolicyUUID, policy)
	if err != nil {
		return nil, err
	}

	ph := &entities.PolicyHistory{
		PolicyHistoryUuid: uuid.New(),
		PolicyUuid:        req.PolicyUUID,
		Document:          req.SaveDocumentRequestBody.Document,
		CreatedAt:         nullable.NewNullTime(time.Now()),
		CreatedBy:         req.UserUuid,
	}

	result := s.gormDB.WithContext(ctx).Create(ph)
	if result.Error != nil {
		err := result.Error
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "policy document already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			if strings.Contains(err.Error(), "fk_policies") {
				return nil, errors.WithMessage(err, "policy not found")
			}

			if strings.Contains(err.Error(), "fk_created_by_users") {
				return nil, errors.WithMessage(err, "user not found")
			}

			if strings.Contains(err.Error(), "fk_updated_by_users") {
				return nil, errors.WithMessage(err, "user not found")
			}

			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		return nil, err
	}

	ph, err = s.getPolicyHistoryByUuid(ctx, &ph.PolicyHistoryUuid)
	if err != nil {
		return nil, err
	}

	return ph, nil
}

func (s *sqlRepository) GetPolicyDocument(ctx context.Context, companyUuid, policyUuid *uuid.UUID, version int) (*entities.PolicyHistory, error) {
	var policyHistory *entities.PolicyHistory

	tx := s.gormDB.WithContext(ctx).Model(&entities.PolicyHistory{}).
		Preload("Policy", "company_uuid = ?", companyUuid).
		Preload("Policy.StatusUpdated").
		Preload("Created")

	if version != 0 {
		tx.Where("version = ?", version)
	}

	result := tx.Order("version desc").Limit(1).
		Find(&policyHistory, "policy_uuid = ?", policyUuid)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &appError.ErrNotFound{Message: "policy document not found"}
	}

	if policyHistory.Policy.CompanyUuid == uuid.Nil {
		return nil, &appError.ErrNotFound{Message: "invalid company id"}
	}

	return policyHistory, nil
}

func (s *sqlRepository) CreatePolicy(ctx context.Context, policy *entities.Policy) (*entities.Policy, error) {
	now := time.Now()
	policy.PolicyUuid = uuid.New()
	policy.CreatedAt = nullable.NewNullTime(now)
	policy.Status = "Draft"
	policy.StatusUpdatedAt = now

	result := s.gormDB.WithContext(ctx).Create(policy)
	if result.Error != nil {
		err := result.Error
		if db.IsAlreadyExistError(err) {
			return nil, errors.WithMessage(err, "policy already exists")
		}

		if db.IsForeignKeyViolationError(err) {
			return nil, errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum policies_status") {
				return nil, errors.WithMessage(err, "invalid input value status")
			}

			return nil, errors.WithMessage(err, "invalid value")
		}

		return nil, err
	}

	return s.getPolicyByUuid(ctx, &policy.PolicyUuid)
}

func (s *sqlRepository) GetAllPolicyHistory(ctx context.Context, companyUuid *uuid.UUID, keyword string) ([]*entities.GetAllPoliciesResponse, error) {
	var rows *sql.Rows
	var err error

	policyHistoryQuery := `
		select p_policy_uuid,
			p_name,
			p_status,
			p_status_updated_at,
			p_created_at,
			ph_last_draft_date,
			coalesce(ph_version, 0),
			phu_user_uuid,
			phu_first_name,
			phu_last_name,
			phu_email,
			su_user_uuid,
			su_first_name,
			su_last_name,
			su_email
		from (
				select RANK() OVER (
						PARTITION BY ph.policy_uuid
						ORDER BY ph.version DESC
					) r_rank,
					p.policy_uuid as p_policy_uuid,
					p.name as p_name,
					p.status as p_status,
					p.status_updated_at as p_status_updated_at,
					p.created_at as p_created_at,
					ph.created_at as ph_last_draft_date,
					ph.version as ph_version,
					phu.user_uuid as phu_user_uuid,
					phu.first_name as phu_first_name,
					phu.last_name as phu_last_name,
					phu.email as phu_email,
					su.user_uuid as su_user_uuid,
					su.first_name as su_first_name,
					su.last_name as su_last_name,
					su.email as su_email
				from policies p
					left join policy_histories ph on ph.policy_uuid = p.policy_uuid
					left join users phu on phu.user_uuid = ph.created_by
					left join users su on su.user_uuid = p.status_updated_by
				where p.company_uuid = ?
				order by p.created_at desc
			) sq
		where sq.r_rank = 1`

	if keyword != "" {
		filter_search := fmt.Sprintf(
			"%s and (p_name ilike '%%%[2]v%%' or p_status::text ilike '%%%[2]v%%' or ph_version::text ilike '%%%[2]v%%' or phu_first_name ilike '%%%[2]v%%' or phu_last_name ilike '%%%[2]v%%' or phu_email ilike '%%%[2]v%%' or to_char(ph_last_draft_date, 'mm/dd/yy') ilike '%%%[2]v%%')",
			policyHistoryQuery, keyword,
		)
		log.Println(filter_search)
		rows, err = s.gormDB.WithContext(ctx).Raw(filter_search, companyUuid).Rows()
	} else {
		rows, err = s.gormDB.WithContext(ctx).Raw(policyHistoryQuery, companyUuid).Rows()
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]*entities.GetAllPoliciesResponse, 0)

	for rows.Next() {
		pol := &entities.GetAllPoliciesResponse{
			Owner:           &entities.UserInfo{},
			StatusUpdatedBy: &entities.UserInfo{},
		}
		err = rows.Scan(
			&pol.PolicyUUID,
			&pol.Name,
			&pol.Status,
			&pol.StatusUpdatedAt,
			&pol.CreatedAt,
			&pol.LastDraftDate,
			&pol.Version,
			&pol.Owner.UserUUID,
			&pol.Owner.FirstName,
			&pol.Owner.LastName,
			&pol.Owner.Email,
			&pol.StatusUpdatedBy.UserUUID,
			&pol.StatusUpdatedBy.FirstName,
			&pol.StatusUpdatedBy.LastName,
			&pol.StatusUpdatedBy.Email,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, pol)
	}

	return response, nil
}

func (s *sqlRepository) UpdatePolicyStatus(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, req *entities.UpdatePolicyDocumentStatusPatchRequestBody) error {
	now := time.Now()

	result := s.gormDB.WithContext(ctx).Model(&entities.Policy{}).Where("policy_uuid = ?", policyUuid).Updates(
		map[string]interface{}{
			"updated_by":        userUuid,
			"updated_at":        nullable.NewNullTime(now),
			"status_updated_by": userUuid,
			"status_updated_at": nullable.NewNullTime(now),
			"status":            req.Status,
		},
	)

	if result.Error != nil {
		err := result.Error

		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}

		if db.IsInvalidValueError(err) {
			if strings.Contains(err.Error(), "invalid input value for enum policies_status") {
				return errors.WithMessage(err, "invalid input value status")
			}

			return errors.WithMessage(err, "invalid value")
		}

		return err
	}

	if result.RowsAffected == 0 {
		return &appError.ErrNotFound{Message: "policy not found"}
	}

	if req.Comment != "" && req.Status == "Rejected" {
		policyHistorySubQuery := s.gormDB.Select("policy_history_uuid").Model(&entities.PolicyHistory{}).Where("policy_uuid = ?", policyUuid).Order("version desc").Limit(1)

		result = s.gormDB.WithContext(ctx).Model(&entities.PolicyHistory{}).Where("policy_history_uuid = (?)", policyHistorySubQuery).Updates(
			map[string]interface{}{
				"updated_by": userUuid,
				"updated_at": nullable.NewNullTime(time.Now()),
				"comment":    req.Comment,
			},
		)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s *sqlRepository) getPolicyByUuid(ctx context.Context, policyUuid *uuid.UUID) (*entities.Policy, error) {
	var policy entities.Policy

	err := s.gormDB.WithContext(ctx).Model(&entities.Policy{}).
		First(&policy, "policy_uuid = ?", policyUuid).Error
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (s *sqlRepository) updatePolicyByUuid(ctx context.Context, userUUID, policyUuid *uuid.UUID, policy *entities.Policy) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.Policy{}).
		Where("policy_uuid = ?", policyUuid).
		Updates(
			map[string]interface{}{
				"updated_by": userUUID,
				"updated_at": nullable.NewNullTime(time.Now()),
				"name":       policy.Name,
				"status":     "Draft",
			})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *sqlRepository) getPolicyHistoryByUuid(ctx context.Context, policyHistoryUuid *uuid.UUID) (*entities.PolicyHistory, error) {
	var ph entities.PolicyHistory

	err := s.gormDB.WithContext(ctx).Model(&entities.PolicyHistory{}).
		First(&ph, "policy_history_uuid = ?", policyHistoryUuid).Error
	if err != nil {
		return nil, err
	}

	return &ph, nil
}

func (s *sqlRepository) DeletePolicy(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID) error {

	result := s.gormDB.WithContext(ctx).Delete(&entities.Policy{}, "policy_uuid = ?", policyUuid)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &appError.ErrNotFound{Message: "policy not found"}
	}

	return nil
}

func (s *sqlRepository) GetDocument(ctx context.Context, companyUuid, userUuid, policyUuid *uuid.UUID, version int) (*entities.PolicyHistory, error) {
	pdoc, err := s.GetPolicyDocument(ctx, companyUuid, policyUuid, version)
	if err != nil {
		return nil, err
	}
	return pdoc, nil
}

func (s *sqlRepository) GetPoliciesStats(ctx context.Context, companyUuid *uuid.UUID) (*entities.GetPoliciesStatsResponse, error) {

	var result []map[string]interface{}
	stats := make(map[string]int)
	var total int

	err := s.gormDB.Debug().WithContext(ctx).Table("public.policies").Select("status, count(status)").Group("status").Find(&result, "company_uuid = ?", companyUuid)
	if err.Error != nil {
		return nil, err.Error
	}

	for _, status := range result {
		status_count := int(status["count"].(int64))
		stats[status["status"].(string)] = status_count
		// the total count includes count of "Draft" policies as well
		total += status_count
	}

	policies_stats := &entities.GetPoliciesStatsResponse{
		Submitted: stats["Submitted"],
		Approved:  stats["Approved"],
		Rejected:  stats["Rejected"],
		Total:     total,
	}
	return policies_stats, nil
}
