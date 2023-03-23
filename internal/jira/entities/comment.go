package entities

type Comment struct {
	Body Body `json:"body"`
}

type Marks struct {
	Type string `json:"type"`
}

type ContentInner struct {
	Type  string  `json:"type"`
	Text  string  `json:"text"`
	Marks []Marks `json:"marks,omitempty"`
}

type ContentMain struct {
	Type    string         `json:"type"`
	Content []ContentInner `json:"content"`
}

type Body struct {
	Type    string        `json:"type"`
	Version int           `json:"version"`
	Content []ContentMain `json:"content"`
}
