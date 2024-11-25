package comp

import (
	"github.com/lpuig/cpmanager/model/consultant"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func AddConsultantModal() g.Node {
	formName := "consultant-form"
	return ModalCard("add-consultant-modal",
		h.Header(h.Class("modal-card-head"),
			h.P(h.Class("modal-card-title"), g.Text("Consultant")),
			h.Button(h.Class("delete"), h.Aria("label", "close"),
				x.Target("closest .modal"), x.Swap("outerHTML"), x.Get("/action/closemodal"),
			),
		),
		h.Section(h.Class("modal-card-body"),
			h.Div(h.Class("panel-block"),
				h.Form(h.ID(formName), h.Class("content"), x.Post("/action/consult/add"), x.Trigger("submit"), x.Target("#consultant-list"),
					h.Div(h.Class("field"),
						h.Label(h.Class("label"), g.Text("Prénom")),
						h.Div(h.Class("control"),
							h.Input(h.Class("input"), h.Type("text"), h.Name("FirstName"), h.Placeholder("John")),
						),
					),
					h.Div(h.Class("field"),
						h.Label(h.Class("label"), g.Text("Nom")),
						h.Div(h.Class("control"),
							h.Input(h.Class("input"), h.Type("text"), h.Name("LastName"), h.Placeholder("Doe")),
						),
					),
				),
			),
		),
		h.Footer(h.Class("modal-card-foot"),
			h.Div(h.Class("buttons"),
				h.Button(h.Class("button is-success"), g.Text("Valider"), h.Type("submit"), h.FormAttr(formName)),
				h.Button(h.Class("button"), g.Text("Annuler"),
					x.Target("closest .modal"),
					x.Swap("outerHTML"),
					x.Get("/action/closemodal"),
				),
			),
		),
	)
}

func ConsultantsTable(cslts *consultant.Container) g.Node {
	return h.Div(h.Style("margin: 15px 48px;"),
		h.Nav(h.Class("level"),
			h.Div(h.Class("level-left"),
				h.Button(h.Class("button is-primary"),
					//htmx attr
					//x.Get("/action/consult/add"),
					//x.Target("#consultant-list"),
					x.Get("/action/consult/addmodal"),
					x.Target(".modal"),
					x.Swap("outerHTML"),
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

func ConsultantLine(cslt *consultant.Consultant) g.Node {
	return h.Div(
		h.Class("box"),
		h.Span(g.Text(cslt.Name())),
	)
}
