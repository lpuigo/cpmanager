package comp

import (
	"fmt"
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/model/consultant"
	"github.com/lpuig/cpmanager/model/date"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"strconv"
)

func missionForm(m *consultant.Mission, formname string, submitNode ...g.Node) g.Node {
	if m == nil {
		m = consultant.NewMission()
		m.StartDay = date.Today().ToDisplay()
	}

	return h.Form(h.ID(formname), h.Class("content full-width"),
		// HTMX for submit
		g.Group(submitNode),

		// Row Title ===================================================
		h.Div(h.Class("field"),
			h.Label(h.Class("label"), g.Text("Intitulé de la mission")),
			h.Div(h.Class("control"),
				h.Input(h.Class("input"), h.Type("text"), h.Name("Title"), h.Placeholder("Business Analyste"), h.Value(m.Title)),
			),
		),

		// Row Client, Manager ===================================================
		h.Div(h.Class("columns"),
			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Client")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("text"), h.Name("Company"), h.Placeholder("Nom Société"), h.Value(m.Company)),
					),
				),
			),

			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Manager")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("text"), h.Name("Manager"), h.Placeholder("Nom Manager"), h.Value(m.Manager)),
					),
				),
			),
		),

		// Cost and Rate Prices ===================================================
		h.Div(h.Class("columns"),
			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Taux Journalier")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("number"),
							h.Name("DailyRate"), h.Value(strconv.Itoa(consultant.GetLast(m.DailyRate).Price)),
						),
					),
				),
			),

			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Coût Journalier")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("number"),
							h.Name("DailyCost"), h.Value(strconv.Itoa(consultant.GetLast(m.DailyCost).Price)),
						),
					),
				),
			),
		),

		// Row Start, End Day ===================================================
		h.Div(h.Class("columns"),
			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Premier jour")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("date"), h.Name("StartDay"), h.Value(m.StartDay)),
					),
				),
			),

			h.Div(h.Class("column"),
				h.Div(h.Class("field"),
					h.Label(h.Class("label"), g.Text("Dernier jour")),
					h.Div(h.Class("control"),
						h.Input(h.Class("input"), h.Type("date"), h.Name("EndDay"), h.Value(m.EndDay)),
					),
				),
			),
		),
	)

}

func AddMissionModal(cslt *consultant.Consultant) g.Node {
	formName := "mission-form"
	return bulmacomp.ModalCardWithWidth("add-mission-modal", "60vw",
		h.Header(h.Class("modal-card-head"),
			h.P(h.Class("modal-card-title"), g.Text("Nouvelle mission pour "+cslt.Name())),
			h.Button(h.Class("delete"), h.Aria("label", "close"),
				x.Get("/action/closemodal"), x.Target("closest .modal"), x.Swap("outerHTML"),
			),
		),
		h.Section(h.Class("modal-card-body"),
			h.Div(h.Class("panel-block"),
				missionForm(nil, formName,
					x.Trigger("submit"),
					x.Post(fmt.Sprintf("/action/consult/%s/addmission", cslt.Id)),
					x.Target(fmt.Sprintf("#consultant-%s", cslt.Id)),
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
						x.Get("/action/closemodal"),
						x.Target("closest .modal"),
						x.Swap("outerHTML"),
					),
				),
			),
		),
	)
}
