package components

import "github.com/pocketbase/pocketbase/tools/template"

// type ComponentLoader interface {
// 	Load() *template.Renderer
// }

type RegistryLoader interface {
	Load(filenames ...string) *template.Renderer
}
