package server

import (
	"os"
	"strings"

	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

type Option func(*server)

func WithMigration(automigrate bool) Option {
	return func(s *server) {
		// loosely check if it was executed using "go run"
		isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

		migratecmd.MustRegister(s.app, s.app.RootCmd, migratecmd.Config{
			// enable auto creation of migration files when making collection changes in the Dashboard
			// (the isGoRun check is to enable it only during development)
			Automigrate: isGoRun,
		})
	}
}
