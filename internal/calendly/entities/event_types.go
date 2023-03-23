package entities

type EventTypesResponse struct {
	EventTypeCollection []EventType `json:"collection"`
	Pagination          Pagination  `json:"pagination"`
}

type Pagination struct {
	Count             int         `json:"count"`
	NextPage          interface{} `json:"next_page"`
	NextPageToken     interface{} `json:"next_page_token"`
	PreviousPage      interface{} `json:"previous_page"`
	PreviousPageToken interface{} `json:"previous_page_token"`
}
