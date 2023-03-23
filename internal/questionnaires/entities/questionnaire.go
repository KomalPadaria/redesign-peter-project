package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type CategoryStats struct {
	Category string `json:"category"`
	Stats    *Stats `json:"stats"`
}

type Category struct {
	Category  string `json:"category" gorm:"column:category"`
	Total     int    `json:"total" gorm:"column:total"`
	Completed int    `json:"completed" gorm:"column:completed"`
}

type Stats struct {
	Total     int `json:"total"`
	Completed int `json:"completed"`
}

// Request Types
type GetQuestionnairesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetQuestionnairesByCategoryRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Category    string    `json:"category"`
}

type PostAnswerRequest struct {
	CompanyUuid uuid.UUID         `json:"company_uuid"`
	UserUuid    uuid.UUID         `json:"user_uuid"`
	Answers     AnswerRequestBody `json:"answer"`
}

type AnswerRequestBody []*AnswerBody

type AnswerBody struct {
	QuestionnairesUuid uuid.UUID `json:"questionnaires_uuid"`
	Answer             Answer    `json:"answer"`
}

type Answer struct {
	Options []uuid.UUID `json:"options"`
	Comment string      `json:"comment"`
}

type CompanyQuestionnaires struct {
	CompanyQuestionnairesUuid uuid.UUID         `json:"company_questionnaires_uuid" gorm:"column:company_questionnaires_uuid"`
	CompanyUuid               uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	Category                  string            `json:"category" gorm:"column:category"`
	CreatedAt                 nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                 nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy                 uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy                 uuid.UUID         `json:"updated_by" gorm:"column:updated_by"`
}

func (m *CompanyQuestionnaires) TableName() string {
	return "company_questionnaires"
}

type Questionnaires struct {
	QuestionnairesUuid uuid.UUID              `json:"questionnaires_uuid" gorm:"column:questionnaires_uuid"`
	Category           string                 `json:"category" gorm:"column:category"`
	Question           string                 `json:"question" gorm:"column:question"`
	OptionType         string                 `json:"option_type" gorm:"column:option_type"`
	CommentType        string                 `json:"comment_type" gorm:"column:comment_type"`
	Options            []QuestionnaireOptions `json:"options" gorm:"foreignKey:QuestionnairesUuid;references:QuestionnairesUuid"`
	Answer             QuestionnaireAnswers   `json:"answer" gorm:"foreignKey:QuestionnairesUuid;references:QuestionnairesUuid"`
	CreatedAt          nullable.NullTime      `json:"-" gorm:"column:created_at"`
	UpdatedAt          nullable.NullTime      `json:"-" gorm:"column:updated_at"`
	CreatedBy          uuid.UUID              `json:"-" gorm:"column:created_by"`
	UpdatedBy          uuid.UUID              `json:"-" gorm:"column:updated_by"`
}

func (m *Questionnaires) TableName() string {
	return "questionnaires"
}

type QuestionnaireOptions struct {
	QuestionnaireOptionsUuid uuid.UUID         `json:"questionnaire_options_uuid" gorm:"column:questionnaire_options_uuid"`
	QuestionnairesUuid       uuid.UUID         `json:"-" gorm:"column:questionnaires_uuid"`
	Value                    int               `json:"value" gorm:"column:value"`
	Label                    string            `json:"label" gorm:"column:label"`
	CreatedAt                nullable.NullTime `json:"-" gorm:"column:created_at"`
	UpdatedAt                nullable.NullTime `json:"-" gorm:"column:updated_at"`
	CreatedBy                uuid.UUID         `json:"-" gorm:"column:created_by"`
	UpdatedBy                uuid.UUID         `json:"-" gorm:"column:updated_by"`
}

func (m *QuestionnaireOptions) TableName() string {
	return "questionnaire_options"
}

type QuestionnaireAnswers struct {
	QuestionnaireAnswersUuid uuid.UUID                   `json:"questionnaire_answers_uuid" gorm:"column:questionnaire_answers_uuid"`
	CompanyUuid              uuid.UUID                   `json:"-" gorm:"column:company_uuid"`
	QuestionnairesUuid       uuid.UUID                   `json:"-" gorm:"column:questionnaires_uuid"`
	Options                  []QuestionnaireOptions      `json:"options" gorm:"many2many:questionnaire_answers_options;foreignKey:QuestionnaireAnswersUuid;joinForeignKey:QuestionnaireAnswersUuid;References:QuestionnaireOptionsUuid;joinReferences:QuestionnaireOptionsUuid"`
	Comment                  string                      `json:"comment" gorm:"column:comment"`
	CreatedAt                nullable.NullTime           `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                nullable.NullTime           `json:"-" gorm:"column:updated_at"`
	CreatedBy                uuid.UUID                   `json:"created_by" gorm:"column:created_by"`
	Created                  entities.User               `json:"-" gorm:"foreignKey:CreatedBy;references:UserUuid"`
	UpdatedBy                uuid.UUID                   `json:"-" gorm:"column:updated_by"`
	Files                    EvidenceFiles               `json:"files" gorm:"column:files"`
	Feedback                 QuestionnaireAnswerFeedback `json:"feedback" gorm:"column:feedback"`
}

func (q *QuestionnaireAnswers) MarshalJSON() ([]byte, error) {
	type Alias QuestionnaireAnswers

	o := &struct {
		*Alias
		QuestionnaireAnswersUuid string                 `json:"questionnaire_answers_uuid,omitempty"`
		Options                  []QuestionnaireOptions `json:"options,omitempty"`
		Comment                  string                 `json:"comment,omitempty"`
		EvidenceFiles            EvidenceFiles          `json:"files,omitempty"`
		CreatedAt                string                 `json:"created_at,omitempty"`
		CreatedBy                string                 `json:"created_by,omitempty"`
	}{
		Alias: (*Alias)(q),
	}

	if q.QuestionnaireAnswersUuid == uuid.Nil {
		o.QuestionnaireAnswersUuid = ""
		o.Options = nil
		o.Comment = ""
		o.CreatedAt = ""
		o.CreatedBy = ""
		o.Files = nil
	} else {
		o.QuestionnaireAnswersUuid = q.QuestionnaireAnswersUuid.String()
		o.Options = q.Options
		o.Comment = q.Comment
		o.CreatedAt = q.CreatedAt.Time.UTC().String()
		o.CreatedBy = fmt.Sprintf("%s %s", q.Created.FirstName, q.Created.LastName)
		o.EvidenceFiles = q.Files
	}
	return json.Marshal(o)
}

func (m *QuestionnaireAnswers) TableName() string {
	return "questionnaire_answers"
}

type CreateQuestionnaireAnswers struct {
	QuestionnaireAnswersUuid uuid.UUID         `json:"questionnaire_answers_uuid" gorm:"column:questionnaire_answers_uuid"`
	CompanyUuid              uuid.UUID         `json:"company_uuid" gorm:"column:company_uuid"`
	QuestionnairesUuid       uuid.UUID         `json:"questionnaires_uuid" gorm:"column:questionnaires_uuid"`
	Comment                  string            `json:"comment" gorm:"column:comment"`
	CreatedAt                nullable.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                nullable.NullTime `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy                uuid.UUID         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy                uuid.UUID         `json:"updated_by" gorm:"column:updated_by"`
	Files                    EvidenceFiles     `json:"files" gorm:"column:files"`
}

func (m *CreateQuestionnaireAnswers) TableName() string {
	return "questionnaire_answers"
}

type QuestionnaireAnswersOptions struct {
	QuestionnaireOptionsUuid uuid.UUID `json:"questionnaire_options_uuid" gorm:"column:questionnaire_options_uuid"`
	QuestionnaireAnswersUuid uuid.UUID `json:"questionnaire_answers_uuid" gorm:"column:questionnaire_answers_uuid"`
}

func (m *QuestionnaireAnswersOptions) TableName() string {
	return "questionnaire_answers_options"
}

type SubmitQuestionnairesRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Status      QuestionnairesStatus
}

type QuestionnairesStatus struct {
	Status string `json:"status"`
}

type AnswerWithEvidenceRequest struct {
	CompanyUuid        uuid.UUID
	UserUuid           uuid.UUID
	QuestionnairesUuid uuid.UUID
	Answer             Answer `json:"answer"`
	Files              []*multipart.FileHeader
}

type UpdateAnswerWithEvidenceRequest struct {
	CompanyUuid        uuid.UUID
	UserUuid           uuid.UUID
	QuestionnairesUuid uuid.UUID
	AnswerUuid         uuid.UUID
	Answer             Answer `json:"answer"`
	Files              []*multipart.FileHeader
}

type DownloadEvidenceRequest struct {
	CompanyUuid        uuid.UUID
	UserUuid           uuid.UUID
	QuestionnairesUuid uuid.UUID
	AnswerUuid         uuid.UUID
	FileId             string
}

type DeleteEvidenceRequest struct {
	CompanyUuid        uuid.UUID
	UserUuid           uuid.UUID
	QuestionnairesUuid uuid.UUID
	AnswerUuid         uuid.UUID
	FileId             string
}

type EngineerFeedbackRequest struct {
	CompanyUuid        uuid.UUID
	UserUuid           uuid.UUID
	QuestionnairesUuid uuid.UUID
	AnswerUuid         uuid.UUID
	Feedback           QuestionnaireAnswerFeedback
}

type QuestionnaireAnswerFeedback struct {
	Title          string `json:"title"`
	Recommendation string `json:"recommendation"`
	Status         string `json:"status"`
	Severity       string `json:"severity"`
}

// Value simply returns the JSON-encoded representation of the struct.
func (m QuestionnaireAnswerFeedback) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan makes the Onboarding map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (m *QuestionnaireAnswerFeedback) Scan(value interface{}) error {
	if value != nil {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("type assertion to []byte failed")
		}
		return json.Unmarshal(b, &m)
	}
	return nil
}

type EvidenceFile struct {
	Id       uuid.UUID `json:"id,omitempty"`
	FileName string    `json:"file_name,omitempty"`
	S3Key    string    `json:"s3key,omitempty"`
}

type EvidenceFiles []EvidenceFile

func (a EvidenceFiles) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *EvidenceFiles) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
