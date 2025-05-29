package server

import (
	"log"

	_ "github.com/rearmid/topup-game/databases/migrations"
	"github.com/rearmid/topup-game/internal/config"
	"github.com/rearmid/topup-game/internal/handlers"
)

func Serve() {
	app := New(
		WithMigration(true),
	)

	cfg := config.Load(app)

	handler := handlers.New(cfg)

	app.OnServe(handler.Register)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
