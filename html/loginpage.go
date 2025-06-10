package html

import (
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/model/manager"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

// LoginPage renders the login page
func LoginPage(m manager.Manager, w http.ResponseWriter) {
	// Create a simple login form
	loginForm := bulmacomp.Section(
		h.Div(h.Class("columns is-centered"),
			h.Div(h.Class("column is-half"),
				h.Div(h.Class("box"),
					h.H2(h.Class("title has-text-centered"), g.Text("Login")),
					h.Form(h.Method("post"), h.Action("/login"),
						h.Div(h.Class("field"),
							h.Label(h.Class("label"), g.Text("User")),
							h.Div(h.Class("control"),
								h.Input(h.Class("input"), h.Type("text"), h.Name("login"), h.Placeholder("Enter your user name")),
							),
						),
						h.Div(h.Class("field"),
							h.Label(h.Class("label"), g.Text("Password")),
							h.Div(h.Class("control"),
								h.Input(h.Class("input"), h.Type("password"), h.Name("password"), h.Placeholder("Enter your password")),
							),
						),
						h.Div(h.Class("field"),
							h.Div(h.Class("control"),
								h.Button(h.Class("button is-primary"), h.Type("submit"), g.Text("Login")),
							),
						),
					),
				),
			),
		),
	)

	// Render the login page
	page := bulmacomp.Page("Login", m, loginForm)
	page.Render(w)
}
