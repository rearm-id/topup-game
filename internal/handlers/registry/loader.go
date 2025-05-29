package registry

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pocketbase/pocketbase/tools/template"
)

func NewLoader(registry *template.Registry, logger *slog.Logger) *loaderImpl {
	return &loaderImpl{
		registry: registry,
		logger:   logger,
	}
}

type loaderImpl struct {
	logger   *slog.Logger
	registry *template.Registry
}

func (r *loaderImpl) Load(filenames ...string) *template.Renderer {
	// get current directory
	dir, err := os.Getwd()
	slog.Info("dir: ", "dir", dir)
	if err != nil {
		r.logger.Error("Failed to get current directory", "error", err)
		return nil
	}

	for i, filename := range filenames {
		filenames[i] = filepath.Join(dir, filename)
	}

	slog.Info("filenames: ", "filenames", filenames)

	return r.registry.LoadFiles(filenames...)
}
