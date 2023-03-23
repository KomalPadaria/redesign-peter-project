package entities

import "time"

type SecurityTest struct {
	CampaignID           int          `json:"campaign_id"`
	PstID                int          `json:"pst_id"`
	Status               string       `json:"status"`
	Name                 string       `json:"name"`
	Groups               []Groups     `json:"groups"`
	PhishPronePercentage float64      `json:"phish_prone_percentage"`
	StartedAt            time.Time    `json:"started_at"`
	Duration             int          `json:"duration"`
	Categories           []Categories `json:"categories"`
	Template             Template     `json:"template"`
	LandingPage          LandingPage  `json:"landing_page"`
	ScheduledCount       int          `json:"scheduled_count"`
	DeliveredCount       int          `json:"delivered_count"`
	OpenedCount          int          `json:"opened_count"`
	ClickedCount         int          `json:"clicked_count"`
	RepliedCount         int          `json:"replied_count"`
	AttachmentOpenCount  int          `json:"attachment_open_count"`
	MacroEnabledCount    int          `json:"macro_enabled_count"`
	DataEnteredCount     int          `json:"data_entered_count"`
	QrCodeScannedCount   int          `json:"qr_code_scanned_count"`
	ReportedCount        int          `json:"reported_count"`
	BouncedCount         int          `json:"bounced_count"`
}
type Groups struct {
	GroupID int    `json:"group_id"`
	Name    string `json:"name"`
}
type Categories struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}
type Template struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Difficulty int    `json:"difficulty"`
	Type       string `json:"type"`
}
type LandingPage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RecipientResult struct {
	RecipientID        int       `json:"recipient_id"`
	PstID              int       `json:"pst_id"`
	User               User      `json:"user"`
	Template           Template  `json:"template"`
	ScheduledAt        time.Time `json:"scheduled_at"`
	DeliveredAt        time.Time `json:"delivered_at"`
	OpenedAt           time.Time `json:"opened_at"`
	ClickedAt          time.Time `json:"clicked_at"`
	RepliedAt          time.Time `json:"replied_at"`
	AttachmentOpenedAt time.Time `json:"attachment_opened_at"`
	MacroEnabledAt     time.Time `json:"macro_enabled_at"`
	DataEnteredAt      time.Time `json:"data_entered_at"`
	QrCodeScanned      time.Time `json:"qr_code_scanned"`
	ReportedAt         time.Time `json:"reported_at"`
	BouncedAt          time.Time `json:"bounced_at"`
	IP                 string    `json:"ip"`
	IPLocation         string    `json:"ip_location"`
	Browser            string    `json:"browser"`
	BrowserVersion     string    `json:"browser_version"`
	Os                 string    `json:"os"`
}
