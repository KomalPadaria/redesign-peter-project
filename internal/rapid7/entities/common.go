package entities

type Links struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type Page struct {
	Number         int `json:"number"`
	Size           int `json:"size"`
	TotalPages     int `json:"totalPages"`
	TotalResources int `json:"totalResources"`
}
