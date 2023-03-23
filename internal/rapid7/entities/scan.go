package entities

type Scans struct {
	Links     []Links `json:"links"`
	Page      Page    `json:"page"`
	Resources []Scan  `json:"resources"`
}

type VulnerabilitiesCount struct {
	Critical int `json:"critical"`
	Moderate int `json:"moderate"`
	Severe   int `json:"severe"`
	Total    int `json:"total"`
}

type Scan struct {
	Assets          string               `json:"assets"`
	Duration        string               `json:"duration"`
	EndTime         string               `json:"endTime"`
	EngineID        string               `json:"engineId"`
	EngineName      string               `json:"engineName"`
	ID              string               `json:"id"`
	Links           []Links              `json:"links"`
	Message         string               `json:"message"`
	ScanName        string               `json:"scanName"`
	ScanType        string               `json:"scanType"`
	StartTime       string               `json:"startTime"`
	StartedBy       string               `json:"startedBy"`
	Status          string               `json:"status"`
	Vulnerabilities VulnerabilitiesCount `json:"vulnerabilities"`
}
