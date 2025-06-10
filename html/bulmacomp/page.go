package bulmacomp

import (
	"github.com/lpuig/cpmanager/model/manager"
	g "maragu.dev/gomponents"
	gc "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func Page(title string, mgr manager.Manager, children ...g.Node) g.Node {
	props := gc.HTML5Props{
		Title:       title,
		Description: "une description",
		Language:    "fr",
		Head: []g.Node{
			h.Link(h.Rel("icon"), h.Type("image/png"), h.Href("/Assets/images/logo_coul.png")),
			h.Link(h.Rel("stylesheet"), h.Href("/Assets/bulma/bulma.min.css")),
			h.Link(h.Rel("stylesheet"), h.Href("/Assets/fontawesome/6.7.1/css/all.min.css")),
			h.Link(h.Rel("stylesheet"), h.Href("/Assets/style/cpmanager.css")),
			h.Script(h.Src("/Assets/script/htmx/2.0.3/htmx.min.js")),
		},
		Body: []g.Node{
			header(mgr),
			g.Group(children),
			footer(),
			ModalHook(),
		},
	}
	return gc.HTML5(props)
}

func header(mgr manager.Manager) g.Node {
	var endNode g.Node

	if mgr.LoggedIn {
		// Create logon widget with user icon and username
		endNode = h.Div(h.Class("navbar-item"),
			h.A(h.Href("/logout"),
				h.Span(h.Class("icon"),
					h.I(h.Class("fas fa-user")),
				),
				h.Span(g.Text(mgr.User.FullName)),
			),
		)
	} else {
		endNode = h.Div()
	}

	return Navbar(false, "5px 48px",
		Image("/Assets/images/logo_coul.png", "to b services", h.Width("72px")),
		h.Div(
			h.P(h.Class("title"), g.Text("To B Services")),
			h.P(h.Class("subtitle"), g.Text("Gestion des Consultants")),
		),
		endNode,
	)
}

func footer() g.Node {
	return Section(
		A("https://www.gomponents.com", "gomponents", true),
	)
}
