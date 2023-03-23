package jwt

import "github.com/pkg/errors"

var (
	ErrInvalidToken     = errors.New("token is invalid")
	ErrExpiredToken     = errors.New("token has expired")
	ErrSignatureInvalid = errors.New("signature is invalid")
	ErrBadRequest       = errors.New("bad request")
)

// IsUnauthorizedError checks if it's any of the above auth errors.
func IsUnauthorizedError(err error) bool {
	cause := errors.Cause(err)

	switch cause {
	case ErrInvalidToken, ErrExpiredToken, ErrSignatureInvalid:
		return true
	default:
		return false
	}
}

// IsBadRequestError checks if it's any of the above auth errors.
func IsBadRequestError(err error) bool {
	cause := errors.Cause(err)

	switch cause {
	case ErrBadRequest:
		return true
	default:
		return false
	}
}
