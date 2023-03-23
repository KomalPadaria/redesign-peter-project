package repository

import "github.com/pkg/errors"

// ErrTechInfoWirelessNotExists error.
var ErrTechInfoWirelessNotExists = errors.New("wireless assessment not exists")

// IsTechInfoWirelessNotExistsError error.
func IsTechInfoWirelessNotExistsError(err error) bool {
	return errors.Cause(err) == ErrTechInfoWirelessNotExists
}
