// Package repository for eventlog.
package repository

import (
	"github.com/pkg/errors"
)

// ErrCompanyAlreadyExists in the eventlog.
var ErrCompanyAlreadyExists = errors.New("company already exists")

// IsCompanyAlreadyExistsError for eventlog.
func IsCompanyAlreadyExistsError(err error) bool {
	return errors.Cause(err) == ErrCompanyAlreadyExists
}
