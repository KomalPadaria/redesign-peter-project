// Package http for conx.
package encode

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	authError "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/errors"
	cognitoErr "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/errors"
	sfError "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/errors"
	httpInternal "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	httpError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors/http"
	"github.com/pkg/errors"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

func Response(_ context.Context, w http.ResponseWriter, r interface{}) error {
	resp := &httpInternal.Response{Error: &httpInternal.Error{}, Data: r}

	w.Header().Set(contentTypeHeader, jsonContentType)

	return encodeJSONToWriter(w, resp)
}

func ViewFileResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	data, err := base64.StdEncoding.DecodeString(fmt.Sprintf("%v", r))
	if err != nil {
		return err
	}

	w.Header().Set(contentTypeHeader, "application/pdf")
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	//data := []byte("this is some data stored as a byte slice in Go Lang!")

	// convert byte slice to io.Reader
	reader := bytes.NewReader(data)

	_, err = io.Copy(w, reader)

	return err
}

func DownloadFileResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	data, ok := r.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert into byte array while downloading file")
	}

	// set the default MIME type to send
	mime := http.DetectContentType(data)

	fileSize := len(string(data))

	// Generate the server headers
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment;")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	// convert byte slice to io.Reader
	reader := bytes.NewReader(data)

	_, err := io.Copy(w, reader)

	return err
}

func Error(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set(contentTypeHeader, jsonContentType)

	var errCode int
	var errMsg string

	switch {
	case db.IsAlreadyExistError(err):
		errCode = http.StatusConflict
		errMsg = strings.Split(err.Error(), ":")[0]
	case db.IsForeignKeyViolationError(err):
		errCode = http.StatusBadRequest
		errMsg = strings.Split(err.Error(), ":")[0]
	case db.IsInvalidValueError(err):
		errCode = http.StatusBadRequest
		errMsg = strings.Split(err.Error(), ":")[0]
	case db.IsNotFoundError(err):
		errCode = http.StatusNotFound
		errMsg = strings.Split(err.Error(), ":")[0]
	case httpError.IsBadRequestBodyError(err):
		errCode = http.StatusBadRequest
		if strings.Contains(err.Error(), "invalid UUID length") {
			errMsg = strings.Split(err.Error(), ":")[0]
		} else {
			errMsg = strings.Split(err.Error(), ":")[1]
		}
	case httpError.IsFileNotSupportedError(err):
		errCode = http.StatusBadRequest
		errMsg = err.Error()
	case appError.IsNotFoundError(err):
		errCode = http.StatusNotFound
		errCause := errors.Cause(err)
		errMsg = errCause.Error()
	case sfError.IsSalesforceError(err):
		sfErr := sfError.SalesforceError(err)
		errCode = sfErr.HttpCode
		errMsg = sfErr.ErrorMessage
	case cognitoErr.IsAccessError(err):
		errCode = http.StatusUnauthorized
		errCause := errors.Cause(err)
		errMsg = errCause.Error()
	case authError.IsNoPermissionError(err):
		errCode = http.StatusForbidden
		errCause := errors.Cause(err)
		errMsg = errCause.Error()
	case client.IsInvalidResponseError(err):
		e := client.InvalidResponseError(err)
		errCode = e.HTTPStatusCode
		errMsg = e.Description
	default:
		errCode = http.StatusInternalServerError
		errMsg = err.Error()
	}

	w.WriteHeader(errCode)

	resp := httpInternal.Response{
		Data:  &struct{}{},
		Error: &httpInternal.Error{Code: errCode, Message: errMsg},
	}

	_ = encodeJSONToWriter(w, resp) // nolint: errcheck
}

func encodeJSONToWriter(w io.Writer, message interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	return encoder.Encode(message)
}
