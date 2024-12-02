package route

import (
	"github.com/lpuig/cpmanager/html"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/manager"
	"net/http"
)

type ManagerHandlerFunc func(manager.Manager, http.ResponseWriter, *http.Request)

// Return main page
func GetMainPage(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	html.MainPage(w, m)
	m.Log.InfoContextWithTime(r.Context(), "Get main page")
}

// Return close modal elem
func GetCloseModal(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	comp.ModalHook().Render(w)
	m.Log.InfoContextWithTime(r.Context(), "close modal")
}
