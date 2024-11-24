package comp

import (
	"github.com/lpuig/cpmanager/model/consultant"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func ConsultantsTable(cslts *consultant.Container) g.Node {
	return h.Div(
		h.Nav(h.Class("level"),
			h.Div(h.Class("level-left"),
				h.Button(h.Class("button is-primary"),
					//htmx attr
					x.Get("/action/consult/add"),
					x.Target("#consultant-list"),
					//children
					h.Span(h.Class("icon"), h.I(h.Class("fas fa-user-plus"))),
					h.Span(g.Text("Nouveau Consultant")),
				),
			),
		),
		h.Div(
			h.ID("consultant-list"),
			ConsultantsList(cslts),
		),
	)
}

func ConsultantsList(cslts *consultant.Container) g.Node {
	return g.Map(cslts.GetSortedByName(), ConsultantLine)
}

func ConsultantLine(cslt *consultant.Consultant) g.Node {
	return h.Div(
		h.Class("box"),
		h.Span(g.Text(cslt.Name())),
	)
}
