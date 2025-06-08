package route

import (
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/model/manager"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

// GetLoginPage renders the login page
func GetLoginPage(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	// Create a simple login form
	loginForm := bulmacomp.Section(
		h.Div(h.Class("columns is-centered"),
			h.Div(h.Class("column is-half"),
				h.Div(h.Class("box"),
					h.H1(h.Class("title has-text-centered"), g.Text("Login")),
					h.Form(h.Method("post"), h.Action("/login"),
						h.Div(h.Class("field"),
							h.Label(h.Class("label"), g.Text("Username")),
							h.Div(h.Class("control"),
								h.Input(h.Class("input"), h.Type("text"), h.Name("username"), h.Placeholder("Enter your username")),
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
	page := bulmacomp.Page("Login", loginForm)
	page.Render(w)

	m.Log.InfoContextWithTime(r.Context(), "Get login page")
}
