package comp

import (
	"fmt"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
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

func Columns(children ...g.Node) g.Node {
	return h.Div(h.Class("columns"),
		g.Map(children, func(n g.Node) g.Node {
			return h.Div(h.Class("column"), n)
		}),
	)
}

func Icon(name string) g.Node {
	return h.Span(h.Class("icon"), h.I(h.Class(name)))
}

func Image(src, alt string, children ...g.Node) g.Node {
	return h.Img(h.Src(src), h.Alt(alt), g.Group(children))
}

func ModalHook() g.Node {
	return h.Div(h.Class("modal"))
}

func modal(id string, children ...g.Node) g.Node {
	return h.Div(h.ID(id), h.Class("modal is-active"), x.Trigger("closeModal from:body"), x.Target("this"), x.Swap("outerHTML"), x.Get("/action/closemodal"),
		h.Div(h.Class("modal-background"),
			x.Trigger("click, keyup[key=='Escape'] from:body")), x.Target("closest .modal"), x.Swap("outerHTML"), x.Get("/action/closemodal"),
		g.Group(children),
	)
}

func Modal(id string, children ...g.Node) g.Node {
	return modal(id,
		h.Div(h.Class("modal-content"),
			g.Group(children),
		),
		h.Button(h.Class("modal-close is-large"), h.Aria("label", "close"), x.Target("closest .modal"), x.Swap("outerHTML"), x.Get("/action/closemodal")),
	)
}

func ModalCard(id string, children ...g.Node) g.Node {
	return modal(id,
		h.Div(h.Class("modal-card"),
			g.Group(children),
		),
	)
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

func Section(children ...g.Node) g.Node {
	return h.Section(h.Class("section"), g.Group(children))
}
