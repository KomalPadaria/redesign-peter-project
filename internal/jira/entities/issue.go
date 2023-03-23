package entities

type SearchResult struct {
	Expand     string   `json:"expand"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Total      int      `json:"total"`
	Issues     []Issues `json:"issues"`
}
type AvatarUrls struct {
	Four8X48  string `json:"48x48"`
	Two4X24   string `json:"24x24"`
	One6X16   string `json:"16x16"`
	Three2X32 string `json:"32x32"`
}
type Author struct {
	Self         string     `json:"self"`
	AccountID    string     `json:"accountId"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
	AccountType  string     `json:"accountType"`
}
type Items struct {
	Field      string `json:"field"`
	Fieldtype  string `json:"fieldtype"`
	FieldID    string `json:"fieldId"`
	From       any    `json:"from"`
	FromString any    `json:"fromString"`
	To         string `json:"to"`
	ToString   string `json:"toString"`
}
type Histories struct {
	ID      string  `json:"id"`
	Author  Author  `json:"author"`
	Created string  `json:"created"`
	Items   []Items `json:"items"`
}
type Changelog struct {
	StartAt    int         `json:"startAt"`
	MaxResults int         `json:"maxResults"`
	Total      int         `json:"total"`
	Histories  []Histories `json:"histories"`
}
type Assignee struct {
	Self         string     `json:"self"`
	AccountID    string     `json:"accountId"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
	AccountType  string     `json:"accountType"`
}
type UpdateAuthor struct {
	Self         string     `json:"self"`
	AccountID    string     `json:"accountId"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
	AccountType  string     `json:"accountType"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ContentList struct {
	Type    string    `json:"type"`
	Content []Content `json:"content"`
}

type WorklogsComment struct {
	Version int           `json:"version"`
	Type    string        `json:"type"`
	Content []ContentList `json:"content"`
}

type Worklogs struct {
	Self             string          `json:"self"`
	Author           Author          `json:"author"`
	UpdateAuthor     UpdateAuthor    `json:"updateAuthor"`
	Comment          WorklogsComment `json:"comment"`
	Created          string          `json:"created"`
	Updated          string          `json:"updated"`
	Started          string          `json:"started"`
	TimeSpent        string          `json:"timeSpent"`
	TimeSpentSeconds int             `json:"timeSpentSeconds"`
	ID               string          `json:"id"`
	IssueID          string          `json:"issueId"`
}
type Worklog struct {
	StartAt    int        `json:"startAt"`
	MaxResults int        `json:"maxResults"`
	Total      int        `json:"total"`
	Worklogs   []Worklogs `json:"worklogs"`
}
type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}
type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}
type Fields struct {
	Summary  string   `json:"summary"`
	Assignee Assignee `json:"assignee"`
	Worklog  Worklog  `json:"worklog"`
	Status   Status   `json:"status"`
}
type Issues struct {
	Expand    string    `json:"expand"`
	ID        string    `json:"id"`
	Self      string    `json:"self"`
	Key       string    `json:"key"`
	Changelog Changelog `json:"changelog"`
	Fields    Fields    `json:"fields"`
}
