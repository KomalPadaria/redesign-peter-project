package entities

import "github.com/google/uuid"

type ListTopRemediationRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	Severity    string
	Top         int
}

type Remediation struct {
	Severity          string            `json:"severity"`
	Instances         int               `json:"instances"`
	Source            string            `json:"source"`
	IssueName         string            `json:"issue_name"`
	Recommendation    string            `json:"recommendation"`
	RemediationDetail RemediationDetail `json:"remediation_detail"`
}

type RemediationDetail struct {
	RemediationName  string   `json:"remediation_name"`
	RemediationLink  string   `json:"remediation_link"`
	IssueDescription string   `json:"issue_description"`
	Recommendations  []string `json:"recommendations"`
}
