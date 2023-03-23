package entities

type ErrorResponse struct {
	Title   string    `json:"title"`
	Message string    `json:"message"`
	Details []Details `json:"details"`
}
type Details struct {
	Parameter string `json:"parameter"`
	Message   string `json:"message"`
}
