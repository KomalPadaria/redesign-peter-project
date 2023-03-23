package entities

import "time"

type ScheduledEventResponse struct {
	ScheduledEvent ScheduledEvent `json:"resource"`
}
type CalendarEvent struct {
	ExternalID string `json:"external_id"`
	Kind       string `json:"kind"`
}
type EventGuests struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}
type EventMemberships struct {
	User string `json:"user"`
}
type InviteesCounter struct {
	Active int `json:"active"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}
type Location struct {
	JoinURL string `json:"join_url"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}
type ScheduledEvent struct {
	CalendarEvent    CalendarEvent      `json:"calendar_event"`
	CreatedAt        time.Time          `json:"created_at"`
	EndTime          time.Time          `json:"end_time"`
	EventGuests      []EventGuests      `json:"event_guests"`
	EventMemberships []EventMemberships `json:"event_memberships"`
	EventType        string             `json:"event_type"`
	InviteesCounter  InviteesCounter    `json:"invitees_counter"`
	Location         Location           `json:"location"`
	Name             string             `json:"name"`
	StartTime        time.Time          `json:"start_time"`
	Status           string             `json:"status"`
	UpdatedAt        time.Time          `json:"updated_at"`
	URI              string             `json:"uri"`
}
