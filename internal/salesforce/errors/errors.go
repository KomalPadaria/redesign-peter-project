package errors

import (
	"encoding/json"
	"fmt"
)

const logPrefix = "salesforce"

type jsonError []struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type ErrSalesforceError struct {
	Message      string
	HttpCode     int
	ErrorCode    string
	ErrorMessage string
}

func (err ErrSalesforceError) Error() string {
	return err.Message
}

func (err ErrSalesforceError) Cause() string {
	return "salesforce"
}

// IsSalesforceError for client.
func IsSalesforceError(err error) bool {
	return SalesforceError(err) != nil
}

// SalesforceError from error
func SalesforceError(err error) *ErrSalesforceError {
	invalidResponseErr, ok := err.(*ErrSalesforceError)
	if !ok {
		return nil
	}

	return invalidResponseErr
}

// Need to get information out of this package.
func ParseSalesforceError(statusCode int, responseBody []byte) (err error) {
	jsonError := jsonError{}
	err = json.Unmarshal(responseBody, &jsonError)
	if err == nil {
		return &ErrSalesforceError{
			Message: fmt.Sprintf(
				logPrefix+" Error. http code: %v Error Message:  %v Error Code: %v",
				statusCode, jsonError[0].Message, jsonError[0].ErrorCode,
			),
			HttpCode:     statusCode,
			ErrorCode:    jsonError[0].ErrorCode,
			ErrorMessage: jsonError[0].Message,
		}
	}

	return &ErrSalesforceError{
		Message:  string(responseBody),
		HttpCode: statusCode,
	}
}
