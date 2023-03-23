package entity

import "github.com/google/uuid"

type User struct {
	Uuid         uuid.UUID
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	UserGroup    string
	IsFirstLogin bool
	Company      *Company
}

type Company struct {
	Uuid         uuid.UUID
	Name         string
	UserRole     string
	Type         string
	IndustryType string
	ExternalId   string
}
