package bulmacomp

import (
	g "maragu.dev/gomponents"
	x "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func ModalHook() g.Node {
	return h.Div(h.Class("modal"))
}

func modal(id string, children ...g.Node) g.Node {
	return h.Div(h.ID(id), h.Class("modal is-active"),
		x.Trigger("closeModal from:body"), x.Target("this"), x.Swap("outerHTML"), x.Get("/action/closemodal"),
		h.Div(h.Class("modal-background"),
			x.Trigger("click, keyup[key=='Escape'] from:body"), x.Target("closest .modal"), x.Swap("outerHTML"), x.Get("/action/closemodal"),
		),
		g.Group(children),
	)
}

func Modal(id string, children ...g.Node) g.Node {
	return modal(id,
		h.Div(h.Class("modal-content"),
			g.Group(children),
		),
		h.Button(h.Class("modal-close is-large"), h.Aria("label", "close"), x.Target("closest .modal"), x.Swap("outerHTML"), x.Get("/action/closemodal")),
	)
}

func ModalCard(id string, children ...g.Node) g.Node {
	return modal(id,
		h.Div(h.Class("modal-card"),
			g.Group(children),
		),
	)
}

func ModalCardWithWidth(id, width string, children ...g.Node) g.Node {
	return modal(id,
		h.Div(h.Class("modal-card"), h.Style("width: "+width+";"),
			g.Group(children),
		),
	)
}
