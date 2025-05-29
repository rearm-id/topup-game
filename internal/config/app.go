package config

import (
	"log/slog"
)

func Load(app Server) *AppConfig {
	return &AppConfig{
		alpine: &Alpine{},
		logger: app.Logger(),
	}
}

type AppConfig struct {
	alpine *Alpine
	logger *slog.Logger
}

func (a *AppConfig) Logger() *slog.Logger {
	return a.logger
}

type Server interface {
	Logger() *slog.Logger
}
