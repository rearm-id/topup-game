package server

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	_ "github.com/rearmid/topup-game/databases/migrations"
	"github.com/rearmid/topup-game/internal/config"
	"github.com/rearmid/topup-game/internal/handlers"
)

func Serve() {
	server := New(
		WithMigration(true),
	)

	cfg := config.Load(server.app)

	handler := handlers.New(cfg)

	server.OnServe(handler.Register)

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(server.app, server.app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
