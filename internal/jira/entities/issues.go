package entities

type CreateIssueRequest struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

type CreateIssueResponse struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}
