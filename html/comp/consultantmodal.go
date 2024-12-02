package comp

import (
	"fmt"
	"github.com/lpuig/cpmanager/model/consultant"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func UpdateConsultantModal(csl consultant.Consultant) g.Node {
	formName := "consultant-form"
	return ModalCard("add-consultant-modal",
		h.Header(h.Class("modal-card-head"),
			h.P(h.Class("modal-card-title"), g.Text("Consultant")),
			h.Button(h.Class("delete"), h.Aria("label", "close"),
				x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
			),
		),
		h.Section(h.Class("modal-card-body"),
			h.Div(h.Class("panel-block"),
				h.Form(h.ID(formName), h.Class("content"), h.Style("width: 100%;"),
					// HTMX for submit : call consult update with current consultant id
					x.Trigger("submit"),
					x.Post(fmt.Sprintf("/action/consult/update/%s", csl.Id)),
					x.Target(fmt.Sprintf("#consultant-%s", csl.Id)),
					x.Swap("outerHTML"),
					h.Div(h.Class("field"),
						h.Label(h.Class("label"), g.Text("Prénom")),
						h.Div(h.Class("control"),
							h.Input(h.Class("input"), h.Type("text"), h.Name("FirstName"), h.Value(csl.FirstName)),
						),
					),
					h.Div(h.Class("field"),
						h.Label(h.Class("label"), g.Text("Nom")),
						h.Div(h.Class("control"),
							h.Input(h.Class("input"), h.Type("text"), h.Name("LastName"), h.Value(csl.LastName)),
						),
					),
				),
			),
		),
		h.Footer(h.Class("modal-card-foot"),
			h.Div(h.Class("buttons"),
				h.Button(h.Class("button is-success"), g.Text("Valider"), h.Type("submit"), h.FormAttr(formName)),
				h.Button(h.Class("button"), g.Text("Annuler"),
					// HTMX for close Modal
					x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
				),
			),
		),
	)
}

func AddConsultantModal() g.Node {
	formName := "consultant-form"
	return ModalCard("add-consultant-modal",
		h.Header(h.Class("modal-card-head"),
			h.P(h.Class("modal-card-title"), g.Text("Consultant")),
			h.Button(h.Class("delete"), h.Aria("label", "close"),
				x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
			),
		),
		h.Section(h.Class("modal-card-body"),
			h.Div(h.Class("panel-block"),
				h.Form(h.ID(formName), h.Class("content"), h.Style("width: 100%;"),
					// HTMX for submit : call consutl add
					x.Post("/action/consult/add"), x.Trigger("submit"), x.Target("#consultant-list"), x.Swap("innerHTML"),
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
					x.Get("/action/closemodal"),
					x.Target("closest .modal"),
					x.Swap("outerHTML"),
				),
			),
		),
	)
}
