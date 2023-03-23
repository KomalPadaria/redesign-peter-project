package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"gorm.io/gorm"
)

const (
	meetingsTable        = "public.meetings"
	companyMeetingsTable = "public.company_meetings"
)

type sqlRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func (s *sqlRepository) CreateMeeting(ctx context.Context, meeting *entities.Meetings) error {
	result := s.gormDB.WithContext(ctx).Create(meeting)
	if result.Error != nil {
		err := result.Error
		if db.IsAlreadyExistError(err) {
			return nil
		}
		return err
	}
	return nil
}

func (s *sqlRepository) GetMeetingByUuid(ctx context.Context, meetingUUID *uuid.UUID) (*entities.Meetings, error) {
	var meetings entities.Meetings

	err := s.gormDB.WithContext(ctx).Table(meetingsTable).Select(
		"meetings_uuid",
		"name",
		"description",
		"duration",
		"data",
		"to_json(company_types)",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by").
		First(&meetings, "meetings_uuid = ?", meetingUUID).Error
	if err != nil {
		return nil, err
	}

	return &meetings, nil
}

func (s *sqlRepository) CreateOrUpdateCompanyMeeting(ctx context.Context, cMeeting *entities.CompanyMeetings) error {
	now := time.Now()
	cMeeting.CreatedAt = nullable.NewNullTime(now)
	cMeeting.UpdatedAt = nullable.NewNullTime(now)

	exists := s.isCompanyMeetingExistByUuid(ctx, &cMeeting.CompanyMeetingsUuid)
	if exists {
		err := s.gormDB.Omit("created_at", "created_by", "company_meetings_uuid").Where("company_meetings_uuid = ?", cMeeting.CompanyMeetingsUuid).Updates(cMeeting).Error
		if err != nil {
			return err
		}
	} else {
		err := s.gormDB.WithContext(ctx).Create(cMeeting).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sqlRepository) GetCompanyMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetCompanyMeetingsResponse, error) {
	rows, err := s.gormDB.WithContext(ctx).Raw("select "+
		"cm.company_meetings_uuid, "+
		"m.meetings_uuid, "+
		"cm.company_uuid, "+
		"cm.start_at, "+
		"cm.utc_start_at, "+
		"m.data->>'name' as meeting_name, "+
		"m.data->'profile'->>'name' as host, "+
		"m.data->>'scheduling_url' as scheduling_url, "+
		"cm.data->'webhook_payload'->>'cancel_url' as cancel_url, "+
		"cm.data->'webhook_payload'->>'reschedule_url' as reschedule_url, "+
		"cm.data->'scheduled_event'->'location'->>'join_url' as location, "+
		"cm.data->'scheduled_event'->>'status' as status, "+
		"m.duration as duration, "+
		"CASE WHEN cm.company_meetings_uuid is null THEN false else true end as scheduled "+
		"from meetings m "+
		"right join company_meetings cm on cm.meetings_uuid = m.meetings_uuid "+
		"where cm.company_uuid = ? and cm.data->'scheduled_event'->>'status' = 'active'", companyUUID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]*entities.GetCompanyMeetingsResponse, 0)

	for rows.Next() {
		res := &entities.GetCompanyMeetingsResponse{
			Data: entities.GetCompanyMeetingsData{},
		}
		err = rows.Scan(
			&res.CompanyMeetingsUuid,
			&res.MeetingsUuid,
			&res.CompanyUuid,
			&res.StartAt,
			&res.UtcStartAt,
			&res.MeetingName,
			&res.Host,
			&res.Data.SchedulingUrl,
			&res.Data.CancelUrl,
			&res.Data.RescheduleUrl,
			&res.Data.Location,
			&res.Data.Status,
			&res.Duration,
			&res.Scheduled,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, res)
	}

	return response, nil

}

func (s *sqlRepository) GetMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.Meetings, error) {
	rows, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(
			"meetings_uuid",
			"name",
			"description",
			"duration",
			"data",
			"to_json(company_types)",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).
		From(meetingsTable).
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

	meetings := make([]*entities.Meetings, 0)

	for rows.Next() {
		meeting := &entities.Meetings{}
		err = rows.Scan(
			&meeting.MeetingsUuid,
			&meeting.Name,
			&meeting.Description,
			&meeting.Duration,
			&meeting.Data,
			&meeting.CompanyTypes,
			&meeting.CreatedAt,
			&meeting.UpdatedAt,
			&meeting.CreatedBy,
			&meeting.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}

		meetings = append(meetings, meeting)
	}

	return meetings, nil
}

func (s *sqlRepository) isCompanyMeetingExistByUuid(ctx context.Context, cMeetingUUID *uuid.UUID) bool {
	var cMeeting entities.CompanyMeetings

	result := s.gormDB.WithContext(ctx).Table(companyMeetingsTable).
		First(&cMeeting, "company_meetings_uuid = ?", cMeetingUUID)

	return result.RowsAffected != 0
}
