package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

type HomePage struct {
	view *template.Renderer
}

func NewHomePage(registry registryLoader) *HomePage {
	return &HomePage{
		view: registry.LoadFiles(
			"./views/layouts/base.html",
			"./views/components/navigation.html",
			"./views/components/carousel.html",
			"./views/components/game-grid.html",
			"./views/home/home.html",
		),
	}
}

func (h *HomePage) Render(e *core.RequestEvent) error {

	data := map[string]any{
		"title":             "Home",
		"alpineAjaxVersion": "0.12.2",
		"alpineVersion":     "3.14.1",
	}

	html, err := h.view.Render(data)
	if err != nil {
		return e.Error(http.StatusInternalServerError, "Failed to render home page", err)
	}

	return e.HTML(http.StatusOK, html)
}
