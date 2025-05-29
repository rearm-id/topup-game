package components

import "github.com/pocketbase/pocketbase/tools/template"

type RegistryLoader interface {
	Load(filenames ...string) *template.Renderer
}
