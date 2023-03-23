package http

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	// ErrBadRequestBody when we can't decode the request body.
	ErrBadRequestBody = errors.New("bad request body")

	// ErrBadPathParameter when there is bad or invalid path parameter.
	ErrBadPathParameter = errors.New("bad or invalid path parameter")

	ErrFileNotSupported = errors.New("This file format is not supported")
)

func IsFileNotSupportedError(err error) bool {
	return errors.Cause(err) == ErrFileNotSupported
}

func IsBadRequestBodyError(err error) bool {
	return errors.Cause(err) == ErrBadRequestBody
}

type ErrBadOrInvalidPathParameter struct {
	Field string
}

func NewErrBadOrInvalidPathParameter(invalidField string) *ErrBadOrInvalidPathParameter {
	return &ErrBadOrInvalidPathParameter{
		Field: invalidField,
	}
}

func (err *ErrBadOrInvalidPathParameter) Error() string {
	return fmt.Sprintf("bad or invalid %s", err.Field)
}

func (err *ErrBadOrInvalidPathParameter) Cause() error {
	return ErrBadPathParameter
}

func IsBadInvalidPathParameterError(err error) bool {
	return errors.Cause(err) == ErrBadPathParameter
}
