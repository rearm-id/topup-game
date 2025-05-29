package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
	"github.com/rearmid/topup-game/internal/handlers/mocks"
)

type HomePage struct {
	cfg  Configurator
	view *template.Renderer
}

func NewHomePage(cfg Configurator, componentLoader ComponentLoader) *HomePage {
	return &HomePage{
		cfg:  cfg,
		view: componentLoader.Load(),
	}
}

func (h *HomePage) Render(e *core.RequestEvent) error {
	data := map[string]any{
		"title":             "Home",
		"alpineAjaxVersion": h.cfg.GetAlpine().AjaxVersion,
		"alpineVersion":     h.cfg.GetAlpine().Version,
		"categories":        mocks.Categories,
	}

	html, err := h.view.Render(data)
	if err != nil {
		return e.Error(http.StatusInternalServerError, "Failed to render home page", err)
	}

	return e.HTML(http.StatusOK, html)
}
