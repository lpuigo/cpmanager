package bulmacomp

import (
	g "maragu.dev/gomponents"
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

func Section(children ...g.Node) g.Node {
	return h.Section(h.Class("section"), g.Group(children))
}
