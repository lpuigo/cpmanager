package http

import (
	"github.com/lpuig/cpmanager/http/route"
	_ "github.com/lpuig/cpmanager/http/session" // Imported for side effects
	"net/http"
)

func (s *Server) setupRoutes() {

	// define middleware
	withManager := func(rh route.ManagerHandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			m := s.Manager.Clone()
			m.Log.StartTimer()
			session := m.Sessions.GetCurrentSessionFromRequest(r)
			m.User = *session.User
			m.Log.InitWith("orig", "ROUTE", "method", r.Method, "path", r.URL.Path)
			if session.IsLoggedIn() {
				m.LoggedIn = true
				m.Log.With("user", session.UserID)
			}
			rh(m, w, r)
		}
	}

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir(s.config.DirAsset))))

	// Add login/logout routes (public routes)
	s.mux.HandleFunc("GET /login", withManager(route.GetLoginPage))
	s.mux.HandleFunc("POST /login", withManager(route.HandleLogin))
	s.mux.HandleFunc("GET /logout", withManager(route.HandleLogout))

	// Add modal close action (needs session but not auth)
	s.mux.Handle("GET /action/closemodal", withManager(route.GetCloseModal))

	// Main page (protected route)
	s.mux.Handle("GET /{$}", withManager(route.GetMainPage))

	// Add New Consultant (protected routes)
	s.mux.Handle("GET /action/consult/addmodal", withManager(route.GetShowNewConsultantModal))
	s.mux.Handle("POST /action/consult/add", withManager(route.PostAddNewConsultantFromModal))

	// Update Consultant (protected routes)
	s.mux.Handle("GET /action/consult/{id}/updatemodal", withManager(route.GetShowUpdateConsultantModal))
	s.mux.Handle("POST /action/consult/{id}/update", withManager(route.PostUpdateConsultantFromModal))

	// Delete Consultant (protected route)
	s.mux.Handle("DELETE /action/consult/{id}", withManager(route.DeleteConsultant))

	// Add NewMission (protected routes)
	s.mux.Handle("GET /action/consult/{id}/addmissionmodal", withManager(route.GetShowAddMissionModal))
	s.mux.Handle("POST /action/consult/{id}/addmission", withManager(route.PostAddMissionFromModal))
}
