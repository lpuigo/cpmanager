package comp

import (
	"fmt"
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/model/consultant"
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func ConsultantsBlock(cslts *consultant.ConsultantsPersister) g.Node {
	return h.Div(h.Style("margin: 15px 48px;"),
		h.Nav(h.Class("level"),
			h.Div(h.Class("level-left"),
				h.Button(h.Class("button is-primary"),
					//htmx attr
					//x.Get("/action/consult/add"),
					//x.Target("#consultant-list"),
					x.Get("/action/consult/addmodal"), x.Target(".modal"), x.Swap("outerHTML"),
					//children
					bulmacomp.Icon("fas fa-user-plus"),
					h.Span(g.Text("Nouveau Consultant...")),
				),
			),
		),
		h.Div(h.ID("consultant-list"),
			ConsultantTable(cslts),
		),
	)
}

//func ConsultantsList(cslts *consultant.ConsultantsPersister) g.Node {
//	return g.Map(cslts.GetSortedByName(), ConsultantLine)
//}
//
//func ConsultantLine(cslt *consultant.Consultant) g.Node {
//	return h.Div(
//		h.ID(fmt.Sprintf("consultant-%s", cslt.Id)),
//		h.Class("box"),
//		h.Span(g.Text(cslt.Name())),
//		h.A(Icon("fas fa-user-pen"),
//			x.Trigger("click"), x.Get(fmt.Sprintf("/action/consult/updatemodal/%s", cslt.Id)), x.Target(".modal"), x.Swap("outerHTML"),
//		),
//	)
//}

func ConsultantTable(cslts *consultant.ConsultantsPersister) g.Node {
	return h.Table(h.Class("table"),
		h.THead(
			h.Th(g.Text("Nom Prénom")),
			h.Th(g.Text("Profile")),
			g.Group(missionColumnHeader()),
			h.Th(g.Text("Actions")),
		),
		h.TBody(
			g.Map(cslts.GetSortedByName(), ConsultantTableRow),
		),
		//h.TFoot(
		//	h.Tr(),
		//),
	)
}

func ConsultantTableRow(cslt *consultant.Consultant) g.Node {
	nameNode := g.Group{g.Text(cslt.Name())}
	if cslt.CrmrId != "" {
		nameNode = append(nameNode,
			h.A(
				bulmacomp.Icon("fas fa-square-arrow-up-right"),
				h.Href(cslt.CrmUrl()),
				h.Target("_blank"),
			),
		)
	}

	return h.Tr(h.ID(fmt.Sprintf("consultant-%s", cslt.Id)),
		h.Td(nameNode),
		h.Td(g.Text(cslt.Profile)),
		g.Group(missionColumnRow(cslt)),
		h.Td(
			h.A(
				bulmacomp.Icon("fas fa-user-pen"),
				x.Trigger("click"), x.Get(fmt.Sprintf("/action/consult/%s/updatemodal", cslt.Id)), x.Target(".modal"), x.Swap("outerHTML"),
			),
			h.A(
				bulmacomp.Icon("fas fa-user-slash"),
				x.Trigger("click"), x.Delete(fmt.Sprintf("/action/consult/%s", cslt.Id)), x.Target(fmt.Sprintf("#consultant-%s", cslt.Id)), x.Swap("outerHTML"),
				x.Confirm(fmt.Sprintf("Supprimer le consultant %s ?", cslt.Name())),
			),
		),
	)
}

func missionColumnHeader() g.Group {
	return g.Group{
		h.Th(g.Text("Statut")),
		h.Th(g.Text("Client")),
		h.Th(g.Text("Manager")),
		h.Th(g.Text("Mission")),
	}
}

func missionColumnRow(cslt *consultant.Consultant) g.Group {
	var missionColumns g.Group

	if cslt.HasActiveMission() {
		m := cslt.LastMission()
		missionColumns = g.Group{
			// Statut
			h.Td(
				g.Text("Actif"),
				h.A(
					bulmacomp.Icon("fas fa-folder-open"),
					x.Trigger("click"), x.Get(fmt.Sprintf("/action/consult/%s/updatemissionmodal", cslt.Id)), x.Target(".modal"), x.Swap("outerHTML"),
				),
			),
			// Client
			h.Td(g.Text(m.Company)),
			// Manager
			h.Td(g.Text(m.Manager)),
			// Mission
			h.Td(g.Text(m.Title)),
		}
	} else {
		missionColumns = g.Group{
			// Statut
			h.Td(
				g.Text("Inactif"),
				h.A(
					bulmacomp.Icon("fas fa-folder-plus"),
					x.Trigger("click"), x.Get(fmt.Sprintf("/action/consult/%s/addmissionmodal", cslt.Id)), x.Target(".modal"), x.Swap("outerHTML"),
				),
			),
			// Client
			h.Td(g.Text("-")),
			// Manager
			h.Td(g.Text("-")),
			// Mission
			h.Td(g.Text("-")),
		}
	}

	return missionColumns
}
