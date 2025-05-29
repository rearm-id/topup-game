package handlers

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func New(cfg Configurator) *handler {
	return &handler{
		cfg:      cfg,
		registry: template.NewRegistry(),
	}
}

type handler struct {
	cfg      Configurator
	registry *template.Registry
}

const (
	publicUser = "/landing"
)

func (h *handler) Register(se *core.ServeEvent) error {
	// Serve static files
	se.Router.GET("/public/{path...}", apis.Static(os.DirFS("./pb_public"), false))

	// Override the root route to redirect to our landing page
	se.Router.GET("/", func(e *core.RequestEvent) error {
		return e.Redirect(http.StatusTemporaryRedirect, publicUser)
	})

	// Set up the landing page group
	landingGroup := se.Router.Group(publicUser)

	registryLoader := &registryLoaderImpl{
		registry: h.registry,
		logger:   h.cfg.Logger(),
	}

	homePage := NewHomePage(registryLoader)
	landingGroup.GET("/", homePage.Render)

	return se.Next()
}

type registryLoader interface {
	LoadFiles(filenames ...string) *template.Renderer
}

type registryLoaderImpl struct {
	logger   *slog.Logger
	registry *template.Registry
}

func (r *registryLoaderImpl) LoadFiles(filenames ...string) *template.Renderer {
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

type Configurator interface {
	Logger() *slog.Logger
}
