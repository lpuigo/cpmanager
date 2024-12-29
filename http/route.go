package http

import (
	"github.com/lpuig/cpmanager/http/route"
	"net/http"
)

func (s *Server) setupRoutes() {

	// define middleware
	withManager := func(rh route.ManagerHandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			m := s.Manager // clone server manager
			m.Log.StartTimer()
			m.Log.With("orig", "ROUTE", "method", r.Method, "path", r.URL.Path)
			rh(m, w, r)
		}
	}

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir(s.config.DirAsset))))

	s.mux.HandleFunc("GET /{$}", withManager(route.GetMainPage))

	s.mux.HandleFunc("GET /action/closemodal", withManager(route.GetCloseModal))

	// Add New Consultant
	s.mux.HandleFunc("GET /action/consult/addmodal", withManager(route.GetShowNewConsultantModal))
	s.mux.HandleFunc("POST /action/consult/add", withManager(route.PostAddNewConsultantFromModal))

	// Update Consultant
	s.mux.HandleFunc("GET /action/consult/updatemodal/{id}", withManager(route.GetShowUpdateConsultantModal))
	s.mux.HandleFunc("POST /action/consult/update/{id}", withManager(route.PostUpdateConsultantFromModal))

}
