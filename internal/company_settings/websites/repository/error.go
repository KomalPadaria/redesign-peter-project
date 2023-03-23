package repository

import "github.com/pkg/errors"

// ErrWebsiteAlreadyExists error.
var ErrWebsiteAlreadyExists = errors.New("website already exists")

// IsWebsiteAlreadyExistsError error.
func IsWebsiteAlreadyExistsError(err error) bool {
	return errors.Cause(err) == ErrWebsiteAlreadyExists
}

// ErrWebsiteNotExists error.
var ErrWebsiteNotExists = errors.New("website not exists")

// IsWebsiteNotExistsError error.
func IsWebsiteNotExistsError(err error) bool {
	return errors.Cause(err) == ErrWebsiteNotExists
}
