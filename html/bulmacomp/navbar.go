package bulmacomp

import (
	"fmt"
	g "maragu.dev/gomponents"
	gc "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func Navbar(isFixedTop bool, margin string, brand, start, end g.Node) g.Node {
	return h.Nav(gc.Classes{
		"navbar":       true,
		"is-fixed-top": isFixedTop,
	}, h.Role("navigation"), h.Aria("label", "main navigation"), h.Style(fmt.Sprintf("margin: %s;", margin)),
		h.Div(h.Class("navbar-brand"), brand),
		h.Div(h.Class("navbar-menu"),
			h.Div(h.Class("navbar-start"), start),
			h.Div(h.Class("navbar-end"), end),
		),
	)
}
