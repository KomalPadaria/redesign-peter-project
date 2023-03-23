// Package http for conx.
package http

import (
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/swagger/static"
	httpTransport "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/transport/http"
)

// RegisterTransport for http.
func RegisterTransport(
	server *httpTransport.Server,
) {
	registerSwaggerUI(server)
}

func registerSwaggerUI(server *httpTransport.Server) {

	// http.FS can be used to create a http Filesystem
	var staticFS = http.FS(static.StaticFiles)
	fs := http.FileServer(staticFS)

	sh := http.StripPrefix("/swaggerui", fs)

	server.HandleWithPathPrefix("/swaggerui", sh)
}
