package errors

import (
	"encoding/json"
	"fmt"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/calendly/entities"
	"github.com/pkg/errors"
)

const logPrefix = "calendly"

type ErrCalendlyError struct {
	Message      string
	HttpCode     int
	ErrorCode    string
	ErrorMessage string
}

func (err ErrCalendlyError) Error() string {
	return err.Message
}

func (err ErrCalendlyError) Cause() string {
	return "salesforce"
}

// IsCalendlyError for client.
func IsCalendlyError(err error) bool {
	return CalendlyError(err) != nil
}

// CalendlyError from error
func CalendlyError(err error) *ErrCalendlyError {
	cErr, ok := err.(*ErrCalendlyError)
	if !ok {
		return nil
	}

	return cErr
}

func ParseCalendlyError(statusCode int, responseBody []byte) (err error) {
	eRes := entities.ErrorResponse{}
	err = json.Unmarshal(responseBody, &eRes)
	if err != nil {
		return errors.Wrap(err, "calendly error response unmarshal error")
	}

	return &ErrCalendlyError{
		Message: fmt.Sprintf(
			logPrefix+" Error. http code: %v, Error Title: %v, Error Message: %v",
			statusCode, eRes.Title, eRes.Message,
		),
		HttpCode: statusCode,
	}
}
