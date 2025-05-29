package components

import (
	"github.com/pocketbase/pocketbase/tools/template"
	"github.com/rearmid/topup-game/internal/handlers/layouts"
)

type homePageComponent struct {
	registry RegistryLoader
}

func NewHomePageComponent(registry RegistryLoader) *homePageComponent {
	return &homePageComponent{
		registry: registry,
	}
}

func (h *homePageComponent) Load() *template.Renderer {
	layoutComponent := layouts.GetFiles(h.Layout())
	components := h.Components()

	allComponents := append(layoutComponent, components...)

	return h.registry.Load(allComponents...)
}

func (h *homePageComponent) Layout() layouts.Type {
	return layouts.LayoutTypeStandard
}

func (h *homePageComponent) Components() []string {
	return []string{
		"./views/components/navigation.html",
		"./views/components/carousel.html",
		"./views/components/game-grid.html",
		"./views/components/quick-categories.html",
		"./views/home/home.html",
	}
}
