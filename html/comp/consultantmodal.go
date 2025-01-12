package comp

import (
	"fmt"
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/model/consultant"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func consultantForm(csl *consultant.Consultant, formName string, submitNode ...g.Node) g.Node {
	if csl == nil {
		csl = consultant.NewConsultant()
	}
	return h.Form(h.ID(formName), h.Class("content full-width"),
		// HTMX for submit : call consult update with current consultant id
		g.Group(submitNode),

		// Row Prénom, Nom, CRMId ===================================================
		h.Div(h.Class("columns"),
			h.Div(h.Class("column is-two-fifths"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Prénom")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("text"), h.Name("FirstName"), h.Placeholder("Jean"), h.Value(csl.FirstName)),
					),
				),
			),

			h.Div(h.Class("column is-two-fifths"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Nom")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("text"), h.Name("LastName"), h.Placeholder("Dupont"), h.Value(csl.LastName)),
					),
				),
			),

			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Id CRM")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("text"), h.Name("CrmrId"), h.Placeholder("1234"), h.Value(csl.CrmrId)),
					),
				),
			),
		),

		// Row Profile ===================================================
		h.Div(h.Class("field"),
			h.Label(h.Class("label"), g.Text("Profil")),
			h.Div(h.Class("control"),
				h.Input(h.Class("input"), h.Type("text"), h.Name("Profile"), h.Placeholder("Developpeur"), h.Value(csl.Profile)),
			),
		),
	)
}

func UpdateConsultantModal(csl *consultant.Consultant) g.Node {
	formName := "consultant-form"
	return bulmacomp.ModalCardWithWidth("update-consultant-modal", "60vw",
		h.Header(h.Class("modal-card-head"),
			h.P(h.Class("modal-card-title"), g.Text("Consultant")),
			h.Button(h.Class("delete"), h.Aria("label", "close"),
				x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
			),
		),
		h.Section(h.Class("modal-card-body"),
			h.Div(h.Class("panel-block"),
				consultantForm(csl, formName,
					x.Trigger("submit"),
					x.Post(fmt.Sprintf("/action/consult/%s/update", csl.Id)),
					x.Target(fmt.Sprintf("#consultant-%s", csl.Id)),
					x.Swap("outerHTML"),
				),
			),
		),
		h.Footer(h.Class("modal-card-foot"),
			h.Div(h.Class("field is-grouped is-grouped-right full-width"),
				h.P(h.Class("control"),
					h.Button(h.Class("button is-success"), g.Text("Valider"), h.Type("submit"), h.FormAttr(formName)),
				),
				h.P(h.Class("control"),
					h.Button(h.Class("button"), g.Text("Annuler"),
						// HTMX for close Modal
						x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
					),
				),
			),
		),
	)
}

func AddConsultantModal() g.Node {
	formName := "consultant-form"
	return bulmacomp.ModalCardWithWidth("add-consultant-modal", "60vw",
		h.Header(h.Class("modal-card-head"),
			h.P(h.Class("modal-card-title"), g.Text("Consultant")),
			h.Button(h.Class("delete"), h.Aria("label", "close"),
				x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
			),
		),
		h.Section(h.Class("modal-card-body"),
			h.Div(h.Class("panel-block"),
				consultantForm(nil, formName,
					x.Post("/action/consult/add"), x.Trigger("submit"), x.Target("#consultant-list"), x.Swap("innerHTML"),
				),
			),
		),
		h.Footer(h.Class("modal-card-foot"),
			h.Div(h.Class("field is-grouped is-grouped-right full-width"),
				h.P(h.Class("control"),
					h.Button(h.Class("button is-success"), g.Text("Valider"), h.Type("submit"), h.FormAttr(formName)),
				),
				h.P(h.Class("control"),
					h.Button(h.Class("button"), g.Text("Annuler"),
						x.Get("/action/closemodal"),
						x.Target("closest .modal"),
						x.Swap("outerHTML"),
					),
				),
			),
		),
	)
}
