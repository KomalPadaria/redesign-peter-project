package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/meetings/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/opentracing/opentracing-go/log"
)

type Service interface {
	GetMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.Meetings, error)
	GetCompanyMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetCompanyMeetingsResponse, error)
	CalendlyWebhook(ctx context.Context, data *entities.WebhookData) error
	CreateMeetingFromCalendly(ctx context.Context) error
}

type service struct {
	repo             repository.Repository
	calendlyClient   calendly.Client
	onboardingClient onboarding.Client
}

func (s *service) CreateMeetingFromCalendly(ctx context.Context) error {
	userInfo, err := s.calendlyClient.GetUserInfo(ctx)
	if err != nil {
		return err
	}

	eventTypes, err := s.calendlyClient.GetEventTypes(ctx, userInfo.URI)
	if err != nil {
		return err
	}

	for _, et := range eventTypes {
		if et.Profile.Name == "REDESIGN Trust Team" {
			uriParts := strings.Split(et.URI, "/event_types/")

			eventUUID := uriParts[1]
			eventType, err := s.calendlyClient.GetEventType(ctx, eventUUID)
			if err != nil {
				return err
			}

			eventTypeBytes, err := json.Marshal(eventType)
			if err != nil {
				return err
			}
			var md entities.MeetingData
			err = json.Unmarshal(eventTypeBytes, &md)
			if err != nil {
				return err
			}

			err = s.repo.CreateMeeting(ctx, &entities.Meetings{
				MeetingsUuid: uuid.MustParse(eventUUID),
				Name:         eventType.Name,
				Description:  eventType.DescriptionPlain,
				Duration:     fmt.Sprintf("%dm", eventType.Duration),
				Data:         &md,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *service) CalendlyWebhook(ctx context.Context, data *entities.WebhookData) error {
	utmContent := data.Payload.Tracking.UtmContent
	if strings.TrimSpace(utmContent) == "" {
		return errors.New("utm_content is not set. please set utm_content ex: company_uuid:<company_uuid>;user_uuid:<user_uuid>")
	}

	utmContentParts := strings.Split(utmContent, ";")
	if len(utmContentParts) < 2 {
		return errors.New("utm_content is not set. please set utm_content ex: company_uuid:<company_uuid>;user_uuid:<user_uuid>")
	}

	companyUuidParts := strings.Split(utmContentParts[0], ":")
	if len(companyUuidParts) < 2 {
		return errors.New("company_uuid is not set properly. please set utm_content properly. ex: company_uuid:<company_uuid>;user_uuid:<user_uuid>")
	}

	companyUuidStr := companyUuidParts[1]
	companyUuid, err := uuid.Parse(companyUuidStr)
	if err != nil {
		return errors.New("invalid company_uuid. uuid parse error")
	}

	userUuidParts := strings.Split(utmContentParts[1], ":")
	if len(userUuidParts) < 2 {
		return errors.New("user_uuid is not set properly. please set utm_content properly. ex: company_uuid:<company_uuid>;user_uuid:<user_uuid>")
	}

	userUuidStr := userUuidParts[1]
	userUuid, err := uuid.Parse(userUuidStr)
	if err != nil {
		return errors.New("invalid user_uuid. uuid parse error")
	}

	eventId := strings.Split(data.Payload.Event, "scheduled_events/")[1]
	event, err := s.calendlyClient.GetScheduledEvent(ctx, eventId)
	if err != nil {
		log.Error(err)
		return err
	}

	meetingUuidStr := strings.Split(event.EventType, "event_types/")[1]
	meetingUuid, err := uuid.Parse(meetingUuidStr)
	if err != nil {
		return err
	}

	meeting, err := s.repo.GetMeetingByUuid(ctx, &meetingUuid)
	if err != nil {
		return err
	}

	eventUuid, err := uuid.Parse(eventId)
	if err != nil {
		return err
	}

	cMeeting := &entities.CompanyMeetings{}
	cMeeting.CompanyMeetingsUuid = eventUuid
	cMeeting.CreatedBy = userUuid
	cMeeting.UpdatedBy = userUuid
	cMeeting.CompanyUuid = companyUuid
	cMeeting.Data = &entities.CompanyMeetingsData{
		WebhookPayload: &data.Payload,
		ScheduledEvent: event,
	}
	cMeeting.MeetingsUuid = meetingUuid
	cMeeting.Host = meeting.Data.Profile.Name
	cMeeting.Name = meeting.Data.Name
	cMeeting.StartAt = nullable.NewNullTime(event.StartTime)
	cMeeting.UtcStartAt = nullable.NewNullTime(event.StartTime)

	err = s.repo.CreateOrUpdateCompanyMeeting(ctx, cMeeting)
	if err != nil {
		return err
	}
	s.UpdateOnboardingStep(ctx, companyUuid, userUuid, data.Event)
	return nil
}

func (s *service) GetMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.Meetings, error) {
	return s.repo.GetMeetings(ctx, companyUUID, userUUID)
}

func (s *service) GetCompanyMeetings(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetCompanyMeetingsResponse, error) {
	return s.repo.GetCompanyMeetings(ctx, companyUUID, userUUID)

}

// Check if "Schedule Meetings" onboarding step can be updated
func (s *service) UpdateOnboardingStep(ctx context.Context, companyUuid, userUuid uuid.UUID, eventType string) {
	meetingsToSkipChecking := []string{"ad hoc"}
	if eventType == "invitee.created" {

		allMeetings, err := s.repo.GetMeetings(ctx, &companyUuid, &userUuid)
		if err != nil {
			log.Error(err)
		}

		// remove meetings to skip from all available meetings
		var filteredMeetings []*entities.Meetings
		for _, e := range meetingsToSkipChecking {
			for i, v := range allMeetings {
				if strings.Contains(strings.ToLower(v.Name), e) {
					filteredMeetings = append(allMeetings[:i], allMeetings[i+1:]...)
				}
			}
		}

		companyMeetings, err := s.repo.GetCompanyMeetings(ctx, &companyUuid, &userUuid)
		if err != nil {
			log.Error(err)
		}
		totalMeetings := len(filteredMeetings)
		if totalMeetings == 0 {
			return
		}
		// count of meetings that are scheduled
		isActiveCount := 0
		for _, meeting := range filteredMeetings {
			for _, cm := range companyMeetings {
				if meeting.MeetingsUuid == cm.MeetingsUuid {
					if cm.Data.Status.Valid && cm.Data.Status.String == "active" {
						isActiveCount += 1
					}
				}
			}
		}
		// all event types in "meetings" table have atleast one active scheduled meeting in "companyMeetings" table
		if isActiveCount >= totalMeetings {
			s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.ScheduleMeetings, companyUuid, userUuid)
		}
	}
}

func New(repo repository.Repository, calendlyClient calendly.Client, onboardingClient onboarding.Client) Service {
	return &service{
		repo:             repo,
		calendlyClient:   calendlyClient,
		onboardingClient: onboardingClient,
	}
}
