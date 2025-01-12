package bulmacomp

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type TabItem struct {
	Tab  g.Node
	Body g.Node
}

func TabsHeader(active int, tabs []TabItem, isBoxed bool) g.Node {
	tabsItems := make([]g.Node, len(tabs))
	for i, tab := range tabs {
		tabsItems[i] = h.Li(g.If(i == active, h.Class("is-active")), h.A(tab.Tab))
	}
	class := "tabs"
	if isBoxed {
		class = "tabs is-boxed"
	}
	return h.Div(h.Class(class),
		h.Ul(g.Group(tabsItems)),
	)
}

func TabsBody(active int, tabs []TabItem) g.Node {
	if active >= len(tabs) {
		return nil
	}
	return tabs[active].Body
}
