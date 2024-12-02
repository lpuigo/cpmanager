package route

import (
	"github.com/lpuig/cpmanager/html"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/manager"
	"log/slog"
	"net/http"
)

type ManagerHandlerFunc func(*manager.Manager, http.ResponseWriter, *http.Request)

// Return main page
func GetMainPage(m *manager.Manager, w http.ResponseWriter, r *http.Request) {
	m.Log.InfoContext(r.Context(), "Get main page")
	html.MainPage(w, m)
}

// Return close modal elem
func GetCloseModal(m *manager.Manager, w http.ResponseWriter, r *http.Request) {
	m.Log.Log(r.Context(), slog.LevelInfo, "close modal")
	comp.ModalHook().Render(w)
}
