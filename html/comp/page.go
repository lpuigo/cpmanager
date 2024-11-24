package comp

import (
	g "maragu.dev/gomponents"
	gc "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func Page(title string, children ...g.Node) g.Node {
	props := gc.HTML5Props{
		Title:       title,
		Description: "une description",
		Language:    "fr",
		Head: []g.Node{
			h.Link(h.Rel("icon"), h.Type("image/png"), h.Href("/Assets/images/logo_coul.png")),
			h.Link(h.Rel("stylesheet"), h.Href("/Assets/bulma/bulma.min.css")),
			h.Link(h.Rel("stylesheet"), h.Href("/Assets/fontawesome/6.7.1/css/all.min.css")),
			h.Script(h.Src("/Assets/script/htmx/2.0.3/htmx.min.js")),
		},
		Body: []g.Node{
			header(),
			container(true,
				h.Div(h.Class("hero-body"),
					g.Group(children),
				),
			),
			footer(),
		},
	}
	return gc.HTML5(props)
}

func header() g.Node {
	return Navbar(false, "5px 48px",
		Image("/Assets/images/logo_coul.png", "to b services", h.Width("72px")),
		h.Div(
			h.P(h.Class("title"), g.Text("To B Services")),
			h.P(h.Class("subtitle"), g.Text("Gestion des Consultants")),
		),
		h.Div(),
	)
}

func container(padY bool, children ...g.Node) g.Node {
	return Section(g.Group(children))
}

func footer() g.Node {
	return Section(
		A("https://www.gomponents.com", "gomponents", true),
	)
}
