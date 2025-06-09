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
			m.Log.With("orig", "ROUTE", "method", r.Method, "path", r.URL.Path)
			rh(m, w, r)
		}
	}

	// Define middleware that adds session to the request context
	withSession := func(next http.Handler) http.Handler {
		return s.Manager.Sessions.WithSession(next)
	}

	// Define middleware that requires authentication
	withAuth := func(next http.Handler) http.Handler {
		return s.Manager.Sessions.RequireAuth(next)
	}

	// Define middleware that combines session and auth
	withSessionAndAuth := func(next http.Handler) http.Handler {
		return withSession(withAuth(next))
	}

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir(s.config.DirAsset))))

	// Add login/logout routes (public routes)
	s.mux.HandleFunc("GET /login", withManager(route.GetLoginPage))
	s.mux.HandleFunc("POST /login", withManager(route.HandleLogin))
	s.mux.HandleFunc("GET /logout", withManager(route.HandleLogout))

	// Add modal close action (needs session but not auth)
	s.mux.Handle("GET /action/closemodal", withSession(http.HandlerFunc(withManager(route.GetCloseModal))))

	// Main page (protected route)
	s.mux.Handle("GET /{$}", withSessionAndAuth(http.HandlerFunc(withManager(route.GetMainPage))))

	// Add New Consultant (protected routes)
	s.mux.Handle("GET /action/consult/addmodal", withSessionAndAuth(http.HandlerFunc(withManager(route.GetShowNewConsultantModal))))
	s.mux.Handle("POST /action/consult/add", withSessionAndAuth(http.HandlerFunc(withManager(route.PostAddNewConsultantFromModal))))

	// Update Consultant (protected routes)
	s.mux.Handle("GET /action/consult/{id}/updatemodal", withSessionAndAuth(http.HandlerFunc(withManager(route.GetShowUpdateConsultantModal))))
	s.mux.Handle("POST /action/consult/{id}/update", withSessionAndAuth(http.HandlerFunc(withManager(route.PostUpdateConsultantFromModal))))

	// Delete Consultant (protected route)
	s.mux.Handle("DELETE /action/consult/{id}", withSessionAndAuth(http.HandlerFunc(withManager(route.DeleteConsultant))))

	// Add NewMission (protected routes)
	s.mux.Handle("GET /action/consult/{id}/addmissionmodal", withSessionAndAuth(http.HandlerFunc(withManager(route.GetShowAddMissionModal))))
	s.mux.Handle("POST /action/consult/{id}/addmission", withSessionAndAuth(http.HandlerFunc(withManager(route.PostAddMissionFromModal))))
}
