package http

import (
	"github.com/lpuig/cpmanager/http/route"
	"net/http"
)

func (s *Server) setupRoutes() {

	// define middleware
	withManager := func(rh route.ManagerHandlerFunc) http.HandlerFunc {
		m := *s.Manager // clone server manager
		m.Log.Reset()
		return func(w http.ResponseWriter, r *http.Request) {
			m.Log.Logger = s.Log.With("method", r.Method, "path", r.URL.Path)
			rh(&m, w, r)
		}
	}

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir(s.config.Dir_Asset))))

	s.mux.HandleFunc("GET /{$}", withManager(route.GetMainPage))

	s.mux.HandleFunc("GET /action/closemodal", withManager(route.GetCloseModal))

	// Add New Consultant
	s.mux.HandleFunc("GET /action/consult/addmodal", withManager(route.GetShowNewConsultantModal))
	s.mux.HandleFunc("POST /action/consult/add", withManager(route.PostAddNewConsultantFromModal))

	// Update Consultant
	s.mux.HandleFunc("GET /action/consult/updatemodal/{id}", withManager(route.GetShowUpdateConsultantModal))
	s.mux.HandleFunc("POST /action/consult/update/{id}", withManager(route.PostUpdateConsultantFromModal))

}
