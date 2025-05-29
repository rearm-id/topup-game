package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

type HomePage struct {
	view *template.Renderer
}

func NewHomePage(componentLoader ComponentLoader) *HomePage {
	return &HomePage{
		view: componentLoader.Load(),
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
