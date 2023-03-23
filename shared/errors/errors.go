package errors

import "github.com/pkg/errors"

type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

func IsNotFoundError(err error) bool {
	cause := errors.Cause(err)

	switch cause.(type) {
	case *ErrNotFound:
		return true
	default:
		return false
	}
}

// ErrBadRouting in the log.
var ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")

// IsBadRoutingError for log.
func IsBadRoutingError(err error) bool {
	return errors.Cause(err) == ErrBadRouting
}
