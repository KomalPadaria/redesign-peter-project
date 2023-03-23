package entities

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type DocusignWebhookData struct {
	Event             string    `json:"event"`
	APIVersion        string    `json:"apiVersion"`
	URI               string    `json:"uri"`
	RetryCount        int       `json:"retryCount"`
	ConfigurationID   int       `json:"configurationId"`
	GeneratedDateTime time.Time `json:"generatedDateTime"`
	Data              Data      `json:"data"`
}
type Sender struct {
	UserName  string `json:"userName"`
	UserID    string `json:"userId"`
	AccountID string `json:"accountId"`
	Email     string `json:"email"`
}
type TextCustomFields struct {
	FieldID  string `json:"fieldId"`
	Name     string `json:"name"`
	Show     string `json:"show"`
	Required string `json:"required"`
	Value    string `json:"value"`
}
type CustomFields struct {
	TextCustomFields []TextCustomFields `json:"textCustomFields"`
}
type SignatureInfo struct {
	SignatureName     string `json:"signatureName"`
	SignatureInitials string `json:"signatureInitials"`
	FontStyle         string `json:"fontStyle"`
}
type Signers struct {
	SignatureInfo          SignatureInfo `json:"signatureInfo"`
	CreationReason         string        `json:"creationReason"`
	IsBulkRecipient        string        `json:"isBulkRecipient"`
	RequireUploadSignature string        `json:"requireUploadSignature"`
	Name                   string        `json:"name"`
	Email                  string        `json:"email"`
	RecipientID            string        `json:"recipientId"`
	RecipientIDGUID        string        `json:"recipientIdGuid"`
	RequireIDLookup        string        `json:"requireIdLookup"`
	UserID                 string        `json:"userId"`
	ClientUserID           string        `json:"clientUserId"`
	RoutingOrder           string        `json:"routingOrder"`
	RoleName               string        `json:"roleName"`
	Status                 string        `json:"status"`
	CompletedCount         string        `json:"completedCount"`
	SignedDateTime         time.Time     `json:"signedDateTime"`
	DeliveredDateTime      time.Time     `json:"deliveredDateTime"`
	SentDateTime           time.Time     `json:"sentDateTime"`
	DeliveryMethod         string        `json:"deliveryMethod"`
	RecipientType          string        `json:"recipientType"`
}
type Recipients struct {
	Signers             []Signers `json:"signers"`
	RecipientCount      string    `json:"recipientCount"`
	CurrentRoutingOrder string    `json:"currentRoutingOrder"`
}
type EnvelopeDocuments struct {
	DocumentID            string `json:"documentId"`
	DocumentIDGUID        string `json:"documentIdGuid"`
	Name                  string `json:"name"`
	Type                  string `json:"type"`
	URI                   string `json:"uri"`
	Order                 string `json:"order"`
	Display               string `json:"display"`
	IncludeInDownload     string `json:"includeInDownload"`
	SignerMustAcknowledge string `json:"signerMustAcknowledge"`
	TemplateRequired      string `json:"templateRequired"`
	AuthoritativeCopy     string `json:"authoritativeCopy"`
	PDFBytes              string `json:"PDFBytes"`
}
type EnvelopeMetadata struct {
	AllowAdvancedCorrect string `json:"allowAdvancedCorrect"`
	EnableSignWithNotary string `json:"enableSignWithNotary"`
	AllowCorrect         string `json:"allowCorrect"`
}
type EnvelopeSummary struct {
	Status                      string              `json:"status"`
	DocumentsURI                string              `json:"documentsUri"`
	RecipientsURI               string              `json:"recipientsUri"`
	AttachmentsURI              string              `json:"attachmentsUri"`
	EnvelopeURI                 string              `json:"envelopeUri"`
	EmailSubject                string              `json:"emailSubject"`
	EnvelopeID                  string              `json:"envelopeId"`
	SigningLocation             string              `json:"signingLocation"`
	CustomFieldsURI             string              `json:"customFieldsUri"`
	NotificationURI             string              `json:"notificationUri"`
	EnableWetSign               string              `json:"enableWetSign"`
	AllowMarkup                 string              `json:"allowMarkup"`
	AllowReassign               string              `json:"allowReassign"`
	CreatedDateTime             time.Time           `json:"createdDateTime"`
	DeliveredDateTime           time.Time           `json:"deliveredDateTime"`
	InitialSentDateTime         time.Time           `json:"initialSentDateTime"`
	SentDateTime                time.Time           `json:"sentDateTime"`
	CompletedDateTime           time.Time           `json:"completedDateTime"`
	StatusChangedDateTime       time.Time           `json:"statusChangedDateTime"`
	DocumentsCombinedURI        string              `json:"documentsCombinedUri"`
	CertificateURI              string              `json:"certificateUri"`
	TemplatesURI                string              `json:"templatesUri"`
	BrandID                     string              `json:"brandId"`
	ExpireEnabled               string              `json:"expireEnabled"`
	ExpireDateTime              time.Time           `json:"expireDateTime"`
	ExpireAfter                 string              `json:"expireAfter"`
	Sender                      Sender              `json:"sender"`
	CustomFields                CustomFields        `json:"customFields"`
	Recipients                  Recipients          `json:"recipients"`
	EnvelopeDocuments           []EnvelopeDocuments `json:"envelopeDocuments"`
	PurgeState                  string              `json:"purgeState"`
	EnvelopeIDStamping          string              `json:"envelopeIdStamping"`
	Is21CFRPart11               string              `json:"is21CFRPart11"`
	SignerCanSignOnMobile       string              `json:"signerCanSignOnMobile"`
	AutoNavigation              string              `json:"autoNavigation"`
	IsSignatureProviderEnvelope string              `json:"isSignatureProviderEnvelope"`
	HasFormDataChanged          string              `json:"hasFormDataChanged"`
	AllowComments               string              `json:"allowComments"`
	HasComments                 string              `json:"hasComments"`
	AllowViewHistory            string              `json:"allowViewHistory"`
	EnvelopeMetadata            EnvelopeMetadata    `json:"envelopeMetadata"`
	EnvelopeLocation            string              `json:"envelopeLocation"`
	IsDynamicEnvelope           string              `json:"isDynamicEnvelope"`
	BurnDefaultTabData          string              `json:"burnDefaultTabData"`
}
type Data struct {
	AccountID       string          `json:"accountId"`
	UserID          string          `json:"userId"`
	EnvelopeID      string          `json:"envelopeId"`
	EnvelopeSummary EnvelopeSummary `json:"envelopeSummary"`
}

// Value simply returns the JSON-encoded representation of the struct.
func (m DocusignWebhookData) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan makes the DocusignWebhookData map implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the map.
func (m *DocusignWebhookData) Scan(value interface{}) error {
	if value != nil {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("type assertion to []byte failed")
		}
		return json.Unmarshal(b, &m)
	}
	return nil
}
