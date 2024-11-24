package comp

import (
	g "maragu.dev/gomponents"
	gc "maragu.dev/gomponents/components"
	gh "maragu.dev/gomponents/html"
)

func Page(title string, children ...g.Node) g.Node {
	props := gc.HTML5Props{
		Title:       title,
		Description: "une description",
		Language:    "fr",
		Head: []g.Node{
			gh.Link(gh.Rel("icon"), gh.Type("image/png"), gh.Href("/Assets/images/logo_coul.png")),
			gh.Link(gh.Rel("stylesheet"), gh.Href("/Assets/bulma/bulma.min.css")),
			gh.Link(gh.Rel("stylesheet"), gh.Href("/Assets/fontawesome/6.7.1/css/all.min.css")),
			gh.Script(gh.Src("/Assets/script/htmx/2.0.3/htmx.min.js")),
		},
		Body: []g.Node{
			header(),
			container(true,
				gh.Div(gh.Class("hero-body"),
					g.Group(children),
				),
			),
			footer(),
		},
	}
	return gc.HTML5(props)
}

func header() g.Node {
	return gh.Section(gh.Class("hero"), gh.Class("is-Small"), gh.Class("is-primary"),
		gh.Div(gh.Class("hero-body"),
			gh.P(gh.Class("title"), g.Text("To B Services")),
			gh.P(gh.Class("subtitle"), g.Text("Gestion des Consultants")),
		),
	)
}

func headerLink(href, text string) g.Node {
	return gh.A(gh.Class("hover:text-indigo-300"), gh.Href(href), g.Text(text))
}

func container(padY bool, children ...g.Node) g.Node {
	return gh.Section(gh.Class("section"),
		g.Group(children),
	)
}

func footer() g.Node {
	return gh.Section(gh.Class("section"),
		gh.A(gh.Href("https://www.gomponents.com"), gh.Target("_blank"), g.Text("gomponents")),
	)
}
