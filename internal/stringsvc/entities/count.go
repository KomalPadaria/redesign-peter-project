// Package entity contains service entities
package entities

// CountRequest represents request
type CountRequest struct {
	Input string `json:"input"`
}

// CountResponse represents response
type CountResponse struct {
	Count int `json:"count"`
}
