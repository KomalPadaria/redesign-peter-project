// Package http for transport.
package http

// Response for http.
type Response struct {
	Error *Error      `json:"error"`
	Data  interface{} `json:"data"`
}

// Error for http.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
