package server

import (
	"log/slog"
	"sync"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func New(opts ...Option) *server {
	s := &server{
		app: pocketbase.NewWithConfig(pocketbase.Config{
			// HideStartBanner: true,
			DefaultDataDir: "./pb_data",
		}),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type server struct {
	app *pocketbase.PocketBase

	onServeOnce sync.Once
}

func (s *server) Start() error {
	return s.app.Start()
}

func (s *server) Logger() *slog.Logger {
	return s.app.Logger()
}

func (s *server) OnServe(fn func(e *core.ServeEvent) error) {
	s.onServeOnce.Do(func() {
		s.app.OnServe().BindFunc(fn)
	})
}
