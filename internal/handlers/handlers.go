package handlers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/rearmid/topup-game/internal/handlers/components"
	"github.com/rearmid/topup-game/internal/handlers/registry"

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

	registryLoader := registry.NewLoader(h.registry, h.cfg.Logger())
	componentLoader := components.NewHomePageComponent(registryLoader)

	homePage := NewHomePage(componentLoader)

	landingGroup.GET("/", homePage.Render)

	return se.Next()
}

type ComponentLoader interface {
	Load() *template.Renderer
}

type Configurator interface {
	Logger() *slog.Logger
}
