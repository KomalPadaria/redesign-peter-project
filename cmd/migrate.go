package cmd

import "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db/migrate/migrate"

func init() {
	rootCmd.AddCommand(migrate.Command)
}
