package entities

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type WebhookData struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	// Either "invitee.scheduled" or "invitee.canceled"
	Event   string         `json:"event"`
	Payload WebhookPayload `json:"payload"`
}

type Tracking struct {
	UtmCampaign    string    `json:"utm_campaign"`
	UtmSource      string    `json:"utm_source"`
	UtmMedium      string    `json:"utm_medium"`
	UtmContent     string    `json:"utm_content"`
	UtmTerm        string    `json:"utm_term"`
	SalesforceUUID uuid.UUID `json:"salesforce_uuid"`
}

type WebhookPayload struct {
	CancelURL             string        `json:"cancel_url"`
	CreatedAt             time.Time     `json:"created_at"`
	Email                 string        `json:"email"`
	Event                 string        `json:"event"`
	FirstName             string        `json:"first_name"`
	LastName              string        `json:"last_name"`
	Name                  string        `json:"name"`
	NewInvitee            interface{}   `json:"new_invitee"`
	NoShow                interface{}   `json:"no_show"`
	OldInvitee            interface{}   `json:"old_invitee"`
	Payment               interface{}   `json:"payment"`
	QuestionsAndAnswers   []interface{} `json:"questions_and_answers"`
	Reconfirmation        interface{}   `json:"reconfirmation"`
	RescheduleURL         string        `json:"reschedule_url"`
	Rescheduled           bool          `json:"rescheduled"`
	RoutingFormSubmission interface{}   `json:"routing_form_submission"`
	Status                string        `json:"status"`
	TextReminderNumber    interface{}   `json:"text_reminder_number"`
	Timezone              string        `json:"timezone"`
	Tracking              Tracking      `json:"tracking"`
	UpdatedAt             time.Time     `json:"updated_at"`
	URI                   string        `json:"uri"`
}

type MeetingData struct {
	Active           bool              `json:"active"`
	AdminManaged     bool              `json:"admin_managed"`
	BookingMethod    string            `json:"booking_method"`
	Color            string            `json:"color"`
	CreatedAt        time.Time         `json:"created_at"`
	CustomQuestions  []CustomQuestions `json:"custom_questions"`
	DeletedAt        interface{}       `json:"deleted_at"`
	DescriptionHTML  string            `json:"description_html"`
	DescriptionPlain string            `json:"description_plain"`
	Duration         int               `json:"duration"`
	InternalNote     interface{}       `json:"internal_note"`
	Kind             string            `json:"kind"`
	KindDescription  string            `json:"kind_description"`
	Name             string            `json:"name"`
	PoolingType      interface{}       `json:"pooling_type"`
	Profile          Profile           `json:"profile"`
	SchedulingURL    string            `json:"scheduling_url"`
	Secret           bool              `json:"secret"`
	Slug             string            `json:"slug"`
	Type             string            `json:"type"`
	UpdatedAt        time.Time         `json:"updated_at"`
	URI              string            `json:"uri"`
}

// Value simply returns the JSON-encoded representation of the struct.
func (m MeetingData) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan makes the Onboarding map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (m *MeetingData) Scan(value interface{}) error {
	if value != nil {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("type assertion to []byte failed")
		}
		return json.Unmarshal(b, &m)
	}
	return nil
}

func (m *MeetingData) MarshalJSON() ([]byte, error) {
	o := &struct {
		URL string `json:"scheduling_url"`
	}{
		URL: m.SchedulingURL,
	}

	return json.Marshal(o)
}

type CustomQuestions struct {
	AnswerChoices []interface{} `json:"answer_choices"`
	Enabled       bool          `json:"enabled"`
	IncludeOther  bool          `json:"include_other"`
	Name          string        `json:"name"`
	Position      int           `json:"position"`
	Required      bool          `json:"required"`
	Type          string        `json:"type"`
}

type Profile struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Type  string `json:"type"`
}
