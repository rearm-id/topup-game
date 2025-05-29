package server

import (
	"log"

	_ "github.com/rearmid/topup-game/databases/migrations"
	"github.com/rearmid/topup-game/internal/handlers"
)

func Serve() {
	app := New(
		WithMigration(true),
	)

	handler := handlers.New(app)

	app.OnServe(handler.Register)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
