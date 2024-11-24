package comp

import (
	"fmt"
	g "maragu.dev/gomponents"
	gc "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func A(href string, text string, newPage bool) g.Node {
	nodes := []g.Node{
		h.Href(href),
		g.Text(text),
	}
	if newPage {
		nodes = append(nodes, h.Target("_blank"))
	}
	return h.A(g.Group(nodes))
}

func Icon(name string) g.Node {
	return h.Span(h.Class("icon"), h.I(h.Class(name)))
}

func Image(src, alt string, children ...g.Node) g.Node {
	return h.Img(h.Src(src), h.Alt(alt), g.Group(children))
}

func Section(children ...g.Node) g.Node {
	return h.Section(h.Class("section"), g.Group(children))
}

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
