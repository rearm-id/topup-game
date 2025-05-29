package layouts

type Type string

const (
	LayoutTypeStandard Type = "standard"
)

var layouts = map[Type][]string{
	LayoutTypeStandard: {
		"./views/layouts/base.html",
		"./views/layouts/footer.html",
	},
}

func GetFiles(layout Type) []string {
	return layouts[layout]
}
