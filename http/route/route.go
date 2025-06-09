package route

import (
	"github.com/lpuig/cpmanager/html"
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/http/session"
	"github.com/lpuig/cpmanager/model/manager"
	"net/http"
)

// ManagerHandlerFunc is a function that handles an HTTP request with a manager
type ManagerHandlerFunc func(manager.Manager, http.ResponseWriter, *http.Request)

// SessionAwareHandlerFunc is a function that handles an HTTP request with a manager and a session
type SessionAwareHandlerFunc func(manager.Manager, *session.Session, http.ResponseWriter, *http.Request)

// Return login page
func GetLoginPage(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	html.LoginPage(m, w)
	m.Log.InfoContextWithTime(r.Context(), "Get login page")
}

// Return main page
func GetMainPage(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	html.MainPage(m, w)
	m.Log.InfoContextWithTime(r.Context(), "Get main page")
}

// Return close modal elem
func GetCloseModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	bulmacomp.ModalHook().Render(w)
	m.Log.InfoContextWithTime(r.Context(), "close modal")
}
