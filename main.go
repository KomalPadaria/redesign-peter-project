// Package main is the main app
package main

import (
	"embed"
	"os"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/cmd"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/permissions"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/swagger/static"
	"github.com/opentracing/opentracing-go/log"
	"gopkg.in/yaml.v3"
)

//go:embed docs/swagger
var staticFiles embed.FS

//go:embed internal/auth/permissions/permissions.yml
var permissoinFile []byte

const exitCode = 1

func execute() error {
	return cmd.Execute()
}

func main() {
	static.StaticFiles = staticFiles

	data := make(map[string]map[string]string)
	err := yaml.Unmarshal(permissoinFile, &data)
	if err != nil {
		log.Error(err)
		os.Exit(exitCode)
	}
	permissions.Permissions = data

	if err := execute(); err != nil {
		os.Exit(exitCode)
	}
}
