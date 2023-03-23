// Package logging contains logs interceptors
package logging

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/log"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/meta"
)

type loggingRoundTripper struct {
	logger *zap.SugaredLogger

	next http.RoundTripper
}

//nolint:funlen
func (h *loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	host, method := r.URL.Hostname(), r.Method
	startTime := time.Now()
	ctx := r.Context()

	ctxLogger := h.logger.With(
		"component", "client",
		"http.host", host,
		"http.path", r.URL.Path,
		"http.method", method,
		"request_id", meta.RequestID(ctx),
		"user_agent", meta.UserAgent(ctx),
		"user_agent_origin", meta.UserAgentOrigin(ctx),
	)

	var reqFields, resFields []interface{}

	var reqPayload, resPayload string

	if r.Body != nil && r.Body != http.NoBody {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			reqFields = append(reqFields, "http.request.read_error", err)
		}

		defer r.Body.Close()

		if json.Valid(body) {
			reqPayload = string(body)
		} else {
			// Queries are sometimes passed as url.Values
			form, err := url.ParseQuery(string(body))
			if err != nil {
				reqFields = append(reqFields, "http.request.parse_query_error", err)
			}

			data := form.Get("data")
			if json.Valid([]byte(data)) {
				reqPayload = data
			}
		}

		// Restore the io.ReadCloser to its original state
		r.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	reqFields = append(reqFields, "http.request.content", log.Mask(reqPayload))

	ctxLogger.Infow("started call", reqFields...)

	response, err := h.next.RoundTrip(r)

	resFields = append(resFields, "http.time_ms", durationToMilliseconds(time.Since(startTime)))

	if response != nil {
		if response.Body != http.NoBody {
			body, ioerr := io.ReadAll(response.Body)
			if ioerr != nil {
				resFields = append(resFields, "http.response.read_error", ioerr)
			}

			if json.Valid(body) {
				resPayload = string(body)
			}

			// Restore the io.ReadCloser to its original state
			response.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		resFields = append(resFields, "http.response.content", log.Mask(resPayload), "http.code", response.StatusCode)
	}

	if err != nil {
		resFields = append(resFields, "error", err)
		ctxLogger.Errorw("finished call", resFields...)
	} else {
		ctxLogger.Infow("finished call", resFields...)
	}

	return response, err
}

// ClientRoundTripper returns a new  client round tripper that adds logger to the context.
func ClientRoundTripper(logger *zap.SugaredLogger, next http.RoundTripper) http.RoundTripper {
	return &loggingRoundTripper{logger, next}
}
