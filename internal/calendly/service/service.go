package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/errors"
)

const (
	logPrefix = "calendly"
	baseURL   = "https://api.calendly.com"
)

type Service interface {
	GetScheduledEvent(ctx context.Context, eventID string) (*entities.ScheduledEvent, error)
	GetUserInfo(ctx context.Context) (*entities.UserMe, error)
	GetEventTypes(ctx context.Context, userURI string) ([]entities.EventType, error)
	GetEventType(ctx context.Context, eventTypeUUID string) (*entities.EventType, error)
}

type service struct {
	Config     config.Config
	httpClient *http.Client
}

func (s *service) GetUserInfo(ctx context.Context) (*entities.UserMe, error) {
	url := fmt.Sprintf("%s/users/me", baseURL)
	data, err := s.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	res := &entities.UserMeResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		log.Println(logPrefix, "json decode failed,", err)
		return nil, err
	}

	return &res.Resource, nil
}

func (s *service) GetEventTypes(ctx context.Context, userURI string) ([]entities.EventType, error) {
	url := fmt.Sprintf("%s/event_types?user=%s", baseURL, userURI)
	data, err := s.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	res := &entities.EventTypesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		log.Println(logPrefix, "json decode failed,", err)
		return nil, err
	}

	return res.EventTypeCollection, nil
}

func (s *service) GetEventType(ctx context.Context, eventTypeUUID string) (*entities.EventType, error) {
	url := fmt.Sprintf("%s/event_types/%s", baseURL, eventTypeUUID)
	data, err := s.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	res := &entities.EventTypeResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		log.Println(logPrefix, "json decode failed,", err)
		return nil, err
	}

	return &res.Resource, nil
}

func (s *service) GetScheduledEvent(ctx context.Context, eventID string) (*entities.ScheduledEvent, error) {
	url := fmt.Sprintf("%s/scheduled_events/%s", baseURL, eventID)
	data, err := s.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	res := &entities.ScheduledEventResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		log.Println(logPrefix, "json decode failed,", err)
		return nil, err
	}

	return &res.ScheduledEvent, nil
}

// httpRequest executes an HTTP request to the salesforce server and returns the response data in byte buffer.
func (s *service) httpRequest(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.Config.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		log.Println(logPrefix, "request failed,", resp.StatusCode)
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, err
		}

		newStr := buf.String()

		sfErr := errors.ParseCalendlyError(resp.StatusCode, buf.Bytes())

		log.Println(logPrefix, "Failed resp.body: ", newStr)
		return nil, sfErr
	}

	return io.ReadAll(resp.Body)
}

func New(cfg config.Config, httpClient *http.Client) Service {
	return &service{cfg, httpClient}
}
