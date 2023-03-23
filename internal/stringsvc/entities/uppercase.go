// Package entity contains service entities
package entities

// UppercaseRequest represents request
type UppercaseRequest struct {
	Input string `json:"input"`
}

// UppercaseResponse represents response
type UppercaseResponse struct {
	Output string `json:"output"`
}
