package errors

import (
	"github.com/pkg/errors"
)

var (
	// ErrTokenInvalid when we can't decode the token.
	ErrTokenInvalid = errors.New("token is invalid")

	// ErrTokenExpired when we can't decode the token.
	ErrTokenExpired = errors.New("token expired")

	// ErrMissingAuthHeader when the authorization is missing.
	ErrMissingAuthHeader = errors.New("missing authorization header")

	// ErrAccessDenied represents access denied error
	ErrAccessDenied = errors.New("access denied")

	// ErrIdTokenInvalid when we can't decode the token.
	ErrIdTokenInvalid = errors.New("invalid identity token")
)

// IsIdTokenInvalidError send back true if the cause of the error is ErrIdTokenInvalid
func IsIdTokenInvalidError(err error) bool {
	return errors.Cause(err) == ErrIdTokenInvalid
}

// IsAccessError checks if it's any of the above auth errors.
func IsAccessError(err error) bool {
	cause := errors.Cause(err)

	switch cause {
	case ErrTokenInvalid, ErrTokenExpired, ErrMissingAuthHeader, ErrAccessDenied:
		return true
	default:
		return false
	}
}
