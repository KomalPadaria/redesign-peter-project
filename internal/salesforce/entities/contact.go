package entities

type Contact struct {
	Email      string `json:"Email"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	Name       string `json:"Name"`
	Phone      string `json:"Phone"`
	ID         string `json:"Id"`
	PhotoUrl   string `json:"PhotoUrl"`
}
