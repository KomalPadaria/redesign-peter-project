package errors

import (
	"github.com/pkg/errors"
)

var (
	// ErrNoPermission represents access denied error
	ErrNoPermission = errors.New("not enough permission")
)

// IsNoPermissionError checks if it's any of the above auth errors.
func IsNoPermissionError(err error) bool {
	cause := errors.Cause(err)

	switch cause {
	case ErrNoPermission:
		return true
	default:
		return false
	}
}
