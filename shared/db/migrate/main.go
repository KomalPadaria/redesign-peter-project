// Package main is db/migrate app
package main

import (
	"os"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db/migrate/migrate"
)

func main() {
	if err := migrate.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
