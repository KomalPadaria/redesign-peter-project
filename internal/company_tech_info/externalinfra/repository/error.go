package repository

import "github.com/pkg/errors"

// ErrTechInfoExternalInfraNotExists error.
var ErrTechInfoExternalInfraNotExists = errors.New("external infra not exists")

// IsTechInfoExternalInfraNotExistsError error.
func IsTechInfoExternalInfraNotExistsError(err error) bool {
	return errors.Cause(err) == ErrTechInfoExternalInfraNotExists
}
