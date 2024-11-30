package comp

import (
	"fmt"
	"github.com/lpuig/cpmanager/model/consultant"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func ConsultantsTable(cslts *consultant.Container) g.Node {
	return h.Div(h.Style("margin: 15px 48px;"),
		h.Nav(h.Class("level"),
			h.Div(h.Class("level-left"),
				h.Button(h.Class("button is-primary"),
					//htmx attr
					//x.Get("/action/consult/add"),
					//x.Target("#consultant-list"),
					x.Get("/action/consult/addmodal"), x.Target(".modal"), x.Swap("outerHTML"),
					//children
					Icon("fas fa-user-plus"),
					h.Span(g.Text("Nouveau Consultant...")),
				),
			),
		),
		h.Div(h.ID("consultant-list"),
			ConsultantsList(cslts),
		),
	)
}

func ConsultantsList(cslts *consultant.Container) g.Node {
	return g.Map(cslts.GetSortedByName(), ConsultantLine)
}

func ConsultantLine(cslt consultant.Consultant) g.Node {
	return h.Div(
		h.ID(fmt.Sprintf("consultant-%s", cslt.Id)),
		h.Class("box"),
		h.Span(g.Text(cslt.Name())),
		h.A(Icon("fas fa-user-pen"),
			x.Trigger("click"), x.Get(fmt.Sprintf("/action/consult/updatemodal/%s", cslt.Id)), x.Target(".modal"), x.Swap("outerHTML"),
		),
	)
}
