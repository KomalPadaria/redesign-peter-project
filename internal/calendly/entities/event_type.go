package entities

import "time"

type EventTypeResponse struct {
	Resource EventType `json:"resource"`
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

type EventType struct {
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
	PoolingType      string            `json:"pooling_type"`
	Profile          Profile           `json:"profile"`
	SchedulingURL    string            `json:"scheduling_url"`
	Secret           bool              `json:"secret"`
	Slug             string            `json:"slug"`
	Type             string            `json:"type"`
	UpdatedAt        time.Time         `json:"updated_at"`
	URI              string            `json:"uri"`
}
