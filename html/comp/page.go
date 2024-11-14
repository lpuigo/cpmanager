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
			gh.Link(gh.Rel("icon"), gh.Type("image/png"), gh.Href("/images/logo_coul.png")),
			gh.Script(gh.Src("script/script.js")),
		},
		Body: []g.Node{
			gh.Class("bg-gradient-to-b from-white to-indigo-100 bg-no-repeat"),
			gh.Div(
				gh.Class("min-h-screen flex flex-col justify-between"),
				header(),
				gh.Div(
					gh.Class("grow"),
					container(true,
						gh.Div(gh.Class("prose prose-lg prose-indigo"),
							g.Group(children),
						),
					),
				),
				footer(),
			),
		},
	}
	return gc.HTML5(props)
}

func header() g.Node {
	return gh.Div(gh.Class("bg-indigo-600 text-white shadow"),
		container(false,
			gh.Div(gh.Class("flex items-center space-x-4 h-8"),
				headerLink("/", "Home"),
				headerLink("/about", "About"),
			),
		),
	)
}

func headerLink(href, text string) g.Node {
	return gh.A(gh.Class("hover:text-indigo-300"), gh.Href(href), g.Text(text))
}

func container(padY bool, children ...g.Node) g.Node {
	return gh.Div(
		gc.Classes{
			"max-w-7xl mx-auto":     true,
			"px-4 md:px-8 lg:px-16": true,
			"py-4 md:py-8":          padY,
		},
		g.Group(children),
	)
}

func footer() g.Node {
	return gh.Div(gh.Class("bg-gray-900 text-white shadow text-center h-16 flex items-center justify-center"),
		gh.A(gh.Href("https://www.gomponents.com"), g.Text("gomponents")),
	)
}
