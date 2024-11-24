package http

import (
	"github.com/lpuig/cpmanager/html"
	"github.com/lpuig/cpmanager/html/comp"
	"log/slog"
	"net/http"
)

func (s *Server) setupRoutes() {

	//TODO use compress middleware github.com/klauspost/compress
	s.mux.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir(s.config.Dir_Asset))))

	s.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Get root")
		html.MainPage(w, s.manager)
	})

	s.mux.HandleFunc("GET /action/consult/add", func(w http.ResponseWriter, r *http.Request) {
		s.log.Log(r.Context(), slog.LevelInfo, "Action Add consult")

		s.manager.Consultants.AddNewConsultant()
		comp.ConsultantsList(s.manager.Consultants).Render(w)
	})

}
