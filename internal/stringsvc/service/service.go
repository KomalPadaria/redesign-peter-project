// Package service contains main app business logic
package service

import "strings"

// Service stores business logic
type Service interface {
	Uppercase(s string) string
	Count(s string) int
}

type service struct{}

func (s *service) Uppercase(input string) string { return strings.ToUpper(input) }

func (s *service) Count(input string) int { return len(input) }

// New service constructor
func New() Service {
	svc := &service{}

	return svc
}
