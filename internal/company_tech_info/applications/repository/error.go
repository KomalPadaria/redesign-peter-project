package repository

import "github.com/pkg/errors"

// ErrApplicationNotExists error.
var ErrApplicationNotExists = errors.New("application not exists")

// IsApplicationNotExistsError error.
func IsApplicationNotExistsError(err error) bool {
	return errors.Cause(err) == ErrApplicationNotExists
}
