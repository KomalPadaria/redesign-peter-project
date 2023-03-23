package entities

import (
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
)

type Penetration struct {
	Severity    string            `json:"severity"`
	Scvss       string            `json:"scvss"`
	Instance    string            `json:"instance"`
	Date        nullable.NullTime `json:"date,omitempty"`
	Issue       string            `json:"issue"`
	Remediation *Remediation      `json:"remediation"`
}

type Remediation struct {
	Category       string `json:"category"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	Risk           string `json:"risk"`
	Recommendation string `json:"recommendation"`
}

type PenetrationStats struct {
	Scans    *Scan     `json:"scan"`
	Critical *Critical `json:"critical"`
	Total    *Total    `json:"total"`
	Count    *Count    `json:"count"`
}

type Scan struct {
	Current nullable.NullTime `json:"current,omitempty"`
	Next    nullable.NullTime `json:"next,omitempty"`
}

type Critical struct {
	Current     int
	Performance int
	Stats       []Stats `json:"stats"`
}

type Stats struct {
	Date  nullable.NullTime `json:"date,omitempty"`
	Value int               `json:"value"`
}

type Count struct {
	Critical int
	High     int
	Medium   int
	Low      int
}

type Total struct {
	Current     int
	Performance int
	Stats       []Stats `json:"stats"`
}

// Request Types
type GetPenetrationTestsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}

type GetPenetrationTestsStatsRequest struct {
	CompanyUuid uuid.UUID `json:"company_uuid"`
	UserUuid    uuid.UUID `json:"user_uuid"`
}
