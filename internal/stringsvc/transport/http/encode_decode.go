// Package http for conx.
package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/stringsvc/entities"

	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &entities.UppercaseRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return req, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &entities.CountRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	resp := &httpTransport.Response{Error: &httpTransport.Error{}, Data: r}

	w.Header().Set(contentTypeHeader, jsonContentType)

	return encodeJSONToWriter(w, resp)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set(contentTypeHeader, jsonContentType)

	errCode := http.StatusInternalServerError
	errMsg := "Unknown error"

	w.WriteHeader(errCode)

	resp := httpTransport.Response{
		Data:  &struct{}{},
		Error: &httpTransport.Error{Code: errCode, Message: errMsg},
	}

	_ = encodeJSONToWriter(w, resp) // nolint: errcheck
}

func encodeJSONToWriter(w io.Writer, message interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	return encoder.Encode(message)
}
