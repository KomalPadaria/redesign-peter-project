// Package repository for eventlog.
package repository

import (
	"github.com/pkg/errors"
)

// ErrCompanyUserAlreadyExists in the eventlog.
var ErrCompanyUserAlreadyExists = errors.New("company_user already exists")

// IsCompanyUserAlreadyExistsError for eventlog.
func IsCompanyUserAlreadyExistsError(err error) bool {
	return errors.Cause(err) == ErrCompanyUserAlreadyExists
}

// ErrUserNotFound in the eventlog.
var ErrUserNotFound = errors.New("user(s) not found")

// IsUserNotFoundError for eventlog.
func IsUserNotFoundError(err error) bool {
	cause := errors.Cause(err)
	switch cause {
	case ErrUserNotFound:
		return true
	}
	return false
}
